package xlog

// Option ...
type Option func(*Config)

// WithFile ...
func WithFile(filepath string, stdouts ...bool) Option {
	return func(c *Config) {
		c.Filepath = filepath
		if len(stdouts) > 0 {
			c.Stdout = stdouts[0]
		}
	}
}

// WithLevel ...
func WithLevel(level Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithLogger ...
func WithLogger(logger Type) Option {
	return func(c *Config) {
		c.Type = logger
	}
}

// WithFormatter ...
func WithFormatter(formatter Formatter) Option {
	return func(c *Config) {
		c.Formatter = formatter
	}
}
