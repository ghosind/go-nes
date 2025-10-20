package logger

type Logger interface {
	Printf(string, ...any)
}
