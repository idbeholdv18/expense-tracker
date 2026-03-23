package ratelimiter

import (
	"context"
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

func (p *PostgreRateLimiter) Allow(ctx context.Context, key string) (bool, error) {
	count, err := p.Repo.Count(ctx, key, p.Window)

	if err != nil {
		return false, err
	}

	if count >= p.Limit {
		return false, ErrReachedRateLimit
	}

	err = p.Repo.Create(ctx, &RateLimitLog{
		Key: key,
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}
