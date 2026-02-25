package expensetypes

import "time"

type ExpenseTypeCreateResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
