package logplus

type Formatter interface {
	Format(*LogEntry) string
}
