package godev

import "os"

func ExampleDebugf() {
	SetDebugOutput(os.Stdout)
	Debugf("hello: %d", 42)
	Debugf("world")
	// Output:
	// [godev:debug] printfdebugging_test.go:7 hello: 42
	// [godev:debug] printfdebugging_test.go:8 world
}
