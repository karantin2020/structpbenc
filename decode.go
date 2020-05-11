package structpbenc

import (
	"fmt"

	pb "google.golang.org/protobuf/types/known/structpb"
)

// Decode decodes pb.Struct to map[string]interface{}
func Decode(pbst *pb.Struct) map[string]interface{} {
	if pbst == nil || len(pbst.Fields) == 0 {
		return nil
	}
	m := make(map[string]interface{})
	for k, pbv := range pbst.Fields {
		m[k] = DecodeValue(pbv)
	}
	return m
}

// DecodeValue decodes pb.Value to interface{}
func DecodeValue(pbv *pb.Value) interface{} {
	switch kind := pbv.GetKind().(type) {
	case *pb.Value_NullValue:
		return nil
	case *pb.Value_NumberValue:
		return kind.NumberValue
	case *pb.Value_StringValue:
		return kind.StringValue
	case *pb.Value_BoolValue:
		return kind.BoolValue
	case *pb.Value_StructValue:
		return Decode(kind.StructValue)
	case *pb.Value_ListValue:
		ret := make([]interface{}, len(kind.ListValue.Values))
		for i, e := range kind.ListValue.Values {
			ret[i] = DecodeValue(e)
		}
		return ret
	default:
		panic(fmt.Sprintf("structpb: failed to decode *pb.Value. unexpected kind (%T)%v", kind, kind))
	}
}
