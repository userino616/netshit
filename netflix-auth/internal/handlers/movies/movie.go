package movies

import (
	"net/http"
	"netflix-auth/internal/services/movies"
	"netflix-auth/pkg/response"
	"netflix-auth/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/userino616/netflix-grpc/movieservice"
)

type Handler struct {
	service movies.Service
}

func NewHandler(ms movies.Service) Handler {
	return Handler{ms}
}

func (h Handler) Search(w http.ResponseWriter, r *http.Request) {
	movieName := mux.Vars(r)["name"]
	moviesList, err := h.service.Search(&movieservice.SearchMovieRequest{Name: movieName})
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	response.SendJSONResponse(w, moviesList)
}

func (h Handler) AddBookmark(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	movieID, err := utils.GetID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	bookmark := &movieservice.AddBookmarkRequest{
		UserId:  userID.String(),
		MovieId: movieID.String(),
	}
	err = h.service.AddBookmark(bookmark)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) AddToWatchedList(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	movieID, err := utils.GetID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	watchedMovie := &movieservice.AddToWatchedListRequest{
		UserId:  userID.String(),
		MovieId: movieID.String(),
	}
	err = h.service.AddToWatchedList(watchedMovie)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetUserWatchedList(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	moviesList, err := h.service.GetWatchedList(userID)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	response.SendJSONResponse(w, moviesList)
}

func (h Handler) GetUserBookmarks(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	moviesList, err := h.service.GetBookmarks(userID)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	response.SendJSONResponse(w, moviesList)
}
