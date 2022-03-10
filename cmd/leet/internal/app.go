package internal

type app struct {
	*Config
	*LeetCode
	*Generator
}

// NewApp ...
func NewApp() *app {
	return &app{
		Config:    newConfig(),
		LeetCode:  newLeetCode(),
		Generator: newGenerator(),
	}
}
