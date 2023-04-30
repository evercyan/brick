package xlog

// Option ...
type Option func(*Config)

// WithLogrus ...
func WithLogrus() Option {
	return func(c *Config) {
		c.Type = TypeLogrus
	}
}

// WithZap ...
func WithZap() Option {
	return func(c *Config) {
		c.Type = TypeZap
	}
}

// WithFile ...
func WithFile(filepath string) Option {
	return func(c *Config) {
		c.Filepath = filepath
	}
}

// WithStdout ...
func WithStdout(v bool) Option {
	return func(c *Config) {
		c.Stdout = v
	}
}

// WithLevel ...
func WithLevel(level Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithFormatter ...
func WithFormatter(formatter Formatter) Option {
	return func(c *Config) {
		c.Formatter = formatter
	}
}
