package ratelimiter

import "context"

type RateLimiter interface {
	Allow(ctx context.Context, action string, token string) (bool, error)
}
