package proxy

import (
	"sync"
	"time"
)

type State string

const (
	Closed   State = "CLOSED"
	Open     State = "OPEN"
	HalfOpen State = "HALF_OPEN"
)

type CircuitBreaker struct {
	mu sync.Mutex

	failures  int
	successes int

	state State

	failureThreshold int
	resetTimeout     time.Duration

	lastFailureTime time.Time
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: 3,
		resetTimeout:     5 * time.Second,
	}
}

func (cb *CircuitBreaker) CanRequest() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {

	case Open:
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.state = HalfOpen
			return true
		}
		return false

	case HalfOpen:
		return true

	case Closed:
		return true
	}

	return false
}

func (cb *CircuitBreaker) OnSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == HalfOpen {
		cb.successes++
		if cb.successes >= 2 {
			cb.reset()
		}
		return
	}

	cb.failures = 0
}

func (cb *CircuitBreaker) OnFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++
	cb.lastFailureTime = time.Now()

	if cb.failures >= cb.failureThreshold {
		cb.state = Open
	}
}

func (cb *CircuitBreaker) reset() {
	cb.failures = 0
	cb.successes = 0
	cb.state = Closed
}
