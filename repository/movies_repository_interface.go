package repositories

import "github.com/teguh-satriya/go-grpc-implementation/models"

type MoviesRepository interface {
	Set(movie models.Movie) (*models.Movie, error)
	Get(id string) (*models.Movie, error)
	List() (map[string]models.Movie, error)
	Delete(id string) error
}
