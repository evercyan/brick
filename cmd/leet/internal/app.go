package internal

// App ...
type App struct {
	*Config
	*LeetCode
	*Generator
}

// NewApp ...
func NewApp() *App {
	return &App{
		Config:    newConfig(),
		LeetCode:  newLeetCode(),
		Generator: newGenerator(),
	}
}
