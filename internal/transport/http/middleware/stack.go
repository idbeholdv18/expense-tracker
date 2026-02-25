package middleware

import (
	httptransport "github/idbeholdv18/expense-tracker/internal/transport/http"
)

func Apply(handler httptransport.AppHandler, middlewares ...Middleware) httptransport.AppHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
