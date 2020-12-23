package pbb

import (
	gogoapi "github.com/ikaiguang/protobuf_benchmarks/gogo_api"
	"testing"
)

func TestGoGo_Marshal(t *testing.T) {
	param := GoGoHandler.getParameter()
	b, err := GoGoHandler.Marshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &gogoapi.Http{}
	err = GoGoHandler.Unmarshal(b, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoGo_Marshal(b *testing.B) {
	param := GoGoHandler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoGoHandler.Marshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(buf)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func TestGoGo_JSONMarshal(t *testing.T) {
	param := GoGoHandler.getParameter()
	b, err := GoGoHandler.JSONMarshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &gogoapi.Http{}
	err = GoGoHandler.JSONUnmarshal(b, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoGo_JSONMarshal(b *testing.B) {
	param := GoGoHandler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoGoHandler.JSONMarshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}
