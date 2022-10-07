package services

import (
	"github.com/google/uuid"
	repositories "github.com/teguh-satriya/go-grpc-implementation/repository"
)

type GetMoviesService interface {
	Call(id uuid.UUID) (*GetMovieResult, error)
}

type GetMoviesServiceImpl struct {
	repo repositories.MoviesRepository
}

type GetMovieParams struct {
	ID string
}

type GetMovieResult struct {
	ID      string
	Title   string
	Summary string
	Rating  uint
}

func (s *GetMoviesServiceImpl) Call(id uuid.UUID) (res *GetMovieResult, err error) {
	data, err := s.repo.Get(id.String())

	if err != nil {
		return nil, err
	}

	res = &GetMovieResult{
		ID:      data.ID,
		Title:   data.Title,
		Summary: data.Summary,
		Rating:  data.Rating,
	}

	return res, nil
}

func NewGetMoviesService(
	repo repositories.MoviesRepository,
) GetMoviesService {
	return &GetMoviesServiceImpl{
		repo: repo,
	}
}
