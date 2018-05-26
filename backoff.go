package lingo

import (
	"errors"
	"time"
)

type backoffConfig struct {
	retries uint
	attempt uint

	base uint
	cap  uint
}

// NewBackoffConfig returns a new backoffConfig struct.
func NewBackoffConfig(retries, base, cap uint) backoffConfig {
	return backoffConfig{
		retries: retries,
		base:    base,
		cap:     cap,
	}
}

func (c *backoffConfig) Retry() error {
	c.attempt++
	if c.attempt > c.retries {
		return errors.New("Backoff failed")
	}

	minimum := min(c.cap, exp(c.base*2, c.attempt))
	time.Sleep(time.Duration(minimum) * time.Millisecond)

	return nil
}

// Reset resets the backoff attempt counter.
func (c *backoffConfig) Reset() {
	c.attempt = 0
}

func exp(x, n uint) uint {
	if n == 0 {
		return 1
	}

	y := uint(1)
	for n > 1 {
		if n%2 == 0 {
			x = x * x
			n = n / 2
		} else {
			y = x * y
			x = x * x
			n = (n - 1) / 2
		}
	}

	return x * y
}

func min(x, y uint) uint {
	if x <= y {
		return x
	}

	return y
}
