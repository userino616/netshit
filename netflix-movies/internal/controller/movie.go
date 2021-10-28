package controller

import (
	"context"
	"github.com/google/uuid"
	"github.com/userino616/netflix-grpc/movieservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"netflix-movies/internal/models"
	"netflix-movies/internal/services/movies"
	"netflix-movies/pkg/postgres"
)

type MovieController struct {
	movieservice.UnimplementedMovieServiceServer
	service movies.MovieService
}

func NewMovieController(ms movies.MovieService) *MovieController {
	return &MovieController{
		service: ms,
	}
}

func (ctrl *MovieController) Search(
	ctx context.Context,
	req *movieservice.SearchMovieRequest,
) (*movieservice.MovieListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	moviesList, err := ctrl.service.Search(req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := &movieservice.MovieListResponse{}
	for _, movie := range moviesList {
		res.Movies = append(res.Movies, marshalMovie(&movie))
	}

	return res, nil
}

func (ctrl *MovieController) GetWatchedList(
	ctx context.Context,
	req *movieservice.UserIDRequest,
) (*movieservice.MovieListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	moviesList, err := ctrl.service.GetWatchedList(uid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &movieservice.MovieListResponse{}
	for _, movie := range moviesList {
		res.Movies = append(res.Movies, marshalMovie(&movie))
	}

	return res, nil
}

func (ctrl *MovieController) GetBookmarks(ctx context.Context,
	req *movieservice.UserIDRequest,
) (*movieservice.MovieListResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	uid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	moviesList, err := ctrl.service.GetBookmarks(uid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &movieservice.MovieListResponse{}
	for _, movie := range moviesList {
		res.Movies = append(res.Movies, marshalMovie(&movie))
	}

	return res, nil
}

func (ctrl *MovieController) AddBookmark(
	ctx context.Context,
	req *movieservice.AddBookmarkRequest,
) (*movieservice.Empty, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	res := &movieservice.Empty{}

	bookmark, err := unmarshalMovieBookmark(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = ctrl.service.AddToBookmark(bookmark)
	if err != nil {
		if postgres.IsDuplicateError(err) {
			return res, status.Error(codes.AlreadyExists, "record already exists")
		}
		if postgres.IsViolatesForeignKeyError(err) {
			return res, status.Error(codes.NotFound, "movie with given id not found")
		}
		return res, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (ctrl *MovieController) AddToWatchedList(
	ctx context.Context,
	req *movieservice.AddToWatchedListRequest,
) (*movieservice.Empty, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	res := &movieservice.Empty{}

	watchedMovie, err := unmarshalWatchedMovie(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = ctrl.service.AddToWatchedList(watchedMovie)

	if err != nil {
		if postgres.IsDuplicateError(err) {
			return res, status.Error(codes.AlreadyExists, "record already exists")
		}
		if postgres.IsViolatesForeignKeyError(err) {
			return res, status.Error(codes.NotFound, "movie with given id not found")
		}

		return res, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func marshalMovie(m *models.Movie) *movieservice.Movie {
	return &movieservice.Movie{
		Id:   m.ID.String(),
		Name: m.Name,
	}
}

func unmarshalMovieBookmark(
	req *movieservice.AddBookmarkRequest,
) (*models.UserMovieBookmark, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	movieID, err := uuid.Parse(req.MovieId)
	if err != nil {
		return nil, err
	}

	bookmark := &models.UserMovieBookmark{
		UserID:  userID,
		MovieID: movieID,
	}
	return bookmark, nil
}

func unmarshalWatchedMovie(
	req *movieservice.AddToWatchedListRequest,
) (*models.UserMovieWatched, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	movieID, err := uuid.Parse(req.MovieId)
	if err != nil {
		return nil, err
	}

	bookmark := &models.UserMovieWatched{
		UserID:  userID,
		MovieID: movieID,
	}
	return bookmark, nil
}
