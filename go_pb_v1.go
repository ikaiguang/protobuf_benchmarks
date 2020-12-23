package pbb

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	goapiv1 "github.com/ikaiguang/protobuf_benchmarks/go_api_v1"
)

// GoPbV1Handler .
var GoPbV1Handler = &GoPbV1{}

// GoPbV1 .
type GoPbV1 struct{}

// getParameter .
func (g *GoPbV1) getParameter() *goapiv1.Http {
	rule1 := &goapiv1.HttpRule{
		Selector: "Selector1",
		Pattern: &goapiv1.HttpRule_Custom{
			Custom: &goapiv1.CustomHttpPattern{
				Kind: "Kind1",
				Path: "Path1",
			},
		},
		Body:               "Body1",
		ResponseBody:       "ResponseBody1",
		AdditionalBindings: []*goapiv1.HttpRule{},
	}
	rule2 := &goapiv1.HttpRule{
		Selector: "Selector2",
		Pattern: &goapiv1.HttpRule_Custom{
			Custom: &goapiv1.CustomHttpPattern{
				Kind: "Kind2",
				Path: "Path2",
			},
		},
		Body:               "Body2",
		ResponseBody:       "ResponseBody2",
		AdditionalBindings: []*goapiv1.HttpRule{rule1},
	}
	rules := []*goapiv1.HttpRule{rule1, rule2}
	anyData, _ := ptypes.MarshalAny(rule2)
	return &goapiv1.Http{
		Rules:                        rules,
		FullyDecodeReservedExpansion: false,
		AnyData:                      anyData,
	}
}

// Marshal .
func (g *GoPbV1) Marshal(param proto.Message) ([]byte, error) {
	return proto.Marshal(param)
}

// Unmarshal .
func (g *GoPbV1) Unmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}

// JSONMarshal .
func (g *GoPbV1) JSONMarshal(param proto.Message) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: true,
	}).Marshal(buf, param)

	return buf, err
}

// JSONUnmarshal .
func (g *GoPbV1) JSONUnmarshal(b *bytes.Buffer, m proto.Message) error {
	return jsonpb.Unmarshal(b, m)
}
