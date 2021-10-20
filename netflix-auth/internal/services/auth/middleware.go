package authentication

import (
	"context"
	"net/http"
	"netflix-auth/pkg/response"

	httperror "netflix-auth/pkg/http_error"
)

type usedID string

const (
	UserIDKey usedID = "userID"
)

func (s authService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			response.SendErrorResponse(w, httperror.NewForbiddenErr(nil, "wrong access token"))

			return
		}
		userID, httpError := s.jwtService.ParseToken(token)
		if httpError != nil {
			response.SendErrorResponse(w, httpError)

			return
		}
		_, httpError = s.GetUserByID(userID)
		if httpError != nil {
			response.SendErrorResponse(w, httpError)

			return
		}

		updatedRequest := r.WithContext(
			context.WithValue(r.Context(), UserIDKey, userID),
		)
		next.ServeHTTP(w, updatedRequest)
	})
}
