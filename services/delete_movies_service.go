package services

import (
	"github.com/google/uuid"
	repositories "github.com/teguh-satriya/go-grpc-implementation/repository"
)

type DeleteMoviesService interface {
	Call(id uuid.UUID) error
}

type DeleteMoviesServiceImpl struct {
	repo repositories.MoviesRepository
}

func (s *DeleteMoviesServiceImpl) Call(id uuid.UUID) (err error) {
	err = s.repo.Delete(id.String())

	if err != nil {
		return err
	}

	return nil
}

func NewDeleteMoviesService(
	repo repositories.MoviesRepository,
) DeleteMoviesService {
	return &DeleteMoviesServiceImpl{
		repo: repo,
	}
}
