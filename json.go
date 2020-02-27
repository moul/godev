package godev

import (
	"encoding/json"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

func JSON(input interface{}) string {
	out, _ := json.Marshal(input)
	return string(out)
}

func PrettyJSON(input interface{}) string {
	out, _ := json.MarshalIndent(input, "", "  ")
	return string(out)
}

func JSONPB(input proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, _ := marshaler.MarshalToString(input)
	return out
}

func PrettyJSONPB(input proto.Message) string {
	marshaler := jsonpb.Marshaler{Indent: "  "}
	out, _ := marshaler.MarshalToString(input)
	return out
}
