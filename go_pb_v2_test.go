package pbb

import (
	goapiv2 "github.com/ikaiguang/protobuf_benchmarks/go_api_v2"
	"testing"
)

func TestGoPbV2_Marshal(t *testing.T) {
	param := GoPbV2Handler.getParameter()
	buf, err := GoPbV2Handler.Marshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &goapiv2.Http{}
	err = GoPbV2Handler.Unmarshal(buf, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoPbV2_Marshal(b *testing.B) {
	param := GoPbV2Handler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoPbV2Handler.Marshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(buf)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func TestGoPbV2_JSONMarshal(t *testing.T) {
	param := GoPbV2Handler.getParameter()
	buf, err := GoPbV2Handler.JSONMarshal(param)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	res := &goapiv2.Http{}
	err = GoPbV2Handler.JSONUnmarshal(buf, res)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func BenchmarkGoPbV2_JSONMarshal(b *testing.B) {
	param := GoPbV2Handler.getParameter()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		buf, err := GoPbV2Handler.JSONMarshal(param)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(buf)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}
