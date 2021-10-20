package handlers

import (
	"net/http"
	"netflix-auth/internal/models"
	"netflix-auth/pkg/response"
	"netflix-auth/pkg/utils"
)

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := utils.UnmarshalAndValidate(r.Body, &input); err != nil {
		return
	}

	user, err := h.services.User.Create(input)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	response.SendJSONResponse(w, user)
}

func (h Handler) Auth(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := utils.UnmarshalAndValidate(r.Body, &input); err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	token, err := h.services.Auth.Authenticate(input)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	response.SendJSONResponse(w, token)
}
