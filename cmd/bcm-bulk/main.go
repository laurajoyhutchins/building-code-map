// Command bcm-bulk executes bounded, project-scoped verification workloads.
// It never substitutes a bulk legal verdict or publishes the provider's response.
package main

import (
 "bufio"
 "bytes"
 "context"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "flag"
 "fmt"
 "io"
 "math"
 "net"
 "net/http"
 "net/url"
 "os"
 "sort"
 "strings"
 "sync"
 "time"
)

type point struct {
 Longitude float64 `json:"longitude"`
 Latitude float64 `json:"latitude"`
}
type sample struct {
 ID string `json:"id"`
 Point *point `json:"point,omitempty"`
 Address string `json:"address,omitempty"`
 ProjectType string `json:"project_type,omitempty"`
 Facts map[string]string `json:"facts,omitempty"`
}
type outcome struct {
 ID string `json:"id"`
 Category string `json:"category"`
 HTTPStatus int `json:"http_status,omitempty"`
 ErrorCode string `json:"error_code,omitempty"`
 Verdict string `json:"verdict,omitempty"`
 JurisdictionID string `json:"jurisdiction_id,omitempty"`
 CodeFamilies []string `json:"code_families,omitempty"`
 EvidenceCount int `json:"evidence_count,omitempty"`
 LatencyMS float64 `json:"latency_ms"`
}
type summary struct {
 SchemaVersion string `json:"schema_version"`
 InputSHA256 string `json:"input_sha256"`
 EndpointPath string `json:"endpoint_path"`
 ApplicabilityDate string `json:"applicability_date"`
 Samples int `json:"samples"`
 Categories map[string]int `json:"categories"`
 HTTPStatuses map[string]int `json:"http_statuses"`
 CodeFamilies map[string]int `json:"code_families"`
 DurationMS int64 `json:"duration_ms"`
 RequestsPerSecond float64 `json:"requests_per_second"`
 LatencyP50MS float64 `json:"latency_p50_ms"`
 LatencyP95MS float64 `json:"latency_p95_ms"`
 LatencyP99MS float64 `json:"latency_p99_ms"`
}
type config struct {
 Input, Results, Summary, Endpoint, Date, ProjectType, TokenEnv string
 Workers, QPS, MaxRequests int
 Timeout time.Duration
 AllowRemote bool
}
func validSample(s sample) error {
 if len(s.ID)==0 || len(s.ID)>128 { return errors.New("sample id length must be 1..128") }
 for _,r:=range s.ID { if !((r>='a'&&r<='z')||(r>='A'&&r<='Z')||(r>='0'&&r<='9')||r=='-'||r=='_'||r==':'||r=='.') {return errors.New("sample id has invalid characters")} }
 if (s.Point==nil)==(strings.TrimSpace(s.Address)=="") {return errors.New("exactly one of point or address required")}
 if s.Point!=nil {
  p:=s.Point
  if math.IsNaN(p.Latitude)||math.IsInf(p.Latitude,0)||math.IsNaN(p.Longitude)||math.IsInf(p.Longitude,0)||p.Latitude < -90||p.Latitude>90||p.Longitude < -180||p.Longitude>180 {return errors.New("invalid coordinates")}
 }
 return nil
}
func validate(c config) error {
 if c.Input==""||c.Results==""||c.Summary=="" {return errors.New("input, results and summary paths are required")}
 if c.Input==c.Results||c.Input==c.Summary||c.Results==c.Summary {return errors.New("input/results/summary paths must differ")}
 if c.Workers<1||c.Workers>64||c.QPS<1||c.QPS>100||c.MaxRequests<1||c.MaxRequests>5000000 {return errors.New("workers 1..64, qps 1..100, max-requests 1..5000000 required")}
 if c.Timeout<time.Second||c.Timeout>120*time.Second {return errors.New("timeout must be between 1s and 120s")}
 if t,e:=time.Parse(time.DateOnly,c.Date);e!=nil||t.Format(time.DateOnly)!=c.Date{return errors.New("date must be YYYY-MM-DD")}
 u,e:=url.Parse(c.Endpoint)
 if e!=nil||u==nil||u.User!=nil||u.RawQuery!=""||u.Fragment!=""||u.Hostname()==""||u.Path!="/v1/project-code-basis" {return errors.New("endpoint must be an origin plus /v1/project-code-basis without credentials or query")}
 if u.Scheme!="https"&&u.Scheme!="http" {return errors.New("endpoint must use HTTP(S)")}
 host:=u.Hostname()
 ip:=net.ParseIP(host)
 local:=strings.EqualFold(host,"localhost")||(ip!=nil&&ip.IsLoopback())
 if !local&&!c.AllowRemote {return errors.New("remote target requires -allow-remote")}
 if !local&&u.Scheme!="https" {return errors.New("remote target requires HTTPS")}
 return nil
}
func classify(status int, body []byte, id, date string) outcome {
 r:=outcome{ID:id,HTTPStatus:status}
 if status==http.StatusOK {
  var value struct {
   ProjectID string `json:"project_id"`
   AsOf string `json:"as_of"`
   SchemaVersion string `json:"schema_version"`
   Verdict string `json:"verdict"`
   JurisdictionID string `json:"jurisdiction_id"`
   CodeSet []struct { Family string `json:"family"`; Status string `json:"status"` } `json:"code_set"`
   Evidence []json.RawMessage `json:"evidence"`
   ExactEvidence []json.RawMessage `json:"exact_evidence"`
  }
  if err:=json.Unmarshal(body,&value);err!=nil||value.ProjectID!=id||value.AsOf!=date||value.SchemaVersion!="0.3"||value.JurisdictionID==""||len(value.CodeSet)==0 {r.Category="invalid_response";return r}
  switch value.Verdict {case "verified","conditional","not_verified":default:r.Category="invalid_response";return r}
  families:=make(map[string]bool)
  for _,c:=range value.CodeSet {if c.Family==""||c.Status=="" {r.Category="invalid_response";return r};families[c.Family]=true}
  if value.Verdict=="verified"&&len(value.Evidence)+len(value.ExactEvidence)==0 {r.Category="unsupported_verification";return r}
  r.Category=value.Verdict;r.Verdict=value.Verdict;r.JurisdictionID=value.JurisdictionID;r.EvidenceCount=len(value.Evidence)+len(value.ExactEvidence)
  for family:=range families {r.CodeFamilies=append(r.CodeFamilies,family)}
  sort.Strings(r.CodeFamilies)
  return r
 }
 var envelope struct {Error struct {Code string `json:"code"`} `json:"error"`}
 if json.Unmarshal(body,&envelope)==nil {r.ErrorCode=envelope.Error.Code}
 if (status==404||status==422)&&map[string]bool{"address_not_found":true,"jurisdiction_profile_missing":true,"address_ambiguous":true,"boundary_ambiguous":true,"outside_supported_coverage":true}[r.ErrorCode] {r.Category="unresolved";return r}
 if status==429||status>=500 {r.Category="service_error";return r}
 r.Category="http_error";return r
}
func query(ctx context.Context, client *http.Client, c config, s sample) outcome {
 requestBody:=map[string]any{"project_id":s.ID,"applicability_date":c.Date}
 if s.Point!=nil {requestBody["point"]=s.Point} else {requestBody["address"]=s.Address}
 if s.ProjectType!="" {requestBody["project_type"]=s.ProjectType} else if c.ProjectType!="" {requestBody["project_type"]=c.ProjectType}
 if len(s.Facts)>0 {requestBody["facts"]=s.Facts}
 data,_:=json.Marshal(requestBody)
 started:=time.Now()
 r:=outcome{ID:s.ID}
 rq,err:=http.NewRequestWithContext(ctx,http.MethodPost,c.Endpoint,bytes.NewReader(data))
 if err==nil {
  rq.Header.Set("Content-Type","application/json")
  rq.Header.Set("X-Request-ID",s.ID)
  if c.TokenEnv!="" {rq.Header.Set("Authorization","Bearer "+os.Getenv(c.TokenEnv))}
  var response *http.Response
  response,err=client.Do(rq)
  if err==nil {
   limited:=io.LimitReader(response.Body,1024*1024+1)
   var body []byte
   body,err=io.ReadAll(limited)
   response.Body.Close()
   if err==nil&&len(body)<=1024*1024 {r=classify(response.StatusCode,body,s.ID,c.Date)} else {r.Category="invalid_response";r.HTTPStatus=response.StatusCode}
  }
 }
 if err!=nil {r.Category="transport_error"}
 r.LatencyMS=float64(time.Since(started).Microseconds())/1000
 return r
}
func percentile(values []float64, fraction float64) float64 {
 if len(values)==0{return 0}
 index:=int(math.Ceil(float64(len(values))*fraction))-1
 if index<0{index=0};if index>=len(values){index=len(values)-1}
 return values[index]
}
func execute(ctx context.Context,c config) (summary,error) {
 result:=summary{SchemaVersion:"1",EndpointPath:"/v1/project-code-basis",ApplicabilityDate:c.Date,Categories:map[string]int{},HTTPStatuses:map[string]int{},CodeFamilies:map[string]int{}}
 if err:=validate(c);err!=nil{return result,err}
 input,err:=os.Open(c.Input);if err!=nil{return result,err};defer input.Close()
 tmp:=c.Results+".tmp"
 output,err:=os.Create(tmp);if err!=nil{return result,err}
 defer func(){output.Close();os.Remove(tmp)}()
 ctx,cancel:=context.WithCancel(ctx);defer cancel()
 client:=&http.Client{Timeout:c.Timeout,CheckRedirect:func(*http.Request,[]*http.Request)error{return errors.New("redirects forbidden")}}
 tick:=time.NewTicker(time.Second/time.Duration(c.QPS));defer tick.Stop()
 jobs:=make(chan sample,c.Workers*2)
 outcomes:=make(chan outcome,c.Workers*2)
 var wg sync.WaitGroup
 for i:=0;i<c.Workers;i++ {
  wg.Add(1)
  go func(){defer wg.Done();for s:=range jobs {select {case <-ctx.Done():return;case <-tick.C:}
    r:=query(ctx,client,c,s);select {case outcomes<-r:case <-ctx.Done():return}
  }}()
 }
 hash:=sha256.New()
 type feederResult struct {n int;err error}
 finished:=make(chan feederResult,1)
 go func(){
  defer close(jobs)
  scanner:=bufio.NewScanner(io.TeeReader(input,hash))
  scanner.Buffer(make([]byte,64*1024),1024*1024)
  seen:=make(map[string]struct{})
  n:=0
  for scanner.Scan() {
   line:=bytes.TrimSpace(scanner.Bytes());if len(line)==0{continue}
   if n>=c.MaxRequests{finished<-feederResult{n,fmt.Errorf("input exceeds max-requests %d",c.MaxRequests)};cancel();return}
   var s sample
   if e:=json.Unmarshal(line,&s);e!=nil{finished<-feederResult{n,fmt.Errorf("invalid JSONL sample: %w",e)};cancel();return}
   if e:=validSample(s);e!=nil{finished<-feederResult{n,fmt.Errorf("sample %q: %w",s.ID,e)};cancel();return}
   if _,ok:=seen[s.ID];ok{finished<-feederResult{n,fmt.Errorf("duplicate sample id %q",s.ID)};cancel();return};seen[s.ID]=struct{}{}
   select{case jobs<-s:n++;case <-ctx.Done():finished<-feederResult{n,ctx.Err()};return}
  }
  finished<-feederResult{n,scanner.Err()}
 }()
 go func(){wg.Wait();close(outcomes)}()
 started:=time.Now()
 encoder:=json.NewEncoder(output)
 latencies:=make([]float64,0,1024)
 var writeErr error
 for r:=range outcomes {
  if e:=encoder.Encode(r);e!=nil&&writeErr==nil {writeErr=e;cancel()}
  result.Samples++;result.Categories[r.Category]++
  result.HTTPStatuses[fmt.Sprint(r.HTTPStatus)]++
  for _,family:=range r.CodeFamilies {result.CodeFamilies[family]++}
  latencies=append(latencies,r.LatencyMS)
 }
 fr:=<-finished
 if fr.err!=nil{return result,fr.err}
 if writeErr!=nil{return result,writeErr}
 if result.Samples!=fr.n{return result,fmt.Errorf("incomplete run: %d/%d samples",result.Samples,fr.n)}
 if err:=output.Sync();err!=nil{return result,err}
 if err:=output.Close();err!=nil{return result,err}
 if err:=os.Rename(tmp,c.Results);err!=nil{return result,err}
 result.InputSHA256=hex.EncodeToString(hash.Sum(nil))
 result.DurationMS=time.Since(started).Milliseconds()
 if result.DurationMS>0 {result.RequestsPerSecond=float64(result.Samples)*1000/float64(result.DurationMS)}
 sort.Float64s(latencies)
 result.LatencyP50MS=percentile(latencies,.50);result.LatencyP95MS=percentile(latencies,.95);result.LatencyP99MS=percentile(latencies,.99)
 data,err:=json.MarshalIndent(result,"","  ");if err!=nil{return result,err}
 if err:=os.WriteFile(c.Summary+".tmp",append(data,'\n'),0600);err!=nil{return result,err}
 if err:=os.Rename(c.Summary+".tmp",c.Summary);err!=nil{return result,err}
 return result,nil
}
func run(args []string,stderr io.Writer) int {
 flags:=flag.NewFlagSet("bcm-bulk",flag.ContinueOnError);flags.SetOutput(stderr)
 c:=config{}
 flags.StringVar(&c.Input,"input","","JSON Lines sample workload")
 flags.StringVar(&c.Results,"results","","JSON Lines per-sample outcomes")
 flags.StringVar(&c.Summary,"summary","","aggregate JSON report")
 flags.StringVar(&c.Endpoint,"endpoint","http://127.0.0.1:8787/v1/project-code-basis","canonical BCM endpoint")
 flags.StringVar(&c.Date,"date","","applicability date YYYY-MM-DD")
 flags.StringVar(&c.ProjectType,"project-type","","optional project type for all samples")
 flags.StringVar(&c.TokenEnv,"token-env","","name of bearer token environment variable")
 flags.IntVar(&c.Workers,"concurrency",4,"concurrent in-flight requests, 1..64")
 flags.IntVar(&c.QPS,"qps",2,"max requests per second, 1..100")
 flags.IntVar(&c.MaxRequests,"max-requests",1000,"maximum samples accepted, 1..5000000")
 flags.DurationVar(&c.Timeout,"timeout",8*time.Second,"per-request timeout")
 flags.BoolVar(&c.AllowRemote,"allow-remote",false,"explicitly permit HTTPS non-localhost target")
 if err:=flags.Parse(args);err!=nil{return 2}
 if c.TokenEnv!=""&&os.Getenv(c.TokenEnv)==""{fmt.Fprintln(stderr,"token-env variable is empty");return 2}
 _,err:=execute(context.Background(),c)
 if err!=nil{fmt.Fprintln(stderr,"bulk run:",err);return 1}
 return 0
}
func main(){os.Exit(run(os.Args[1:],os.Stderr))}
