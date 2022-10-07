package server

import (
	"context"

	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
	"github.com/teguh-satriya/go-grpc-implementation/services"
)

func (s *MoviesServer) CreateMovie(ctx context.Context, req *moviesv1.CreateMovieRequest) (*moviesv1.CreateMovieResponse, error) {
	request := &services.CreateMovieParams{
		Title:   req.Title,
		Summary: req.Summary,
		Rating:  uint(req.Rating),
	}

	result, err := s.createMoviesService.Call(request)
	if err != nil {
		return nil, err
	}

	res := moviesv1.CreateMovieResponse{
		Data: &moviesv1.Movie{
			Id:      result.ID,
			Title:   result.Title,
			Summary: result.Summary,
			Rating:  uint32(result.Rating),
		},
	}

	return &res, nil
}
