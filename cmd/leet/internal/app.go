package internal

type app struct {
	*Config
	*Leet
	*File
}

// NewService ...
func NewApp() *app {
	return &app{
		Config: newConfig(),
		Leet:   newLeet(),
		File:   newFile(),
	}
}
