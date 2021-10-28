package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"netflix-auth/pkg/jwt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	authentication "netflix-auth/internal/services/auth"
	httperror "netflix-auth/pkg/http_error"
)

func Unmarshal(r io.Reader, target interface{}) httperror.HTTPError {
	body, err := io.ReadAll(r)
	if err != nil {
		return httperror.NewInternalServerErr(err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return httperror.NewBadRequestErr(err, err.Error())
	}

	return nil
}

func UnmarshalAndValidate(r io.Reader, target validation.Validatable) httperror.HTTPError {
	if err := Unmarshal(r, target); err != nil {
		return httperror.NewBadRequestErr(err, err.Error())
	}

	if err := target.Validate(); err != nil {
		return httperror.NewBadRequestErr(err, err.Error())
	}

	return nil
}

func GetID(r *http.Request) (uuid.UUID, httperror.HTTPError) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return uuid.UUID{}, httperror.NewBadRequestErr(nil, "id wasn't provided")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, httperror.NewInternalServerErr(err)
	}

	return uid, nil
}

func GetUserID(r *http.Request) (id uuid.UUID, err httperror.HTTPError) {
	userID := r.Context().Value(authentication.UserIDKey)
	id, ok := userID.(uuid.UUID)
	if !ok {
		return id, httperror.NewInternalServerErr(err)
	}

	return
}

func GetTokenClaims(r *http.Request) (*jwt.TokenClaims, httperror.HTTPError) {
	val := r.Context().Value(authentication.TokenClaimsKey)
	claims, ok := val.(*jwt.TokenClaims)
	if !ok {
		return nil, httperror.NewInternalServerErr(jwt.ErrTokenWrongType)
	}

	return claims, nil
}
