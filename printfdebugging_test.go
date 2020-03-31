package godev

import (
	"fmt"
	"os"
)

func ExampleDebug() {
	SetDebugOutput(os.Stdout)
	Debug()
	Debug()
	// Output:
	// [godev:debug] printfdebugging_test.go:10
	// [godev:debug] printfdebugging_test.go:11
}

func ExampleDebugf() {
	SetDebugOutput(os.Stdout)
	Debugf("hello: %d", 42)
	Debugf("world")
	// Output:
	// [godev:debug] printfdebugging_test.go:19 hello: 42
	// [godev:debug] printfdebugging_test.go:20 world
}

func ExampleSdebug() {
	fmt.Println(Sdebug())
	fmt.Println(Sdebug())
	// Output:
	// [godev:debug] printfdebugging_test.go:27
	// [godev:debug] printfdebugging_test.go:28
}

func ExampleSdebugf() {
	fmt.Println(Sdebugf("hello: %d", 42))
	fmt.Println(Sdebugf("world"))
	// Output:
	// [godev:debug] printfdebugging_test.go:35 hello: 42
	// [godev:debug] printfdebugging_test.go:36 world
}
