package handlers

import (
	"net/http"
	"netflix-auth/internal/services"
	"netflix-auth/pkg/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *services.Service
}

func New(s *services.Service) *Handler {
	return &Handler{services: s}
}

func (h Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/v1").Subrouter()

	users := router.PathPrefix("/users").Subrouter()
	usersAuth := router.PathPrefix("/users").Subrouter()
	usersAuth.Use(h.services.Auth.Middleware)

	{
		users.HandleFunc("/create", h.CreateUser).Methods(http.MethodPost)
		users.HandleFunc("/auth", h.Auth).Methods(http.MethodPost)
		usersAuth.HandleFunc("/watched", h.GetUserWatchedList).Methods(http.MethodGet)
		usersAuth.HandleFunc("/bookmarks", h.GetUserBookmarks).Methods(http.MethodGet)
	}

	movies := router.PathPrefix("/movies").Subrouter()
	moviesAuth := router.PathPrefix("/movies").Subrouter()
	moviesAuth.Use(h.services.Auth.Middleware)

	{
		movies.Path("/search/{name}").HandlerFunc(h.Search).Methods(http.MethodGet)
		moviesAuth.HandleFunc("/{id}/add-bookmark", h.AddBookmark).Methods(http.MethodPost)
		moviesAuth.HandleFunc("/{id}/add-to-watch-list", h.AddToWatchedList).Methods(http.MethodPost)
	}

	router.Use(utils.PanicRecoveryMiddleware)

	return router
}
