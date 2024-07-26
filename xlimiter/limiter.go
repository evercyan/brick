package xlimiter

import (
	"golang.org/x/time/rate"
)

// Limiter ...
type Limiter struct {
	*rate.Limiter
}

// New ...
func New(r float64, b int) *Limiter {
	return &Limiter{
		Limiter: rate.NewLimiter(rate.Limit(r), b),
	}
}
