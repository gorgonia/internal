// +build debug
// +build linux

package debug

import (
	"runtime"
	"syscall"
)

// Logtid logs the thread ID of the given category.
// In release builds or debug builds that are not linux, this is a NOOP.
func Logtid(category string, logCaller int) {
	tid := syscall.Gettid()
	format := category + "- tid %v"
	if logCaller > 0 {
		pc, _, _, _ := runtime.Caller(logCaller + 1)
		format += ", called by %v"
		Logf(format, tid, runtime.FuncForPC(pc).Name())
		return
	}
	Logf(format, tid)
}
