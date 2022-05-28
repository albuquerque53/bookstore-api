package middleware

import (
	"bookstoreapi/internal/application/helper"
	"net/http"
)

func PanicHandler(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				helper.HandleHttpError(w, 500, "unexpected error", nil)
			}
		}()

		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
