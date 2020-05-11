package structpbenc_test

import (
	"fmt"
	"testing"

	structpb "github.com/karantin2020/structpbenc"
	"github.com/stretchr/testify/assert"
	pb "google.golang.org/protobuf/types/known/structpb"
)

func ExampleEncode() {
	src := map[string]interface{}{
		"str": "str",
		"int": 1,
		"map": map[string]interface{}{
			"nested_str": "nested_str",
			"nested_int": 10,
		},
		"struct": struct {
			Field string
		}{
			Field: "Field",
		},
		"slice": []string{"one"},
		"nil":   interface{}(nil),
	}

	var pbst pb.Struct
	structpb.Encode(src, &pbst)

	fmt.Println(
		pbst.Fields["str"].GetStringValue(),
		pbst.Fields["int"].GetNumberValue(),
		pbst.Fields["map"].GetStructValue().Fields["nested_str"].GetStringValue(),
		pbst.Fields["map"].GetStructValue().Fields["nested_int"].GetNumberValue(),
		pbst.Fields["struct"].GetStructValue().Fields["Field"].GetStringValue(),
		pbst.Fields["slice"].GetListValue().Values[0].GetStringValue(),
	)

	_, ok := pbst.Fields["nil"].GetKind().(*pb.Value_NullValue)
	fmt.Println(
		ok,
	)

	// Output:
	// str 1 nested_str 10 Field one
	// true
}

func ExampleEncodeFromStruct() {
	src := struct {
		StrField   string
		IntField   int
		unexported string
	}{
		StrField:   "str",
		IntField:   10,
		unexported: "unexported",
	}

	var pbst pb.Struct
	structpb.EncodeFromStruct(src, &pbst)

	fmt.Println(
		pbst.Fields["StrField"].GetStringValue(),
		pbst.Fields["IntField"].GetNumberValue(),
		pbst.Fields["unexported"],
	)

	// Output:
	// str 10 <nil>
}

func TestEncode(t *testing.T) {
	tt := []struct {
		input map[string]interface{}
		want  *pb.Struct
	}{
		{
			input: map[string]interface{}{},
			want:  &pb.Struct{},
		},
		{
			input: map[string]interface{}{
				"nil":      interface{}(nil),
				"int":      int(1),
				"intp":     intp(1),
				"int8":     int8(8),
				"int8p":    int8p(8),
				"int16":    int16(16),
				"int16p":   int16p(16),
				"int32":    int32(32),
				"int32p":   int32p(32),
				"int64":    int64(64),
				"int64p":   int64p(64),
				"uint":     uint(1),
				"uintp":    uintp(1),
				"uint8":    uint8(8),
				"uint8p":   uint8p(8),
				"uint16":   uint16(16),
				"uint16p":  uint16p(16),
				"uint32":   uint32(32),
				"uint32p":  uint32p(32),
				"uint64":   uint64(64),
				"uint64p":  uint64p(64),
				"float32":  float32(32),
				"float32p": float32p(32),
				"float64":  float64(64.64),
				"float64p": float64p(64.64),
				"str":      "str",
				"strp":     strp("strp"),
				"bool":     true,
				"boolp":    boolp(true),
				"struct": struct {
					ExportedField    string
					notExportedField string
					NestedStruct     struct {
						NestedField string
					}
				}{
					ExportedField:    "exported",
					notExportedField: "not_exported",
					NestedStruct: struct {
						NestedField string
					}{
						NestedField: "nested_field",
					},
				},
				"map": map[string]interface{}{
					"mapf": "mapf",
					"nested_map": map[string]interface{}{
						"nested_mapf": "nested_mapf",
					},
				},
				"arr":   [2]string{"one", "two"},
				"slice": []string{"three", "four"},
			},
			want: &pb.Struct{
				Fields: map[string]*pb.Value{
					"nil":      {Kind: &pb.Value_NullValue{}},
					"int":      {Kind: &pb.Value_NumberValue{NumberValue: 1}},
					"intp":     {Kind: &pb.Value_NumberValue{NumberValue: 1}},
					"int8":     {Kind: &pb.Value_NumberValue{NumberValue: 8}},
					"int8p":    {Kind: &pb.Value_NumberValue{NumberValue: 8}},
					"int16":    {Kind: &pb.Value_NumberValue{NumberValue: 16}},
					"int16p":   {Kind: &pb.Value_NumberValue{NumberValue: 16}},
					"int32":    {Kind: &pb.Value_NumberValue{NumberValue: 32}},
					"int32p":   {Kind: &pb.Value_NumberValue{NumberValue: 32}},
					"int64":    {Kind: &pb.Value_NumberValue{NumberValue: 64}},
					"int64p":   {Kind: &pb.Value_NumberValue{NumberValue: 64}},
					"uint":     {Kind: &pb.Value_NumberValue{NumberValue: 1}},
					"uintp":    {Kind: &pb.Value_NumberValue{NumberValue: 1}},
					"uint8":    {Kind: &pb.Value_NumberValue{NumberValue: 8}},
					"uint8p":   {Kind: &pb.Value_NumberValue{NumberValue: 8}},
					"uint16":   {Kind: &pb.Value_NumberValue{NumberValue: 16}},
					"uint16p":  {Kind: &pb.Value_NumberValue{NumberValue: 16}},
					"uint32":   {Kind: &pb.Value_NumberValue{NumberValue: 32}},
					"uint32p":  {Kind: &pb.Value_NumberValue{NumberValue: 32}},
					"uint64":   {Kind: &pb.Value_NumberValue{NumberValue: 64}},
					"uint64p":  {Kind: &pb.Value_NumberValue{NumberValue: 64}},
					"float32":  {Kind: &pb.Value_NumberValue{NumberValue: 32}},
					"float32p": {Kind: &pb.Value_NumberValue{NumberValue: 32}},
					"float64":  {Kind: &pb.Value_NumberValue{NumberValue: 64.64}},
					"float64p": {Kind: &pb.Value_NumberValue{NumberValue: 64.64}},
					"str":      {Kind: &pb.Value_StringValue{StringValue: "str"}},
					"strp":     {Kind: &pb.Value_StringValue{StringValue: "strp"}},
					"bool":     {Kind: &pb.Value_BoolValue{BoolValue: true}},
					"boolp":    {Kind: &pb.Value_BoolValue{BoolValue: true}},
					"struct": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
						Fields: map[string]*pb.Value{
							"ExportedField": {Kind: &pb.Value_StringValue{StringValue: "exported"}},
							"NestedStruct": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
								Fields: map[string]*pb.Value{
									"NestedField": {Kind: &pb.Value_StringValue{StringValue: "nested_field"}},
								},
							}}},
						},
					}}},
					"map": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
						Fields: map[string]*pb.Value{
							"mapf": {Kind: &pb.Value_StringValue{StringValue: "mapf"}},
							"nested_map": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
								Fields: map[string]*pb.Value{
									"nested_mapf": {Kind: &pb.Value_StringValue{StringValue: "nested_mapf"}},
								},
							}}},
						},
					}}},
					"arr": {Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{Values: []*pb.Value{
						{Kind: &pb.Value_StringValue{StringValue: "one"}},
						{Kind: &pb.Value_StringValue{StringValue: "two"}},
					}}}},
					"slice": {Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{Values: []*pb.Value{
						{Kind: &pb.Value_StringValue{StringValue: "three"}},
						{Kind: &pb.Value_StringValue{StringValue: "four"}},
					}}}},
				},
			},
		},
		{
			input: map[string]interface{}{
				"empty_map":    make(map[string]interface{}),
				"empty_struct": struct{}{},
				"empty_arr":    [0]string{},
				"empty_slice":  []string{},
			},
			want: &pb.Struct{
				Fields: map[string]*pb.Value{
					"empty_map":    {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{}}},
					"empty_struct": {Kind: &pb.Value_StructValue{StructValue: &pb.Struct{}}},
					"empty_arr":    {Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{}}},
					"empty_slice":  {Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{}}},
				},
			},
		},
	}
	for i, te := range tt {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := &pb.Struct{}
			if err := structpb.Encode(te.input, got); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, te.want, got)
		})
	}
}

func TestEncodeFromStruct(t *testing.T) {
	tt := []struct {
		input interface{}
		want  *pb.Struct
	}{
		{
			input: struct{}{},
			want:  &pb.Struct{},
		},
	}
	for i, te := range tt {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := &pb.Struct{}
			if err := structpb.EncodeFromStruct(te.input, got); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, te.want, got)
		})
	}
}

func intp(i int) *int {
	return &i
}

func int8p(i int8) *int8 {
	return &i
}

func int16p(i int16) *int16 {
	return &i
}

func int32p(i int32) *int32 {
	return &i
}

func int64p(i int64) *int64 {
	return &i
}

func uintp(i uint) *uint {
	return &i
}

func uint8p(i uint8) *uint8 {
	return &i
}

func uint16p(i uint16) *uint16 {
	return &i
}

func uint32p(i uint32) *uint32 {
	return &i
}

func uint64p(i uint64) *uint64 {
	return &i
}

func float32p(f float32) *float32 {
	return &f
}

func float64p(f float64) *float64 {
	return &f
}

func strp(s string) *string {
	return &s
}

func boolp(b bool) *bool {
	return &b
}
