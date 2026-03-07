package ratelimiter

import (
	"context"
	"time"
)

type RateLimitLog struct {
	ID        int
	Key       string
	Action    string
	CreatedAt time.Time
}

type RateLimitLogRepository interface {
	Count(ctx context.Context, key string, action string, window time.Duration) (int, error)
	Create(ctx context.Context, rll *RateLimitLog) error
}
