package pbb

import (
	gogofastapi "github.com/ikaiguang/protobuf_benchmarks/gogo_fast_api"
	"testing"
)

func TestGoGoFast_Marshal(t *testing.T) {
	param := GoGoFastHandler.getParameter()
	b, err := GoGoFastHandler.Marshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &gogofastapi.Http{}
	err = GoGoFastHandler.Unmarshal(b, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoGoFast_Marshal(b *testing.B) {
	param := GoGoFastHandler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoGoFastHandler.Marshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(buf)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func TestGoGoFast_JSONMarshal(t *testing.T) {
	param := GoGoFastHandler.getParameter()
	b, err := GoGoFastHandler.JSONMarshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &gogofastapi.Http{}
	err = GoGoFastHandler.JSONUnmarshal(b, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoGoFast_JSONMarshal(b *testing.B) {
	param := GoGoFastHandler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoGoFastHandler.JSONMarshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}
