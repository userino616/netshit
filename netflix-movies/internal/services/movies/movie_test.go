package movies

import (
	"netflix-movies/internal/fixtures"
	"netflix-movies/internal/models"
	"netflix-movies/internal/repository"
	"netflix-movies/pkg/postgres"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	r repository.MovieRepository
	s *movieService

	existsUserID, _ = uuid.Parse("06ea7596-210d-11ec-a866-0242ac140003")

	firstMovieID, _  = uuid.Parse("9e7f9371-51a1-43f1-924f-004cfd6e6571")
	secondMovieID, _ = uuid.Parse("d49f5dab-22d4-45b2-9ac4-e4bc089d74b5")

	movie1 = models.Movie{
		ID:   firstMovieID,
		Name: "Movie 1",
	}

	movie2 = models.Movie{
		ID:   secondMovieID,
		Name: "Movie 2",
	}
)

// Если запускать тесты по отдельность то сначала вызывается инит и база приводится в ожидаемый вид
// Если запускать тесты разом то тест AddToWatchList записывает в бд новую запись и тогда тест GetWatchedList падает
func init() {
	fixtures.PrepareFixtures()
	r = repository.NewMovieRepository(fixtures.GetDB())
	s = NewMovieService(r)
}

func TestMovieService_AddToBookmark(t *testing.T) {
	tableCases := []struct {
		testName  string
		bookmark  *models.UserMovieBookmark
		pgCodeErr string
	}{
		{
			testName: "valid test",
			bookmark: &models.UserMovieBookmark{
				UserID:  existsUserID,
				MovieID: secondMovieID,
			},
			pgCodeErr: "",
		},
		{
			testName: "movie doesn't exists",
			bookmark: &models.UserMovieBookmark{
				UserID:  existsUserID,
				MovieID: uuid.New(),
			},
			pgCodeErr: postgres.ViolatesForeignKeyCode,
		},
		{
			testName: "user doesn't exists",
			bookmark: &models.UserMovieBookmark{
				UserID:  uuid.New(),
				MovieID: firstMovieID,
			},
			pgCodeErr: postgres.ViolatesForeignKeyCode,
		},
	}
	for _, tc := range tableCases {
		err := s.AddToBookmark(tc.bookmark)
		assert.Equal(t, tc.pgCodeErr, postgres.GetPgErrCode(err))
	}
}

func TestMovieService_AddToWatchedList(t *testing.T) {
	tableCases := []struct {
		testName  string
		watched   *models.UserMovieWatched
		pgCodeErr string
	}{
		{
			testName: "valid test",
			watched: &models.UserMovieWatched{
				UserID:  existsUserID,
				MovieID: secondMovieID,
			},
			pgCodeErr: "",
		},
		{
			testName: "movie doesn't exists",
			watched: &models.UserMovieWatched{
				UserID:  existsUserID,
				MovieID: uuid.New(),
			},
			pgCodeErr: postgres.ViolatesForeignKeyCode,
		},
		{
			testName: "user doesn't exists",
			watched: &models.UserMovieWatched{
				UserID:  uuid.New(),
				MovieID: firstMovieID,
			},
			pgCodeErr: postgres.ViolatesForeignKeyCode,
		},
	}
	for _, tc := range tableCases {
		err := s.AddToWatchedList(tc.watched)
		assert.Equal(t, tc.pgCodeErr, postgres.GetPgErrCode(err))
	}
}

func TestMovieService_GetBookmarks(t *testing.T) {
	bookmarks, err := s.GetBookmarks(existsUserID)
	expectedBookmarks := models.Movies{
		movie1,
	}
	assert.Equal(t, expectedBookmarks, bookmarks)
	assert.Equal(t, err, nil)
}

func TestMovieService_GetWatchedList(t *testing.T) {
	watchedList, err := s.GetWatchedList(existsUserID)
	expectedWatchedList := models.Movies{
		movie1,
	}
	assert.Equal(t, expectedWatchedList, watchedList)
	assert.Equal(t, err, nil)
}

func TestMovieService_Search(t *testing.T) {
	tableCases := []struct {
		searchName string
		expected   models.Movies
		err        error
	}{
		{
			searchName: "1",
			expected: models.Movies{
				movie1,
			},
			err: nil,
		}, {
			searchName: "movie",
			expected: models.Movies{
				movie1,
				movie2,
			},
			err: nil,
		}, {
			searchName: "nothing",
			expected:   models.Movies(nil),
			err:        nil,
		},
	}
	for _, tc := range tableCases {
		movies, err := s.Search(tc.searchName)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.expected, movies)
	}
}
