package services

import (
	"netflix-auth/internal/config"
	"netflix-auth/internal/repository"
	authentication "netflix-auth/internal/services/auth"
	"netflix-auth/internal/services/movies"
	"netflix-auth/internal/services/users"
	"netflix-auth/pkg/jwt"
	"netflix-auth/pkg/passwords"

	"google.golang.org/grpc"
)

type Service struct {
	Auth  authentication.AuthService
	User  users.UserService
	Movie movies.MovieService
}

func New(
	repos *repository.Repository,
	grpcConn grpc.ClientConnInterface,
	cfg *config.Config,
) *Service {
	jwtService := jwt.NewJWTService(cfg)
	passwordService := passwords.NewPasswordService(cfg)

	return &Service{
		Auth:  authentication.NewAuthService(repos.User, jwtService, passwordService),
		User:  users.NewUserService(repos.User, passwordService),
		Movie: movies.NewMovieClient(grpcConn, cfg),
	}
}
