package godev

import "fmt"

func ExamplePrettyJSON() {
	fmt.Println(PrettyJSON([]string{"hello", "world"}))
	fmt.Println(PrettyJSON(42))
	fmt.Println(PrettyJSON(nil))
	// Output:
	// [
	//   "hello",
	//   "world"
	// ]
	// 42
	// null
}

func ExampleJSON() {
	fmt.Println(JSON([]string{"hello", "world"}))
	fmt.Println(JSON(42))
	fmt.Println(JSON(nil))
	// Output:
	// ["hello","world"]
	// 42
	// null
}
