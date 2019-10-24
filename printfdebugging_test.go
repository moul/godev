package godev

import "os"

func ExampleDebug() {
	SetDebugOutput(os.Stdout)
	Debug()
	Debug()
	// Output:
	// [godev:debug] printfdebugging_test.go:7
	// [godev:debug] printfdebugging_test.go:8
}

func ExampleDebugf() {
	SetDebugOutput(os.Stdout)
	Debugf("hello: %d", 42)
	Debugf("world")
	// Output:
	// [godev:debug] printfdebugging_test.go:16 hello: 42
	// [godev:debug] printfdebugging_test.go:17 world
}
