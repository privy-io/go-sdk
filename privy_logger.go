package privyclient

import (
	"fmt"
	"io"
	"os"
)

// LogLevel represents the severity level for logging
type LogLevel int

const (
	// LogLevelNone disables all logging (default)
	LogLevelNone LogLevel = iota
	// LogLevelError shows only error messages
	LogLevelError
	// LogLevelInfo shows error and info messages
	LogLevelInfo
	// LogLevelDebug shows error, info, and debug messages
	LogLevelDebug
	// LogLevelVerbose shows all messages (error, info, debug, verbose)
	LogLevelVerbose
)

// String returns the string representation of a LogLevel
func (l LogLevel) String() string {
	switch l {
	case LogLevelNone:
		return "NONE"
	case LogLevelError:
		return "ERROR"
	case LogLevelInfo:
		return "INFO"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelVerbose:
		return "VERBOSE"
	default:
		return "UNKNOWN"
	}
}

// logger defines the interface for logging operations in the SDK
type logger interface {
	// SetLevel sets the minimum log level to display
	SetLevel(level LogLevel)
	// Error logs an error message (always displayed unless LogLevelNone)
	Error(msg string)
	// Info logs an informational message
	Info(msg string)
	// Debug logs a debug message
	Debug(msg string)
	// Verbose logs a verbose message (most detailed)
	Verbose(msg string)
}

// privyLogger is a basic implementation of the Logger interface
// that outputs formatted messages to stdout
type privyLogger struct {
	level  LogLevel
	writer io.Writer
}

// newPrivyLogger creates a new PrivyLogger instance with the specified log level.
// If level is not explicitly set (zero value), it defaults to LogLevelNone (no logging).
func newPrivyLogger(level LogLevel) *privyLogger {
	return &privyLogger{
		level:  level,
		writer: os.Stdout,
	}
}

// SetLevel sets the minimum log level for the logger.
// Note: This should only be called during initialization, not concurrently.
func (l *privyLogger) SetLevel(level LogLevel) {
	l.level = level
}

// Error logs an error message if log level is Error or higher
func (l *privyLogger) Error(msg string) {
	l.log(LogLevelError, msg)
}

// Info logs an informational message if log level is Info or higher
func (l *privyLogger) Info(msg string) {
	l.log(LogLevelInfo, msg)
}

// Debug logs a debug message if log level is Debug or higher
func (l *privyLogger) Debug(msg string) {
	l.log(LogLevelDebug, msg)
}

// Verbose logs a verbose message if log level is Verbose
func (l *privyLogger) Verbose(msg string) {
	l.log(LogLevelVerbose, msg)
}

// log handles the actual logging with hierarchical level checking
func (l *privyLogger) log(level LogLevel, msg string) {
	// Don't log if current level doesn't meet or exceed the message level
	if l.level < level {
		return
	}

	// Format: [Privy][LEVEL] message
	formattedMsg := fmt.Sprintf("[Privy][%s] %s\n", level.String(), msg)
	l.writer.Write([]byte(formattedMsg))
}
