package main

import (
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "os"
 "path/filepath"
 "strings"
 "testing"
)

const demoDOLA = `[{"id":"11","city":"90001","first_city":"FICTIONAL-CITY","the_geom":{"type":"MultiPolygon","coordinates":[[[[-105,40],[-104,40],[-104,41],[-105,40]]]]}}]`
const zeroHash = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func TestNormalizeDOLASyntheticOnly(t *testing.T){
 r,err:=normalizeDOLA(strings.NewReader(demoDOLA),zeroHash,1);if err!=nil{t.Fatal(err)}
 if len(r.Features)!=1||r.Features[0].Properties["status"]!="geographic_observation_only"{t.Fatalf("unexpected candidate result %+v",r)}
 if r.Features[0].Properties["jurisdiction_name"]!="FICTIONAL-CITY"||r.Features[0].Properties["source_city_code"]!="90001"{t.Fatalf("name/code mapping incorrect: %+v",r.Features[0].Properties)}
 if _,exists:=r.Features[0].Properties["code_adopted"];exists{t.Fatal("geometry importer asserted adoption")}
}
func TestNormalizeRejectsMalformedAndDuplicate(t *testing.T){
 cases:=[]string{
  `[]`,
  `[{"id":1,"city":"90001","first_city":"FICTIONAL","the_geom":{"type":"Point","coordinates":[]}}]`,
  `[{"id":1,"city":"90001","first_city":"FICTIONAL","the_geom":{"type":"MultiPolygon","coordinates":[[[[-105,40],[-104,40],[-104,41],[-105,40]]]]}},{"id":1,"city":"90001","first_city":"FICTIONAL","the_geom":{"type":"MultiPolygon","coordinates":[[[[-105,40],[-104,40],[-104,41],[-105,40]]]]}}]`,
  `[{"id":1,"city":"90001","first_city":"FICTIONAL","the_geom":{"type":"MultiPolygon","coordinates":[[[[-105,40],[-104,40],[-104,41],[-105,40]]]]}}] {}`,
  `[{"id":1,"city":"90001","first_city":"FICTIONAL","the_geom":{"type":"MultiPolygon","coordinates":[[[[-999,40],[-104,40],[-104,41],[-999,40]]]]}}]`,
 }
 for _,raw:=range cases{if _,err:=normalizeDOLA(strings.NewReader(raw),zeroHash,1);err==nil{t.Errorf("accepted invalid input: %s",raw)}}
 if _,err:=normalizeDOLA(strings.NewReader(demoDOLA),"wrong",1);err==nil{t.Fatal("accepted invalid digest")}
}
func TestFetchContentIdentifiedAndNoOverwrite(t *testing.T){
 data:=[]byte("FICTIONAL synthetic source bytes")
 server:=httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.Write(data)}));defer server.Close()
 src:=Source{ID:"synthetic",Publisher:"FICTIONAL",URL:server.URL,Format:"txt"};out:=filepath.Join(t.TempDir(),"raw.txt")
 h:=sha256.Sum256(data);expected:=hex.EncodeToString(h[:]);receipt,err:=fetch(server.Client(),src,out,expected,false);if err!=nil{t.Fatal(err)}
 if receipt.SHA256!=expected||receipt.State!="hash_matched_not_legal_evidence"{t.Fatalf("unexpected receipt %+v",receipt)}
 raw,err:=os.ReadFile(out);if err!=nil||string(raw)!=string(data){t.Fatalf("missing download: %v",err)}
 b,err:=os.ReadFile(out+".receipt.json");if err!=nil{t.Fatal(err)};var decoded Receipt;if err=json.Unmarshal(b,&decoded);err!=nil||decoded.SHA256!=expected{t.Fatalf("bad receipt: %v",err)}
 if _,err=fetch(server.Client(),src,out,expected,false);err==nil{t.Fatal("overwrote immutable download")}
}
func TestFetchFailsClosedOnHashAndSize(t *testing.T){
 server:=httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.Write([]byte("a"))}));defer server.Close()
 dir:=t.TempDir();src:=Source{ID:"demo",URL:server.URL,Large:true}
 if _,err:=fetch(server.Client(),src,filepath.Join(dir,"large"),"",false);err==nil{t.Fatal("large archive not guarded")}
 out:=filepath.Join(dir,"mismatch");if _,err:=fetch(server.Client(),src,out,zeroHash,true);err==nil{t.Fatal("incorrect hash accepted")}
 if _,err:=os.Stat(out);!os.IsNotExist(err){t.Fatalf("wrote unverified raw file: %v",err)}
 if _,err:=os.Stat(out+".receipt.json");!os.IsNotExist(err){t.Fatalf("wrote unverified receipt: %v",err)}
}
func TestNormalizeFileVerifiesReceipt(t *testing.T){
 dir:=t.TempDir();raw:=filepath.Join(dir,"source.json");os.WriteFile(raw,[]byte(demoDOLA),0600)
 sum:=sha256.Sum256([]byte(demoDOLA));sha:=hex.EncodeToString(sum[:]);receiptPath:=raw+".receipt.json"
 if err:=writeNewJSON(receiptPath,Receipt{SourceID:"co-dola-municipal",SourceURL:sources[4].URL,SHA256:sha,Bytes:int64(len(demoDOLA)),State:"unreviewed"});err!=nil{t.Fatal(err)}
 out:=filepath.Join(dir,"features.geojson");if err:=normalizeFile(raw,receiptPath,out,1);err!=nil{t.Fatal(err)}
 if _,err:=os.Stat(out);err!=nil{t.Fatal(err)}
 os.WriteFile(raw,[]byte("altered"),0600)
 if err:=normalizeFile(raw,receiptPath,filepath.Join(dir,"second.geojson"),1);err==nil{t.Fatal("tampered source admitted")}
}

func TestDOLARejectsMissingNameOrIncompleteSnapshot(t *testing.T){
 raw:=strings.Replace(demoDOLA,`"first_city":"FICTIONAL-CITY"`,`"first_city":""`,1)
 if _,err:=normalizeDOLA(strings.NewReader(raw),zeroHash,1);err==nil{t.Fatal("missing municipality name accepted")}
 if _,err:=normalizeDOLA(strings.NewReader(demoDOLA),zeroHash,274);err==nil{t.Fatal("truncated source accepted")}
 if _,err:=normalizeDOLA(strings.NewReader(demoDOLA),zeroHash,0);err==nil{t.Fatal("missing expected count accepted")}
}
func TestDOLARejectsReceiptSourceSpoofing(t *testing.T){
 dir:=t.TempDir();raw:=filepath.Join(dir,"raw.json")
 if err:=os.WriteFile(raw,[]byte(demoDOLA),0600);err!=nil{t.Fatal(err)}
 sum:=sha256.Sum256([]byte(demoDOLA));hash:=hex.EncodeToString(sum[:])
 receipt:=filepath.Join(dir,"spoof.json")
 if err:=writeNewJSON(receipt,Receipt{SourceID:"co-dola-municipal",SourceURL:"https://example.com/other.json",SHA256:hash,Bytes:int64(len(demoDOLA)),State:"unreviewed"});err!=nil{t.Fatal(err)}
 if err:=normalizeFile(raw,receipt,filepath.Join(dir,"out.geojson"),1);err==nil{t.Fatal("spoofed receipt source accepted")}
}
