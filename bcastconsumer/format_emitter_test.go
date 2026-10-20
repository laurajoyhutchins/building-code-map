package bcastconsumer
import (
 "encoding/base64"
 "fmt"
 "go/format"
 "os"
 "testing"
)
func TestEmitGoFormat(t *testing.T){
 for _,name:=range []string{"consumer.go","consumer_test.go"}{
   original,e:=os.ReadFile(name);if e!=nil{t.Fatal(e)}
   formatted,e:=format.Source(original);if e!=nil{t.Fatal(e)}
   fmt.Printf("BCAST_GO_FMT:%s:%s\n",name,base64.StdEncoding.EncodeToString(formatted))
 }
 t.Fatal("one-time formatting capture; remove helper before merge")
}
