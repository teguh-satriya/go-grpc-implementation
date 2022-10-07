package services

import (
	"github.com/google/uuid"
	"github.com/teguh-satriya/go-grpc-implementation/models"
	repositories "github.com/teguh-satriya/go-grpc-implementation/repository"
)

type CreateMoviesService interface {
	Call(params *CreateMovieParams) (*CreateMovieResult, error)
}

type CreateMoviesServiceImpl struct {
	repo repositories.MoviesRepository
}

type CreateMovieParams struct {
	Title   string
	Summary string
	Rating  uint
}

type CreateMovieResult struct {
	ID      string
	Title   string
	Summary string
	Rating  uint
}

func (s *CreateMoviesServiceImpl) Call(movie *CreateMovieParams) (res *CreateMovieResult, err error) {
	data, err := s.repo.Set(models.Movie{
		ID:      uuid.New().String(),
		Title:   movie.Title,
		Summary: movie.Summary,
		Rating:  movie.Rating,
	})

	if err != nil {
		return nil, err
	}

	res = &CreateMovieResult{
		ID:      data.ID,
		Title:   data.Title,
		Summary: data.Summary,
		Rating:  data.Rating,
	}

	return res, nil
}

func NewCreateMoviesService(
	repo repositories.MoviesRepository,
) CreateMoviesService {
	return &CreateMoviesServiceImpl{
		repo: repo,
	}
}
