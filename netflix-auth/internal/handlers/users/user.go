package users

import (
	"net/http"

	"netflix-auth/internal/models"
	"netflix-auth/internal/services/users"
	"netflix-auth/pkg/response"
	"netflix-auth/pkg/utils"
)

type Handler struct {
	service users.Service
}

func NewHandler(us users.Service) Handler {
	return Handler{us}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := utils.UnmarshalAndValidate(r.Body, &input); err != nil {
		return
	}

	user, err := h.service.Create(input)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	response.SendJSONResponse(w, user)
}
