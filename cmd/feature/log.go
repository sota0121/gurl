package feature

import "fmt"

// LogLevel is the level of the log.
type LogLevel int

const (
	Invalid LogLevel = iota
	Silent
	Normal
	Verbose
)

type LeveledLogger struct {
	// contains filtered or unexported fields
	Level LogLevel
}

func NewLeveledLogger(level LogLevel) *LeveledLogger {
	return &LeveledLogger{
		Level: level,
	}
}

// Println prints the log with the specified level.
// this method is the wrapper of fmt.Println.
func (l *LeveledLogger) Println(verbose bool, v ...interface{}) {
	if verbose {
		if l.Level >= Verbose {
			fmt.Println(v...)
		}
	} else {
		if l.Level >= Normal {
			fmt.Println(v...)
		}
	}
}
