package middleware

import httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"

type Middleware func(next httptransport.AppHandler) httptransport.AppHandler
