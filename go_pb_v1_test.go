package pbb

import (
	goapiv1 "github.com/ikaiguang/protobuf_benchmarks/go_api_v1"
	"testing"
)

func TestGoPbV1_Marshal(t *testing.T) {
	param := GoPbV1Handler.getParameter()
	buf, err := GoPbV1Handler.Marshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &goapiv1.Http{}
	err = GoPbV1Handler.Unmarshal(buf, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoPbV1_Marshal(b *testing.B) {
	param := GoPbV1Handler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoPbV1Handler.Marshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(buf)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func TestGoPbV1_JSONMarshal(t *testing.T) {
	param := GoPbV1Handler.getParameter()
	buf, err := GoPbV1Handler.JSONMarshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &goapiv1.Http{}
	err = GoPbV1Handler.JSONUnmarshal(buf, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoPbV1_JSONMarshal(b *testing.B) {
	param := GoPbV1Handler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoPbV1Handler.JSONMarshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}
