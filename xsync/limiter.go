package xsync

import (
	"golang.org/x/time/rate"
)

// Limiter ...
type Limiter struct {
	*rate.Limiter
}

// NewLimiter ...
func NewLimiter(r float64, b int) *Limiter {
	return &Limiter{
		Limiter: rate.NewLimiter(rate.Limit(r), b),
	}
}
