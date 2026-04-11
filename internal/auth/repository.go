package auth

import (
	"context"
	"time"
)

type RegistrationStatus string

const (
	StatusPending   RegistrationStatus = "pending"
	StatusConfirmed RegistrationStatus = "confirmed"
	StatusExpired   RegistrationStatus = "expired"
	StatusCancelled RegistrationStatus = "cancelled"
)

type RegistrationRequest struct {
	Id           int
	Email        string
	Username     string
	TokenHash    string
	PasswordHash string
	Status       RegistrationStatus
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RegistrationRequestRepository interface {
	GetById(ctx context.Context, id int) (*RegistrationRequest, error)
	UpdateStatusById(ctx context.Context, id int, status RegistrationStatus) error
	Create(ctx context.Context, rr *RegistrationRequest) error
}
