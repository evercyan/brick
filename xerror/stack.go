package xerror

import "runtime"

// callers returns a stack trace. the argument skip is the number of stack frames to skip before recording
// in pc, with 0 identifying the frame for Callers itself and 1 identifying the caller of Callers.
func callers(skip int) *stack {
	const depth = 64
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0 : n-2] // todo: change this to filtering out runtime instead of hardcoding n-2
	return &st
}

// stack is an array of program counters.
type stack []uintptr
