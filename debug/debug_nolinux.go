// +build debug
// +build !linux

package debug

// Logtid logs the thread ID of the given category.
// In release builds or debug builds that are not linux, this is a NOOP.
func Logtid(category string, logCaller int) {}
