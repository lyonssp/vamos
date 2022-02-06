package vamos

import (
	"flag"
	"fmt"
)

var (
	seedFlag = flag.Int64("vamos.seed", 0, "seed to use to reproduce observed behaviors")

	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type Property[O any] struct {
	Description string
	Generator   Generator[O]
	Check       func(*V[O])
}

func green(s string) string {
	return fmt.Sprintln(colorGreen, s, colorReset)
}

func red(s string) string {
	return fmt.Sprintln(colorRed, s, colorReset)
}
