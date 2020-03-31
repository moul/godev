package godev

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// https://stackoverflow.com/questions/189562/what-is-the-proper-name-for-doing-debugging-by-adding-print-statements

var (
	logWriter  io.Writer = os.Stderr
	_, b, _, _           = runtime.Caller(0)
	basepath             = filepath.Dir(b)
)

func SetDebugOutput(writer io.Writer) {
	logWriter = writer
}

func Debug() {
	_, file, line, _ := runtime.Caller(1)
	file = strings.TrimPrefix(file, basepath+"/") // FIXME: try to keep the whole package
	fmt.Fprintf(logWriter, "[godev:debug] %s:%d\n", file, line)
}

func Sdebug() string {
	_, file, line, _ := runtime.Caller(1)
	file = strings.TrimPrefix(file, basepath+"/") // FIXME: try to keep the whole package
	return fmt.Sprintf("[godev:debug] %s:%d", file, line)
}

func Debugf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)

	_, file, line, _ := runtime.Caller(1)
	file = strings.TrimPrefix(file, basepath+"/") // FIXME: try to keep the whole package
	fmt.Fprintf(logWriter, "[godev:debug] %s:%d %v\n", file, line, message)
}

func Sdebugf(format string, args ...interface{}) string {
	message := fmt.Sprintf(format, args...)

	_, file, line, _ := runtime.Caller(1)
	file = strings.TrimPrefix(file, basepath+"/") // FIXME: try to keep the whole package
	return fmt.Sprintf("[godev:debug] %s:%d %v", file, line, message)
}
