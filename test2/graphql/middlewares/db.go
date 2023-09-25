package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseKeyType string

var DatabaseKey DatabaseKeyType = "db"

func DatabaseMiddleware(db *pgxpool.Pool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqWithStore := r.WithContext(context.WithValue(r.Context(), DatabaseKey, db))
		next.ServeHTTP(w, reqWithStore)
	})
}

func GetDatabaseFromContext(ctx context.Context) (*pgxpool.Pool, error) {
	obj, ok := ctx.Value(DatabaseKey).(*pgxpool.Pool)
	if !ok {
		return nil, errors.New("database not found")
	}
	return obj, nil
}
