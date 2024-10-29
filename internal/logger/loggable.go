package logger

type Loggable interface {
	GetLogger() *Logger
	SetLogger(logger *Logger)
}
