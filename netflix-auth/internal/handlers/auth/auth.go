package auth

import (
	"net/http"
	"netflix-auth/internal/models"
	"netflix-auth/internal/services/auth"
	"netflix-auth/pkg/response"
	"netflix-auth/pkg/utils"
)

type Handler struct {
	service auth.Service
}

func NewHandler(as auth.Service) Handler {
	return Handler{as}
}

func (h Handler) Auth(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := utils.UnmarshalAndValidate(r.Body, &input); err != nil {
		response.SendErrorResponse(w, err)
		return
	}

	token, err := h.service.Authenticate(input)
	if err != nil {
		response.SendErrorResponse(w, err)
		return
	}

	response.SendJSONResponse(w, token)
}

func (h Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	tokenClaims, httpError := utils.GetTokenClaims(r)
	if httpError != nil {
		response.SendErrorResponse(w, httpError)

		return
	}
	err := h.service.LogOut(tokenClaims.Id, tokenClaims.ExpiresAt)
	if err != nil {
		response.SendErrorResponse(w, err)
		return
	}
	w.WriteHeader(200)
}
