package handlers

import (
	"net/http"
	"netflix-auth/pkg/response"
	"netflix-auth/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/userino616/netflix-grpc/movieservice"
)

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	movieName := mux.Vars(r)["name"]
	movies, err := h.services.Movie.Search(&movieservice.SearchMovieRequest{Name: movieName})
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	response.SendJSONResponse(w, movies)
}

func (h *Handler) AddBookmark(w http.ResponseWriter, r *http.Request) {
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
	err = h.services.Movie.AddBookmark(bookmark)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddToWatchedList(w http.ResponseWriter, r *http.Request) {
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
	err = h.services.Movie.AddToWatchedList(watchedMovie)
	if err != nil {
		response.SendErrorResponse(w, err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetUserWatchedList(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	movies, err := h.services.Movie.GetWatchedList(userID)
	if err != nil {
		response.SendErrorResponse(w, err)
	}
	response.SendJSONResponse(w, movies)
}

func (h *Handler) GetUserBookmarks(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.GetUserID(r)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}
	movies, err := h.services.Movie.GetBookmarks(userID)
	if err != nil {
		response.SendErrorResponse(w, err)

		return
	}

	response.SendJSONResponse(w, movies)
}
