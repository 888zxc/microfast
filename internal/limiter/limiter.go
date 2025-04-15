package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	mu          sync.Mutex
	curRequests int
	lastTick    time.Time
	limitPerSec int
}

func NewLimiter(limit int) *Limiter {
	return &Limiter{
		limitPerSec: limit,
		lastTick:    time.Now(),
	}
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	if now.Second() != l.lastTick.Second() {
		l.curRequests = 0
		l.lastTick = now
	}
	if l.curRequests < l.limitPerSec {
		l.curRequests++
		return true
	}
	return false
}
