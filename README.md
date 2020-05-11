# structpbenc

Encode map (or struct) to protobuf Struct, and decode from protobuf Struct to map

## Installation

```bash
go get "github.com/karantin2020/structpbenc"
```

## Examples
### Encode to protobuf Struct

```go
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
```

### Decode to map

```go
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
```
