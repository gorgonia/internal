// +build debug

package debug

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var tc int

const DEBUG = true

var _logger_ = log.New(os.Stderr, "", 0)

var ch = make(chan func(), 64) // 64 buffer items should be long enough.

// EnterLoggingContext creates the next logged line with an additional tab.
func EnterLoggingContext() { ch <- func() { tc++ } }

// LeaveLoggingContext shortens tabs.
func LeaveLoggingContext() {
	ch <- func() {
		tc--
		if tc < 0 {
			tc = 0
		}
	}
}

// Logf is a concurrent logger whose prints are serialized by a channel.
func Logf(format string, others ...interface{}) {
	ch <- func() {
		replacement := "\n" + strings.Repeat("\t", tc)
		s := fmt.Sprintf(format, others...)
		s = strings.Replace(s, "\n", replacement, -1)
		_logger_.SetPrefix(strings.Repeat("\t", tc))
		_logger_.Println(s)
	}

}

// GetFuncName returns the name of a function. If `a` is not a function, it will panic.
func GetFuncName(a interface{}) string {
	if a == nil {
		return "nil"
	}
	return runtime.FuncForPC(reflect.ValueOf(a).Pointer()).Name()
}

// LogCaller logs the caller of a function.
func LogCaller(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	format += " Called by %s"
	args = append(args, runtime.FuncForPC(pc).Name())
	Logf(format, args...)
}

func init() {
	Logf("DEBUG MODE ON")

	go func() {
		for fn := range ch {
			fn() // wheeee side effects for the win! Who needs monads and effects systems?
		}
	}()
}
