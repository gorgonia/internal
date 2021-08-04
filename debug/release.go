// +build !debug

package debug

const DEBUG = false

var tc uint32

// var _logger_ = log.New(os.Stderr, "", 0)
// var replacement = "\n"

// EnterLoggingContext creates the next logged line with an additional tab.
// In release builds, this is a NOOOP.
func EnterLoggingContext() {}

// LeaveLoggingContext shortens tabs.
// In release builds, this is a NOOP.
func LeaveLoggingContext() {}

// Logf is a concurrent logger to stdout whose prints are serialized by a channel.
// In release builds, this is a NOOP
func Logf(format string, others ...interface{}) {}

// GetFuncName returns the name of a function. If `a` is not a function, it will panic.
// In release builds, this is a NOOP.
func GetFuncName(a interface{}) string { return "GetFuncName doesn't work in release builds." }

// LogCaller logs the caller of the function with a message.
// In release builds, this is a NOOP.
func LogCaller(format string, args ...interface{}) {}

// Logtid logs the thread ID of the given category.
// In release builds or debug builds that are not linux, this is a NOOP.
func Logtid(category string, logCaller int) {}
