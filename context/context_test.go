package context_test

import (
	"context"
	"testing"

	cmnCtx "github.com/kedifeles/go-my-commons/context"
	"github.com/stretchr/testify/require"
)

func TestContext(t *testing.T) {
	ctx := cmnCtx.With(
		context.Background(),
		cmnCtx.HTTPReqID("httpReqID01"),
		cmnCtx.HTTPSessID("httpSessID02"),
		cmnCtx.SessID("sessID03"))

	require.Equal(t, "httpReqID01", ctx.Value(cmnCtx.HTTPReqIDKey))
	require.Equal(t, "httpSessID02", ctx.Value(cmnCtx.HTTPSessIDKey))
	require.Equal(t, "sessID03", ctx.Value(cmnCtx.SessIDKey))
}
