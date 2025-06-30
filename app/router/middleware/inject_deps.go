package middleware

import (
	"context"
	"net/http"

	"github.com/zgack/stocks/config"
	"github.com/zgack/stocks/pkg/contextKeys"
)

func InjectDeps(cfg *config.Conf) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, contextkeys.CtxKeyConfig, cfg)
            // TODO: add logger
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
