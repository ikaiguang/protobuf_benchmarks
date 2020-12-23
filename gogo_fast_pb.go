package pbb

import (
	"bytes"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	gogofastapi "github.com/ikaiguang/protobuf_benchmarks/gogo_api"
)

// GoGoFastHandler
var GoGoFastHandler = &GoGoFast{}

// GoGoFast .
type GoGoFast struct{}

// getParameter .
func (g *GoGoFast) getParameter() *gogofastapi.Http {
	rule1 := &gogofastapi.HttpRule{
		Selector: "Selector1",
		Pattern: &gogofastapi.HttpRule_Custom{
			Custom: &gogofastapi.CustomHttpPattern{
				Kind: "Kind1",
				Path: "Path1",
			},
		},
		Body:               "Body1",
		ResponseBody:       "ResponseBody1",
		AdditionalBindings: []*gogofastapi.HttpRule{},
	}
	rule2 := &gogofastapi.HttpRule{
		Selector: "Selector2",
		Pattern: &gogofastapi.HttpRule_Custom{
			Custom: &gogofastapi.CustomHttpPattern{
				Kind: "Kind2",
				Path: "Path2",
			},
		},
		Body:               "Body2",
		ResponseBody:       "ResponseBody2",
		AdditionalBindings: []*gogofastapi.HttpRule{rule1},
	}
	rules := []*gogofastapi.HttpRule{rule1, rule2}
	anyData, _ := types.MarshalAny(rule2)
	return &gogofastapi.Http{
		Rules:                        rules,
		FullyDecodeReservedExpansion: false,
		AnyData:                      anyData,
	}
}

// Marshal .
func (g *GoGoFast) Marshal(param proto.Message) ([]byte, error) {
	return proto.Marshal(param)
}

// Unmarshal .
func (g *GoGoFast) Unmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}

// JSONMarshal .
func (g *GoGoFast) JSONMarshal(param proto.Message) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: true,
	}).Marshal(buf, param)

	return buf, err
}

// JSONUnmarshal .
func (g *GoGoFast) JSONUnmarshal(b *bytes.Buffer, m proto.Message) error {
	return jsonpb.Unmarshal(b, m)
}
