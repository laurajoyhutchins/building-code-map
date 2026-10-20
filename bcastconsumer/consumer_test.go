package bcastconsumer

import (
 "context"
 "encoding/json"
 "net/http"
 "net/http/httptest"
 "os"
 "strings"
 "testing"
)
func example(t *testing.T) map[string]any {
 t.Helper()
 bytes,err:=os.ReadFile("testdata/package-0.2.0.json");if err!=nil{t.Fatal(err)}
 decoded,err:=decode(bytes);if err!=nil{t.Fatal(err)}
 return decoded.(map[string]any)
}
func server(t *testing.T,p map[string]any)*httptest.Server{
 t.Helper()
 id:=asString(p,"package_id")
 return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
  w.Header().Set("Content-Type","application/json")
  root:="/packages/"+id
  var body any
  switch r.URL.Path{
  case root:body=p
  case root+"/metadata":meta:=map[string]any{};for _,k:=range []string{"schema_version","package_id","package_version","publication"}{meta[k]=p[k]};body=meta
  default:
   for _,element:=range p["objects"].([]any){
    o:=element.(map[string]any);path:=root+"/objects/"+asString(o,"object_id")
    if r.URL.Path==path{body=o;break}
    if r.URL.Path==path+"/children"{
     children:=[]map[string]any{};for _,child:=range p["objects"].([]any){x:=child.(map[string]any);if x["parent_id"]==o["object_id"]{children=append(children,x)}}
     body=children;break
    }
   }
  }
  if body==nil{http.NotFound(w,r);return}
  if err:=json.NewEncoder(w).Encode(body);err!=nil{t.Error(err)}
 }))
}
func TestBCASTReferenceFixtureAndNormalHTTPNavigation(t *testing.T){
 data:=example(t);id:=asString(data,"package_id")
 s:=server(t,data);defer s.Close()
 c,e:=New(s.URL);if e!=nil{t.Fatal(e)}
 pkg,e:=c.Fetch(context.Background(),id);if e!=nil{t.Fatal(e)}
 if pkg.PublicationID==""||len(pkg.Objects)!=2{t.Fatal("incomplete BCAST fixture")}
 if err:=c.VerifyNavigation(context.Background(),pkg);err!=nil{t.Fatal(err)}
 root:=pkg.Objects[0];child:=pkg.Children(asString(root,"object_id"))
 if len(child)!=1||asString(child[0],"locator")!="1.1"{t.Fatal("source-independent navigation failed")}
 // This is content availability, not evidence that a jurisdiction adopted it.
 if asString(child[0],"status")!="resolved"{t.Fatal("unexpected availability state")}
}
func TestPackageDigestAndPublicationSubstitutionsRejected(t *testing.T){
 original:=example(t);id:=asString(original,"package_id")
 for title,modify:=range map[string]func(map[string]any){
  "package":func(p map[string]any){p["package_version"]="bad version"},
  "publication":func(p map[string]any){p["publication"].(map[string]any)["edition"]="2030"},
  "object":func(p map[string]any){p["objects"].([]any)[1].(map[string]any)["locator"]="1.2"},
  "text":func(p map[string]any){p["objects"].([]any)[1].(map[string]any)["content"].(map[string]any)["text"]="substituted"},
  "source":func(p map[string]any){p["objects"].([]any)[1].(map[string]any)["provenance"].(map[string]any)["citation"]="forged"},
 } {
  t.Run(title,func(t *testing.T){
   b,_:=json.Marshal(original);var p map[string]any;_ = json.Unmarshal(b,&p);modify(p)
   if _,err:=validate(p,id);err==nil{t.Fatal("accepted BCAST substitution")}
  })
 }
}
func TestDuplicateAndUnknownJSONRejected(t *testing.T){
 payload,_:=os.ReadFile("testdata/package-0.2.0.json")
 for _,input:=range []string{
  strings.Replace(string(payload),"\"package_id\":","\"package_id\":\"forged\",\"package_id\":",1),
  strings.Replace(string(payload),"\"schema_version\":","\"provider_secret\":\"leak\",\"schema_version\":",1),
  strings.Replace(string(payload),"\"edition\": \"2026\",","\"edition\": null,",1),
 } {
  v,err:=decode([]byte(input));if err==nil {_,err=validate(v.(map[string]any),asString(example(t),"package_id"))}
  if err==nil{t.Fatal("accepted non-contract JSON")}
 }
}
func TestBoundedEndpointAndUnsafeUnicodeEdition(t *testing.T){
 for _,s:=range []string{"file:///tmp/data","https://u:p@host","https://host/path","https://host/?key=1"}{if _,e:=New(s);e==nil{t.Fatal("accepted unsafe endpoint")}}
 p:=example(t);p["publication"].(map[string]any)["edition"]="2026\u0301"
 if _,e:=validate(p,asString(p,"package_id"));e==nil{t.Fatal("accepted Unicode identity without NFC verification")}
}
func TestUnicodeCanonicalJSONMatchesPublicPython(t *testing.T){
 got,err:=id("fixture",map[string]any{"text":"A < & > \u2028 ✓","list":[]any{"ñ","\n","\\","\""}})
 if err!=nil||got!="fixture:sha256:69fdb80ce446e0a67ed96465fec923f3772e93e365eac7dbdfbf6d5f1352b27f"{t.Fatalf("canonical mismatch %s: %v",got,err)}
}
