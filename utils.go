// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package babycli

import (
	"fmt"
	"io"
)

func panicf(msg string, args ...any) {
	s := fmt.Sprintf(msg, args...)
	s = "babycli: " + s
	panic(s)
}

func write(output io.Writer, msg string) {
	_, _ = io.WriteString(output, msg)
	_, _ = io.WriteString(output, "\n")
}

func writef(output io.Writer, msg string, args ...any) {
	s := fmt.Sprintf(msg, args...)
	s += "\n"
	_, _ = io.WriteString(output, s)
}
