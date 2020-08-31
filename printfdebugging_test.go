package godev_test

import (
	"fmt"
	"os"

	"moul.io/godev"
)

func ExampleDebug() {
	godev.SetDebugOutput(os.Stdout)
	godev.Debug()
	godev.Debug()
	// Output:
	// [godev:debug] printfdebugging_test.go:12
	// [godev:debug] printfdebugging_test.go:13
}

func ExampleDebugf() {
	godev.SetDebugOutput(os.Stdout)
	godev.Debugf("hello: %d", 42)
	godev.Debugf("world")
	// Output:
	// [godev:debug] printfdebugging_test.go:21 hello: 42
	// [godev:debug] printfdebugging_test.go:22 world
}

func ExampleSdebug() {
	fmt.Println(godev.Sdebug())
	fmt.Println(godev.Sdebug())
	// Output:
	// [godev:debug] printfdebugging_test.go:29
	// [godev:debug] printfdebugging_test.go:30
}

func ExampleSdebugf() {
	fmt.Println(godev.Sdebugf("hello: %d", 42))
	fmt.Println(godev.Sdebugf("world"))
	// Output:
	// [godev:debug] printfdebugging_test.go:37 hello: 42
	// [godev:debug] printfdebugging_test.go:38 world
}
