package server

import (
	"context"

	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
	"github.com/teguh-satriya/go-grpc-implementation/services"
)

func (s *MoviesServer) UpdateMovie(ctx context.Context, req *moviesv1.UpdateMovieRequest) (*moviesv1.UpdateMovieResponse, error) {
	request := &services.UpdateMovieParams{
		ID:      req.Id,
		Title:   req.Title,
		Summary: req.Summary,
		Rating:  uint(req.Rating),
	}

	result, err := s.updateMoviesService.Call(request)
	if err != nil {
		return nil, err
	}

	res := moviesv1.UpdateMovieResponse{
		Data: &moviesv1.Movie{
			Id:      result.ID,
			Title:   result.Title,
			Summary: result.Summary,
			Rating:  uint32(result.Rating),
		},
	}

	return &res, nil
}
