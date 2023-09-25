package middlewares

import (
	"context"
	"errors"
	"graphql/token"
	"net/http"
)

type TokenMakerKeyType string

var TokenMakerKey TokenMakerKeyType = "token-maker"

func TokenMakerMiddleware(tm *token.PasetoMaker, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqWithStore := r.WithContext(context.WithValue(r.Context(), TokenMakerKey, tm))
		next.ServeHTTP(w, reqWithStore)
	})
}

func GetTokenMakerFromContext(ctx context.Context) (*token.PasetoMaker, error) {
	obj, ok := ctx.Value(TokenMakerKey).(*token.PasetoMaker)
	if !ok {
		return nil, errors.New("unable to get token maker from context")
	}
	return obj, nil
}
