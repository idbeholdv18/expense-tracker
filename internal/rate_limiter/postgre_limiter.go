package ratelimiter

import (
	"context"
	"log"
	"time"
)

var (
	LoginAction         = "login"
	RegisterAction      = "register"
	ResendAction        = "resend"
	VerifyAction        = "verify"
	ResetPasswordAction = "reset_password"
)

type PostgreRateLimiter struct {
	Repo   RateLimitLogRepository
	Limit  int
	Window time.Duration
}

func (p *PostgreRateLimiter) Allow(ctx context.Context, action string, key string) (bool, error) {
	count, err := p.Repo.Count(ctx, key, action, p.Window)

	if err != nil {
		return false, err
	}

	if count >= p.Limit {
		return false, ErrReachedRateLimit
	}

	err = p.Repo.Create(ctx, &RateLimitLog{
		Key:    key,
		Action: action,
	})

	if err != nil {
		return false, nil
	}

	log.Printf("allowed: %d", count)
	return true, nil
}
