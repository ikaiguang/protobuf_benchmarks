package pbb

import (
	"github.com/golang/protobuf/ptypes"
	goapiv2 "github.com/ikaiguang/protobuf_benchmarks/go_api_v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// GoPbV2Handler .
var GoPbV2Handler = &GoPbV2{}

// GoPbV2 .
type GoPbV2 struct{}

// getParameter .
func (g *GoPbV2) getParameter() *goapiv2.Http {
	rule1 := &goapiv2.HttpRule{
		Selector: "Selector1",
		Pattern: &goapiv2.HttpRule_Custom{
			Custom: &goapiv2.CustomHttpPattern{
				Kind: "Kind1",
				Path: "Path1",
			},
		},
		Body:               "Body1",
		ResponseBody:       "ResponseBody1",
		AdditionalBindings: []*goapiv2.HttpRule{},
	}
	rule2 := &goapiv2.HttpRule{
		Selector: "Selector2",
		Pattern: &goapiv2.HttpRule_Custom{
			Custom: &goapiv2.CustomHttpPattern{
				Kind: "Kind2",
				Path: "Path2",
			},
		},
		Body:               "Body2",
		ResponseBody:       "ResponseBody2",
		AdditionalBindings: []*goapiv2.HttpRule{rule1},
	}
	rules := []*goapiv2.HttpRule{rule1, rule2}
	anyData, _ := ptypes.MarshalAny(rule2)
	return &goapiv2.Http{
		Rules:                        rules,
		FullyDecodeReservedExpansion: false,
		AnyData:                      anyData,
	}
}

// Marshal .
func (g *GoPbV2) Marshal(param proto.Message) ([]byte, error) {
	return proto.Marshal(param)
}

// Unmarshal .
func (g *GoPbV2) Unmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}

// JSONMarshal .
func (g *GoPbV2) JSONMarshal(param proto.Message) ([]byte, error) {
	buf, err := (protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}).Marshal(param)

	return buf, err
}

// JSONUnmarshal .
func (g *GoPbV2) JSONUnmarshal(b []byte, m proto.Message) error {
	return protojson.Unmarshal(b, m)
}
