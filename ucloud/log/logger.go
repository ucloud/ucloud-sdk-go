/*
Package log is the log utilities of sdk
*/
package log

import (
	"io"
	"os"

	"github.com/Sirupsen/logrus"
)

// Logger is the logger (wrapper for logrus)
type Logger struct {
	*logrus.Logger
}

// Level is the log level of logger (wrapper for logrus)
type Level logrus.Level

// Formatter is the formatter of logger (wrapper for logrus)
type Formatter logrus.Formatter

// New will return a logger pointer
func New() *Logger {
	logger := &Logger{logrus.New()}
	logger.Out = os.Stdout
	logger.Level = logrus.Level(DebugLevel)
	logger.Formatter = &logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	}
	return logger
}

// SetOutput sets the logger output.
func (logger *Logger) SetOutput(out io.Writer) {
	logger.Out = out
}

// SetFormatter sets the logger formatter.
func (logger *Logger) SetFormatter(formatter Formatter) {
	logger.Formatter = logrus.Formatter(formatter)
}

// SetLevel sets the logger level.
func (logger *Logger) SetLevel(level Level) {
	logger.Level = logrus.Level(level)
}

// GetLevel returns the logger level.
func (logger *Logger) GetLevel() Level {
	return Level(logger.Level)
}

var (
	PanicLevel = Level(logrus.PanicLevel)
	FatalLevel = Level(logrus.FatalLevel)
	ErrorLevel = Level(logrus.ErrorLevel)
	WarnLevel  = Level(logrus.WarnLevel)
	InfoLevel  = Level(logrus.InfoLevel)
	DebugLevel = Level(logrus.DebugLevel)

	SetLevel     = func(level Level) { logrus.SetLevel(logrus.Level(level)) }
	GetLevel     = func() Level { return Level(logrus.GetLevel()) }
	SetOutput    = logrus.SetOutput
	SetFormatter = logrus.SetFormatter

	WithError = logrus.WithError
	WithField = logrus.WithField

	Debug   = logrus.Debug
	Print   = logrus.Print
	Info    = logrus.Info
	Warn    = logrus.Warn
	Warning = logrus.Warning
	Error   = logrus.Error
	Panic   = logrus.Panic
	Fatal   = logrus.Fatal

	Debugf   = logrus.Debugf
	Printf   = logrus.Printf
	Infof    = logrus.Infof
	Warnf    = logrus.Warnf
	Warningf = logrus.Warningf
	Errorf   = logrus.Errorf
	Panicf   = logrus.Panicf
	Fatalf   = logrus.Fatalf
)

// Init (Deprecated) will init with level and default output (stdout) and formatter (text without color) to global logger
func Init(level Level) {
	logrus.SetLevel(logrus.Level(level))
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}
