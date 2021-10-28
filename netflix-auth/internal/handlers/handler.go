package handlers

import (
	"net/http"

	"github.com/justinas/alice"
	"google.golang.org/grpc"
	"netflix-auth/internal/config"
	"netflix-auth/internal/handlers/auth"
	"netflix-auth/internal/handlers/movies"
	"netflix-auth/internal/handlers/users"
	"netflix-auth/internal/middlewares"
	"netflix-auth/internal/repository"
	authService "netflix-auth/internal/services/auth"
	moviesService "netflix-auth/internal/services/movies"
	userService "netflix-auth/internal/services/users"
	"netflix-auth/pkg/hash"
	"netflix-auth/pkg/jwt"
)

const version1 = "/v1"

type Handler struct {
	js jwt.Service
	us userService.Service

	user  users.Handler
	auth  auth.Handler
	movie movies.Handler
}

func New(repos *repository.Repository, conn *grpc.ClientConn, cfg *config.Config) *Handler {
	js := jwt.NewJWTService(cfg, repos.MemStore)
	hs := hash.NewHashService(cfg)
	us := userService.NewUserService(repos.User, hs)

	as := authService.NewAuthService(us, js, hs)
	ms := moviesService.NewMovieService(conn, cfg)

	return &Handler{
		js: js,
		us: us,

		user:  users.NewHandler(us),
		auth:  auth.NewHandler(as),
		movie: movies.NewHandler(ms),
	}
}

func (h *Handler) InitRoutes() *router {
	var (
		r              = newRouter()
		v1             = r.subRouter(version1)
		authMiddleware = middlewares.NewAuth(h.js, h.us)

		pubChain                 = alice.New(middlewares.PanicRecoveryMiddleware)
		authChain                = pubChain.Append(authMiddleware.WithUserID)
		authChainWithTokenClaims = pubChain.Append(authMiddleware.WithTokenClaims)

		usersRouter              = v1.subRouter("/users")
		moviesRouter             = v1.subRouter("/movies")
	)

	usersRouter.chain(pubChain).handle("/create", h.user.Create, http.MethodPost)
	usersRouter.chain(pubChain).handle("/auth", h.auth.Auth, http.MethodPost)
	usersRouter.chain(authChainWithTokenClaims).handle("/logout", h.auth.LogOut, http.MethodPost)

	usersRouter.chain(authChain).handle("/watched", h.movie.GetUserWatchedList, http.MethodGet)
	usersRouter.chain(authChain).handle("/bookmarks", h.movie.GetUserBookmarks, http.MethodGet)

	moviesRouter.chain(pubChain).handle("/search/{name}", h.movie.Search, http.MethodGet)
	moviesRouter.chain(authChain).handle("/{id}/add-bookmark", h.movie.AddBookmark, http.MethodPost)
	moviesRouter.chain(authChain).handle("/{id}/add-to-watch-list", h.movie.AddToWatchedList, http.MethodPost)

	return r
}
