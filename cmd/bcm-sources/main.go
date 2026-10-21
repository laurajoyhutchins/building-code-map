// Command bcm-sources acquires public research inputs; it never certifies code adoption.
package main

import (
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "flag"
 "fmt"
 "io"
 "net/http"
 "os"
 "path/filepath"
 "strings"
 "time"
)

type Source struct { ID, Publisher, URL, Format, Purpose string; Large bool }
var sources = []Source{
 {"doe-residential", "US DOE Building Energy Codes Program", "https://www.energycodes.gov/sites/default/files/2026-09/StateLevelResidentialCodesEnergyUseIndex.xlsx", "xlsx", "State residential energy-code equivalence; not adopted legal text", false},
 {"doe-commercial", "US DOE Building Energy Codes Program", "https://www.energycodes.gov/sites/default/files/2026-09/StateLevelCommercialCodesEnergyUseIndex.xlsx", "xlsx", "State commercial energy-code equivalence; not adopted legal text", false},
 {"fema-bcat-2026-06", "Public Environmental Data Partners (FEMA BCAT preservation)", "https://zenodo.org/records/21084129/files/FEMA%20BCAT.gdb.zip?download=1", "zip", "Historical third-party BCAT extract; adoption research leads only", true},
 {"usdot-nad-2026-09", "US Department of Transportation", "https://data.transportation.gov/download/fc2s-wawr/application/x-zip-compressed", "zip", "National address-point input; incomplete coverage", true},
 {"co-dola-municipal", "Colorado DOLA via Colorado Information Marketplace", "https://data.colorado.gov/resource/w6jq-7em5.json?%24limit=50000&%24order=id", "json", "Best-effort municipal geometry, not legal boundaries", false},
 {"census-co-places-2026", "US Census Bureau", "https://www2.census.gov/geo/tiger/TIGER2026/PLACE/tl_2026_08_place.zip", "zip", "2026 Colorado census place geometries, not code authority", false},
}
func lookup(id string) (Source,error) { for _,s:=range sources {if s.ID==id{return s,nil}};return Source{},fmt.Errorf("unknown source %q",id) }

type Receipt struct {
 SourceID string `json:"source_id"`
 Publisher string `json:"publisher"`
 SourceURL string `json:"source_url"`
 RetrievedAt string `json:"retrieved_at"`
 SHA256 string `json:"sha256"`
 Bytes int64 `json:"bytes"`
 State string `json:"state"`
 Purpose string `json:"purpose"`
}

func fetch(client *http.Client,source Source,dest,expected string,allowLarge bool)(Receipt,error){
 if source.Large&&!allowLarge{return Receipt{},errors.New("large source requires -allow-large")}
 if dest==""{return Receipt{},errors.New("-out is required")}
 if expected!=""&&(len(expected)!=64||strings.Trim(expected,"0123456789abcdef")!=""){return Receipt{},errors.New("expected SHA256 must be 64 lowercase hex characters")}
 if _,e:=os.Stat(dest);e==nil{return Receipt{},errors.New("output already exists; choose a new immutable path")}else if !errors.Is(e,os.ErrNotExist){return Receipt{},e}
 resp,err:=client.Get(source.URL);if err!=nil{return Receipt{},err};defer resp.Body.Close()
 if resp.StatusCode!=http.StatusOK{return Receipt{},fmt.Errorf("source returned HTTP %d",resp.StatusCode)}
 limit:=int64(128<<20);if allowLarge{limit=4<<30}
 if resp.ContentLength>limit{return Receipt{},errors.New("source exceeds configured byte limit")}
 if err=os.MkdirAll(filepath.Dir(dest),0700);err!=nil{return Receipt{},err}
 tmp,err:=os.CreateTemp(filepath.Dir(dest),".bcm-source-*");if err!=nil{return Receipt{},err}
 defer os.Remove(tmp.Name());defer tmp.Close()
 h:=sha256.New();n,err:=io.Copy(io.MultiWriter(tmp,h),io.LimitReader(resp.Body,limit+1));if err!=nil{return Receipt{},err}
 if n>limit{return Receipt{},errors.New("source exceeded configured byte limit")}
 sum:=hex.EncodeToString(h.Sum(nil));if expected!=""&&sum!=expected{return Receipt{},fmt.Errorf("SHA256 mismatch: got %s",sum)}
 if err=tmp.Sync();err!=nil{return Receipt{},err}
 if err=tmp.Close();err!=nil{return Receipt{},err}
 receipt:=Receipt{source.ID,source.Publisher,source.URL,time.Now().UTC().Format(time.RFC3339),sum,n,"unreviewed",source.Purpose}
 if expected!=""{receipt.State="hash_matched_not_legal_evidence"}
 // Raw sources and receipts stay in local/private custody. Nothing here grants regulatory authority.
 if err=writeNewJSON(dest+".receipt.json",receipt);err!=nil{return Receipt{},err}
 if err=os.Rename(tmp.Name(),dest);err!=nil{os.Remove(dest+".receipt.json");return Receipt{},err}
 return receipt,nil
}
func writeNewJSON(path string,v any)error{
 f,err:=os.OpenFile(path,os.O_WRONLY|os.O_CREATE|os.O_EXCL,0600);if err!=nil{return err}
 defer f.Close();e:=json.NewEncoder(f);e.SetIndent("","  ");if err=e.Encode(v);err!=nil{os.Remove(path);return err};return f.Sync()
}
type geom struct {Type string `json:"type"`;Coordinates [][][][]float64 `json:"coordinates"`}
type dolaRow struct {ID json.Number `json:"id"`;City string `json:"city"`;FirstCity string `json:"first_city"`;Geometry geom `json:"the_geom"`}
type feature struct {Type string `json:"type"`;Geometry geom `json:"geometry"`;Properties map[string]string `json:"properties"`}
type featureCollection struct {Type string `json:"type"`;Features []feature `json:"features"`}
func normalizeDOLA(input io.Reader,sourceHash string)(featureCollection,error){
 if len(sourceHash)!=64||strings.Trim(sourceHash,"0123456789abcdef")!=""{return featureCollection{},errors.New("source SHA256 required")}
 dec:=json.NewDecoder(input);dec.UseNumber();var rows []dolaRow
 if err:=dec.Decode(&rows);err!=nil{return featureCollection{},err}
 if len(rows)==0||len(rows)>10000{return featureCollection{},errors.New("unexpected DOLA record count")}
 var extra any;if err:=dec.Decode(&extra);err!=io.EOF{return featureCollection{},errors.New("trailing JSON")}
 result:=featureCollection{Type:"FeatureCollection",Features:make([]feature,0,len(rows))};seen:=map[string]bool{}
 for _,r:=range rows{
  id:=r.ID.String();if id==""||strings.TrimSpace(r.City)==""||seen[id]{return featureCollection{},fmt.Errorf("missing/duplicate ID or city: %q",id)};seen[id]=true
  if r.Geometry.Type!="MultiPolygon"||len(r.Geometry.Coordinates)==0{return featureCollection{},fmt.Errorf("invalid geometry for ID %s",id)}
  for _,polygon:=range r.Geometry.Coordinates{if len(polygon)==0{return featureCollection{},fmt.Errorf("empty polygon %s",id)};for _,ring:=range polygon{
   if len(ring)<4{return featureCollection{},fmt.Errorf("short ring %s",id)}
   for _,p:=range ring{if len(p)!=2||p[0]< -180||p[0]>180||p[1]< -90||p[1]>90{return featureCollection{},fmt.Errorf("invalid lon/lat %s",id)}}
   a,b:=ring[0],ring[len(ring)-1];if a[0]!=b[0]||a[1]!=b[1]{return featureCollection{},fmt.Errorf("open ring %s",id)}
  }}
  result.Features=append(result.Features,feature{Type:"Feature",Geometry:r.Geometry,Properties:map[string]string{
   "source_id":"co-dola-municipal","source_record_id":id,"jurisdiction_name":strings.TrimSpace(r.City),"source_sha256":sourceHash,"status":"geographic_observation_only",
  }})
 }
 return result,nil
}
func normalizeFile(input,receiptPath,out string)error{
 var r Receipt;rb,err:=os.ReadFile(receiptPath);if err!=nil{return err};if err=json.Unmarshal(rb,&r);err!=nil{return err}
 if r.SourceID!="co-dola-municipal"{return errors.New("receipt is not a Colorado DOLA source")}
 f,err:=os.Open(input);if err!=nil{return err};defer f.Close()
 h:=sha256.New();if _,err=io.Copy(h,f);err!=nil{return err};if hex.EncodeToString(h.Sum(nil))!=r.SHA256{return errors.New("source content does not match receipt SHA256")}
 if _,err=f.Seek(0,io.SeekStart);err!=nil{return err};collection,err:=normalizeDOLA(f,r.SHA256);if err!=nil{return err}
 return writeNewJSON(out,collection)
}
func main(){
 if len(os.Args)<2{fmt.Fprintln(os.Stderr,"usage: bcm-sources list | fetch -source ID -out FILE [-sha256 SHA] [-allow-large] | normalize-colorado -in FILE -receipt FILE -out FILE");os.Exit(2)}
 var err error
 switch os.Args[1]{
 case "list":for _,s:=range sources{fmt.Printf("%s\t%s\t%s\n",s.ID,s.Format,s.Purpose)}
 case "fetch":fs:=flag.NewFlagSet("fetch",flag.ExitOnError);id:=fs.String("source","","listed source ID");out:=fs.String("out","","private destination");sha:=fs.String("sha256","","expected content hash");large:=fs.Bool("allow-large",false,"opt in to large archives");fs.Parse(os.Args[2:]);var src Source;src,err=lookup(*id);if err==nil{client:=&http.Client{Timeout:30*time.Minute};client.CheckRedirect=func(req *http.Request,via []*http.Request)error{if req.URL.Scheme!="https"{return errors.New("redirect left HTTPS")};if len(via)>=5{return errors.New("too many redirects")};return nil};var receipt Receipt;receipt,err=fetch(client,src,*out,*sha,*large);if err==nil{fmt.Printf("acquired %s (%d bytes), SHA256 %s, status %s\n",*id,receipt.Bytes,receipt.SHA256,receipt.State)}}
 case "normalize-colorado":fs:=flag.NewFlagSet("normalize-colorado",flag.ExitOnError);input:=fs.String("in","","downloaded JSON");receipt:=fs.String("receipt","","matching fetch receipt");out:=fs.String("out","","private candidate GeoJSON output");fs.Parse(os.Args[2:]);if *input==""||*receipt==""||*out==""{err=errors.New("-in, -receipt, and -out are required")}else{err=normalizeFile(*input,*receipt,*out)}
 default:err=fmt.Errorf("unknown command %q",os.Args[1])
 }
 if err!=nil{fmt.Fprintln(os.Stderr,"bcm-sources:",err);os.Exit(1)}
}
