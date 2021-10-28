package middlewares

import (
	"context"
	"net/http"

	authentication "netflix-auth/internal/services/auth"
	"netflix-auth/internal/services/users"
	httperror "netflix-auth/pkg/http_error"
	"netflix-auth/pkg/jwt"
	"netflix-auth/pkg/response"
)

type Auth struct {
	js jwt.Service
	us users.Service
}

func NewAuth(js jwt.Service, us users.Service) Auth {
	return Auth{
		js: js,
		us: us,
	}
}

func (a Auth) extractClaims(r *http.Request) (*jwt.TokenClaims, httperror.HTTPError) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, httperror.NewForbiddenErr(nil, "wrong access token")
	}
	claims, httpError := a.js.ParseToken(token)
	if httpError != nil {
		return nil, httpError
	}

	logOuted, httpError := a.js.IsBlacklisted(claims.Id)
	if httpError != nil {
		return nil, httpError
	}
	if logOuted {
		return nil, httperror.NewForbiddenErr(nil, "wrong access token")
	}

	return claims, nil
}

func (a Auth) WithUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, httpError := a.extractClaims(r)
		if httpError != nil {
			response.SendErrorResponse(w, httpError)

			return
		}
		_, httpError = a.us.GetByID(claims.UserID)
		if httpError != nil {
			response.SendErrorResponse(w, httpError)

			return
		}
		updatedRequest := r.WithContext(
			context.WithValue(r.Context(), authentication.UserIDKey, claims.UserID),
		)
		next.ServeHTTP(w, updatedRequest)
	})
}


func (a Auth) WithTokenClaims(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, httpError := a.extractClaims(r)
		if httpError != nil {
			response.SendErrorResponse(w, httpError)

			return
		}
		_, httpError = a.us.GetByID(claims.UserID)
		if httpError != nil {
			response.SendErrorResponse(w, httpError)

			return
		}

		updatedRequest := r.WithContext(
			context.WithValue(r.Context(), authentication.TokenClaimsKey, claims),
		)
		next.ServeHTTP(w, updatedRequest)
	})
}
