package server

import (
	"context"

	"github.com/google/uuid"
	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
)

func (s *MoviesServer) GetMovie(ctx context.Context, req *moviesv1.GetMovieRequest) (*moviesv1.GetMovieResponse, error) {
	result, err := s.getMoviesService.Call(uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}

	res := moviesv1.GetMovieResponse{
		Data: &moviesv1.Movie{
			Id:      result.ID,
			Title:   result.Title,
			Summary: result.Summary,
			Rating:  uint32(result.Rating),
		},
	}

	return &res, nil
}
