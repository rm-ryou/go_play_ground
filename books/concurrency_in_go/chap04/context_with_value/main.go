package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserId ctxKey = iota
	ctxAuthToken
)

func UserId(ctx context.Context) string {
	return ctx.Value(ctxUserId).(string)
}

func AuthToken(ctx context.Context) string {
	return ctx.Value(ctxAuthToken).(string)
}

func ProcessRequest(userId, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserId, userId)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (auth: %v)",
		UserId(ctx),
		AuthToken(ctx),
	)
}

func main() {
	ProcessRequest("jane", "abc123")
}
