package server

import (
	"context"

	"github.com/google/uuid"
	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
)

func (s *MoviesServer) DeleteMovie(ctx context.Context, req *moviesv1.DeleteMovieRequest) (*moviesv1.DeleteMovieResponse, error) {
	err := s.deleteMoviesService.Call(uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}

	res := moviesv1.DeleteMovieResponse{}

	return &res, nil
}
