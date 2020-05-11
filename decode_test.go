package structpbenc_test

import (
	"fmt"
	"testing"

	structpb "github.com/karantin2020/structpbenc"
	"github.com/stretchr/testify/assert"
	pb "google.golang.org/protobuf/types/known/structpb"
)

func ExampleDecode() {
	src := &pb.Struct{
		Fields: map[string]*pb.Value{
			"null":   {Kind: &pb.Value_NullValue{}},
			"number": {Kind: &pb.Value_NumberValue{NumberValue: 10}},
			"str":    {Kind: &pb.Value_StringValue{StringValue: "str"}},
			"bool":   {Kind: &pb.Value_BoolValue{BoolValue: true}},
			"struct": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
				Fields: map[string]*pb.Value{
					"nested": {Kind: &pb.Value_StringValue{StringValue: "nested"}},
				},
			}}},
			"slice": {Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{
				Values: []*pb.Value{
					{Kind: &pb.Value_StringValue{StringValue: "one"}},
					{Kind: &pb.Value_StringValue{StringValue: "two"}},
				},
			}}},
		},
	}

	dest := structpb.Decode(src)

	fmt.Println(
		dest["null"],
		dest["number"],
		dest["str"],
		dest["bool"],
	)
	if nested, ok := dest["struct"].(map[string]interface{}); ok {
		fmt.Println(nested["nested"])
	}
	if slice, ok := dest["slice"].([]interface{}); ok {
		fmt.Println(slice[0], slice[1])
	}

	// Output:
	// <nil> 10 str true
	// nested
	// one two
}

func TestDecode(t *testing.T) {
	tt := []struct {
		input *pb.Struct
		want  map[string]interface{}
	}{
		{
			input: &pb.Struct{},
			want:  nil,
		},
		{
			input: &pb.Struct{
				Fields: map[string]*pb.Value{
					"null":   {Kind: &pb.Value_NullValue{}},
					"number": {Kind: &pb.Value_NumberValue{NumberValue: 10}},
					"str":    {Kind: &pb.Value_StringValue{StringValue: "str"}},
					"bool":   {Kind: &pb.Value_BoolValue{BoolValue: true}},
					"struct": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
						Fields: map[string]*pb.Value{
							"nested": {Kind: &pb.Value_StringValue{StringValue: "nested"}},
						},
					}}},
					"slice": {Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{
						Values: []*pb.Value{
							{Kind: &pb.Value_StringValue{StringValue: "one"}},
							{Kind: &pb.Value_StringValue{StringValue: "two"}},
						},
					}}},
				},
			},
			want: map[string]interface{}{
				"null":   nil,
				"number": float64(10),
				"str":    "str",
				"bool":   true,
				"struct": map[string]interface{}{
					"nested": "nested",
				},
				"slice": []interface{}{"one", "two"},
			},
		},
	}
	for i, te := range tt {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := structpb.Decode(te.input)
			assert.Equal(t, te.want, got)
		})
	}
}
