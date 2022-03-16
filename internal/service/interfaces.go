package service

// ILogger -
type ILogger interface {
	Log(message string)
	WithPrefix(prefix string)
}
