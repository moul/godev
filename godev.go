package godev

import (
	"encoding/json"
)

func JSON(input interface{}) string {
	out, _ := json.Marshal(input)
	return string(out)
}

func PrettyJSON(input interface{}) string {
	out, _ := json.MarshalIndent(input, "", "  ")
	return string(out)
}
