package domain

import (
	"log"
	"testing"
)

var data = []byte(`
{
   "time":1593709730877594291,
   "namespace":"prod",
   "source":"app_1",
   "host":"127.100.0.1:50000",
   "level":"debug",
   "trace_id":"1c7fuhpo0ln2dcq",
   "message":"read failed: some message 4",
   "build_commit":"db957a22b3c1d6e508c0828917a5e14c572fb007",
   "config_hash":"130362a6fd10cf2f939dd0cfc0ab222cee6a99ec",
   "key1":[
      1,
      2,
      3,
      4,
      5,
      6
   ],
   "key2":[
      "a",
      "b",
      "c",
      "d",
      "e",
      "f"
   ],
   "key3":{
      "key4":{
         "key5":[
            11,
            11,
            11,
            "WWW"
         ]
      }
   }
}`)

func TestEntry_UnmarshalJSON(t *testing.T) {
	entry := &Entry{}

	if err := entry.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	log.Println(entry.Time)
	log.Println(entry.Namespace)
	log.Println(entry.Source)
	log.Println(entry.Host)
	log.Println(entry.Level)
	log.Println(entry.TraceID)
	log.Println(entry.Message)
	log.Println(entry.BuildCommit)
	log.Println(entry.ConfigHash)
	log.Println(entry.Params)
	log.Println(entry.StringKey)
	log.Println(entry.StringVal)
	log.Println(entry.FloatKey)
	log.Println(entry.FloatVal)
}

func BenchmarkEntry_UnmarshalJSON(b *testing.B) {
	entry := &Entry{}
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := entry.UnmarshalJSON(data); err != nil {
			b.Fatal(err)
		}
	}
}
