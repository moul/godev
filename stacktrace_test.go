package godev_test

import (
	"fmt"

	"moul.io/godev"
)

func ExampleStacktrace() {
	fmt.Println(godev.Stacktrace())
}
