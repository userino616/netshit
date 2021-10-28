package middlewares

import (
	"net/http"
	"netflix-auth/pkg/logger"

	httperror "netflix-auth/pkg/http_error"
)

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				l := logger.GetLogger()
				l.Error(err)
				http.Error(w, httperror.InternalServerError, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
