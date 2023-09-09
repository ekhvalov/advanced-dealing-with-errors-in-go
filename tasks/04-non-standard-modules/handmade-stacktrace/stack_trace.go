package stacktrace

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const maxStacktraceDepth = 32

type Frame uintptr

func (f Frame) pc() uintptr {
	return uintptr(f) - 1
}

func (f Frame) String() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, line := fn.FileLine(f.pc())
	return fmt.Sprintf("%s\n%s:%d", prettifyFuncName(fn.Name()), prettifyFileName(file), line)
}

type StackTrace []Frame

func (s StackTrace) String() string {
	sb := strings.Builder{}
	for _, frame := range s {
		sb.WriteString(frame.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Trace возвращает стектрейс глубиной не более maxStacktraceDepth.
// Возвращаемый стектрейс начинается с того места, где была вызвана Trace.
func Trace() StackTrace {
	pc := make([]uintptr, maxStacktraceDepth)
	l := runtime.Callers(2, pc)
	st := make([]Frame, l)
	for i := 0; i < l; i++ {
		st[i] = Frame(pc[i])
	}
	return st
}

func prettifyFuncName(name string) string {
	return keepFromTailToByte(name, filepath.Separator, 0)
}

func prettifyFileName(name string) string {
	return keepFromTailToByte(name, filepath.Separator, 1)
}

func keepFromTailToByte(name string, b byte, skip int) string {
	i := lastIndexByte(name, b, skip)
	if i < 0 {
		return name
	}
	return name[i+1:]
}

func lastIndexByte(s string, b byte, skip int) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == b {
			skip--
			if skip < 0 {
				return i
			}
		}
	}
	return -1
}
