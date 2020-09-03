package godev

import "runtime/debug"

func Stacktrace() string {
	return string(debug.Stack())
}
