package pbb

import (
	"bytes"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	gogoapi "github.com/ikaiguang/protobuf_benchmarks/gogo_api"
)

// GoGoHandler .
var GoGoHandler = &GoGo{}

// GoGo .
type GoGo struct{}

// getParameter .
func (g *GoGo) getParameter() *gogoapi.Http {
	rule1 := &gogoapi.HttpRule{
		Selector: "Selector1",
		Pattern: &gogoapi.HttpRule_Custom{
			Custom: &gogoapi.CustomHttpPattern{
				Kind: "Kind1",
				Path: "Path1",
			},
		},
		Body:               "Body1",
		ResponseBody:       "ResponseBody1",
		AdditionalBindings: []*gogoapi.HttpRule{},
	}
	rule2 := &gogoapi.HttpRule{
		Selector: "Selector2",
		Pattern: &gogoapi.HttpRule_Custom{
			Custom: &gogoapi.CustomHttpPattern{
				Kind: "Kind2",
				Path: "Path2",
			},
		},
		Body:               "Body2",
		ResponseBody:       "ResponseBody2",
		AdditionalBindings: []*gogoapi.HttpRule{rule1},
	}
	rules := []*gogoapi.HttpRule{rule1, rule2}
	anyData, _ := types.MarshalAny(rule2)
	return &gogoapi.Http{
		Rules:                        rules,
		FullyDecodeReservedExpansion: false,
		AnyData:                      anyData,
	}
}

// Marshal .
func (g *GoGo) Marshal(param proto.Message) ([]byte, error) {
	return proto.Marshal(param)
}

// Unmarshal .
func (g *GoGo) Unmarshal(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}

// JSONMarshal .
func (g *GoGo) JSONMarshal(param proto.Message) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: true,
	}).Marshal(buf, param)

	return buf, err
}

// JSONUnmarshal .
func (g *GoGo) JSONUnmarshal(b *bytes.Buffer, m proto.Message) error {
	return jsonpb.Unmarshal(b, m)
}
