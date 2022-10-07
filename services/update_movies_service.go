package services

import (
	"github.com/teguh-satriya/go-grpc-implementation/models"
	repositories "github.com/teguh-satriya/go-grpc-implementation/repository"
)

type UpdateMoviesService interface {
	Call(params *UpdateMovieParams) (*UpdateMovieResult, error)
}

type UpdateMoviesServiceImpl struct {
	repo repositories.MoviesRepository
}

type UpdateMovieParams struct {
	ID      string
	Title   string
	Summary string
	Rating  uint
}

type UpdateMovieResult struct {
	ID      string
	Title   string
	Summary string
	Rating  uint
}

func (s *UpdateMoviesServiceImpl) Call(movie *UpdateMovieParams) (res *UpdateMovieResult, err error) {
	data, err := s.repo.Set(models.Movie{
		ID:      movie.ID,
		Title:   movie.Title,
		Summary: movie.Summary,
		Rating:  movie.Rating,
	})

	if err != nil {
		return nil, err
	}

	res = &UpdateMovieResult{
		ID:      data.ID,
		Title:   data.Title,
		Summary: data.Summary,
		Rating:  data.Rating,
	}

	return res, nil
}

func NewUpdateMoviesService(
	repo repositories.MoviesRepository,
) UpdateMoviesService {
	return &UpdateMoviesServiceImpl{
		repo: repo,
	}
}
