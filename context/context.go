package context

import "context"

// Represent context key for storing information in context
const (
	HTTPReqIDKey ctxKey = iota
	HTTPSessIDKey
	SessIDKey
)

type ctxKey int

func (c ctxKey) String() string {
	switch c {
	case HTTPReqIDKey:
		return "httpReqID"
	case HTTPSessIDKey:
		return "httpSessID"
	case SessIDKey:
		return "sessID"
	}
	return ""
}

// BuildCtxFunc is special type to build a new context
type BuildCtxFunc func(ctx context.Context) context.Context

// With help to build a new context fluently
func With(ctx context.Context, buildFuncs ...BuildCtxFunc) context.Context {
	for _, buildCtxFunc := range buildFuncs {
		ctx = buildCtxFunc(ctx)
	}
	return ctx
}

// HTTPReqID compose http request ID information in context using With function
func HTTPReqID(httpReqID string) BuildCtxFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, HTTPReqIDKey, httpReqID)
	}
}

// HTTPSessID compose http session ID information in context using With function
func HTTPSessID(httpSessID string) BuildCtxFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, HTTPSessIDKey, httpSessID)
	}
}

// SessID compose BCA session ID information in context using With function
func SessID(bcaSessID string) BuildCtxFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, SessIDKey, bcaSessID)
	}
}
