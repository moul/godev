package godev

import "fmt"

func ExampleYAML() {
	fmt.Println(YAML([]string{"hello", "world"}))
	fmt.Println(YAML(42))
	fmt.Println(YAML(nil))
	// Output:
	// - hello
	// - world
	// 42
	// null
}
