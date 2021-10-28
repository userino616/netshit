package movies

import (
	"context"
	"netflix-auth/internal/config"
	"time"

	"github.com/google/uuid"
	"github.com/userino616/netflix-grpc/movieservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	httperror "netflix-auth/pkg/http_error"
)

type Service struct {
	grpcClient movieservice.MovieServiceClient
	timeOut    time.Duration
}

func NewMovieService(conn *grpc.ClientConn, cfg *config.Config) Service {
	return Service{
		grpcClient: movieservice.NewMovieServiceClient(conn),
		timeOut:    time.Duration(cfg.Server.GRPCTimeout) * time.Second,
	}
}

func (s Service) Search(
	req *movieservice.SearchMovieRequest,
) (*movieservice.MovieListResponse, httperror.HTTPError) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), s.timeOut)
	defer cancelFunc()
	resp, err := s.grpcClient.Search(ctx, req)
	if err != nil {
		return nil, httperror.NewInternalServerErr(err)
	}

	res := &movieservice.MovieListResponse{Movies: resp.GetMovies()}

	return res, nil
}

func (s Service) AddBookmark(
	bookmark *movieservice.AddBookmarkRequest,
) httperror.HTTPError {
	ctx, cancelFunc := context.WithTimeout(context.Background(), s.timeOut)
	defer cancelFunc()
	_, err := s.grpcClient.AddBookmark(ctx, bookmark)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.AlreadyExists {
				return httperror.NewBadRequestErr(err, st.Message())
			}
			if st.Code() == codes.NotFound {
				return httperror.NewNotFoundErr(err, st.Message())
			}
		}

		return httperror.NewInternalServerErr(err)
	}

	return nil
}

func (s Service) AddToWatchedList(
	watchedMovie *movieservice.AddToWatchedListRequest,
) httperror.HTTPError {
	ctx, cancelFunc := context.WithTimeout(context.Background(), s.timeOut)
	defer cancelFunc()
	_, err := s.grpcClient.AddToWatchedList(ctx, watchedMovie)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.AlreadyExists {
				return httperror.NewBadRequestErr(err, st.Message())
			}
			if st.Code() == codes.NotFound {
				return httperror.NewNotFoundErr(err, st.Message())
			}
		}

		return httperror.NewInternalServerErr(err)
	}

	return nil
}

func (s Service) GetWatchedList(
	userID uuid.UUID,
) (*movieservice.MovieListResponse, httperror.HTTPError) {
	req := &movieservice.UserIDRequest{Id: userID.String()}

	ctx, cancelFunc := context.WithTimeout(context.Background(), s.timeOut)
	defer cancelFunc()
	resp, err := s.grpcClient.GetWatchedList(ctx, req)
	if err != nil {
		return nil, httperror.NewInternalServerErr(err)
	}

	res := &movieservice.MovieListResponse{Movies: resp.GetMovies()}

	return res, nil
}

func (s Service) GetBookmarks(
	userID uuid.UUID,
) (*movieservice.MovieListResponse, httperror.HTTPError) {
	req := &movieservice.UserIDRequest{Id: userID.String()}

	ctx, cancelFunc := context.WithTimeout(context.Background(), s.timeOut)
	defer cancelFunc()
	resp, err := s.grpcClient.GetBookmarks(ctx, req)
	if err != nil {
		return nil, httperror.NewInternalServerErr(err)
	}

	res := &movieservice.MovieListResponse{Movies: resp.GetMovies()}

	return res, nil
}
