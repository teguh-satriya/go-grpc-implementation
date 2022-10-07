package services

import (
	"github.com/teguh-satriya/go-grpc-implementation/models"
	repositories "github.com/teguh-satriya/go-grpc-implementation/repository"
)

type ListMoviesService interface {
	Call() (res *ListMovieResult, err error)
}

type ListMoviesServiceImpl struct {
	repo repositories.MoviesRepository
}

type ListMovieParams struct{}

type ListMovieResult struct {
	Movies map[string]models.Movie
}

func (s *ListMoviesServiceImpl) Call() (res *ListMovieResult, err error) {
	data, err := s.repo.List()

	if err != nil {
		return nil, err
	}

	res = &ListMovieResult{
		Movies: data,
	}

	return res, nil
}

func NewListMoviesService(
	repo repositories.MoviesRepository,
) ListMoviesService {
	return &ListMoviesServiceImpl{
		repo: repo,
	}
}
