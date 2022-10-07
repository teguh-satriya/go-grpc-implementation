package repositories

import (
	"sync"

	"github.com/teguh-satriya/go-grpc-implementation/models"
)

type MoviesRepositoryImpl struct {
	data  map[string]models.Movie
	mutex sync.Mutex
}

func NewMoviesRepository() MoviesRepository {
	return &MoviesRepositoryImpl{
		data:  make(map[string]models.Movie),
		mutex: sync.Mutex{},
	}
}

func (r *MoviesRepositoryImpl) Set(movie models.Movie) (*models.Movie, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.data[movie.ID] = movie
	res := r.data[movie.ID]
	return &res, nil
}

func (r *MoviesRepositoryImpl) List() (map[string]models.Movie, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.data, nil
}

func (r *MoviesRepositoryImpl) Get(id string) (*models.Movie, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	res := r.data[id]
	return &res, nil
}

func (r *MoviesRepositoryImpl) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.data, id)
	return nil
}
