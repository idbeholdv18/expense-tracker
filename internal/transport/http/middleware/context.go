package middleware

import "context"

type userIDKey struct{}

func SetUserID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIDKey{}, id)
}

func GetUserID(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDKey{}).(int)
	return id, ok
}
