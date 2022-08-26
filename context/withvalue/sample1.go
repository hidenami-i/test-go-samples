package main

import (
	"context"
	"errors"
	"fmt"
)

type TraceId string

const ZeroTraceId = ""

const traceKey = iota

func SetTraceId(ctx context.Context, tid TraceId) context.Context {
	return context.WithValue(ctx, traceKey, tid)
}

func GetTraceId(ctx context.Context) TraceId {
	if id, ok := ctx.Value(traceKey).(TraceId); ok {
		return id
	}
	return ZeroTraceId
}

type User struct {
	name string
	age  int
}

type contextKey string

const userKey contextKey = "userKey"

func ContextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func UserFromContext(ctx context.Context) (*User, error) {
	if user, ok := ctx.Value(userKey).(*User); ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceId(ctx))
	ctx = SetTraceId(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceId(ctx))

	fmt.Printf("user = %q\n", GetTraceId(ctx))
	ctx = ContextWithUser(ctx, &User{
		name: "abcde",
		age:  1,
	})
	user, err := UserFromContext(ctx)
	if err != nil {
		return
	}
	fmt.Printf("user = %+v\n", user)
}
