package godev_test

import (
	"fmt"

	pb "github.com/gogo/protobuf/jsonpb/jsonpb_test_proto"
	"moul.io/godev"
)

func ExamplePrettyJSON() {
	fmt.Println(godev.PrettyJSON([]string{"hello", "world"}))
	fmt.Println(godev.PrettyJSON(42))
	fmt.Println(godev.PrettyJSON(nil))
	// Output:
	// [
	//   "hello",
	//   "world"
	// ]
	// 42
	// null
}

func ExampleJSON() {
	fmt.Println(godev.JSON([]string{"hello", "world"}))
	fmt.Println(godev.JSON(42))
	fmt.Println(godev.JSON(nil))
	// Output:
	// ["hello","world"]
	// 42
	// null
}

func ExampleJSONPB() {
	msg := pb.Mappy{Enumy: map[string]pb.Numeral{"XIV": pb.Numeral_ROMAN}}
	fmt.Println(godev.JSONPB(&msg))
	// Output:
	// {"enumy":{"XIV":"ROMAN"}}
}

func ExamplePrettyJSONPB() {
	msg := pb.Mappy{Enumy: map[string]pb.Numeral{"XIV": pb.Numeral_ROMAN}}
	fmt.Println(godev.PrettyJSONPB(&msg))
	// Output:
	// {
	//   "enumy": {
	//     "XIV": "ROMAN"
	//   }
	// }
}
