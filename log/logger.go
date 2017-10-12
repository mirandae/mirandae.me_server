package log

import (
	"fmt"
	"os"
)

type Logger struct {
	out string
}

func NewLogger(filename string) Logger {
	return Logger{
		out: filename,
	}
}

func (l Logger) Log(messge string) error {
	file, err := os.Open("output.txt")
	if err != nil {
		// TODO wrap
		return nil
	}
	file.Close()
	return nil
}
