package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     float64
	refillRate float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     float64(capacity),
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) refill() {

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()

	tb.tokens += elapsed * tb.refillRate

	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}

	tb.lastRefill = now
}

func (tb *TokenBucket) Allow() bool {

	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}

	return false
}

var (
	buckets = make(map[string]*TokenBucket)
	mu      sync.Mutex
)

func GetBucket(ip string) *TokenBucket {

	mu.Lock()
	defer mu.Unlock()

	bucket, exists := buckets[ip]

	if !exists {
		bucket = NewTokenBucket(10, 5) // burst 10, refill 5/sec
		buckets[ip] = bucket
	}

	return bucket
}