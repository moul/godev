package godev_test

import (
	"fmt"

	"moul.io/godev"
)

func ExampleYAML() {
	fmt.Println(godev.YAML([]string{"hello", "world"}))
	fmt.Println(godev.YAML(42))
	fmt.Println(godev.YAML(nil))
	// Output:
	// - hello
	// - world
	// 42
	// null
}
