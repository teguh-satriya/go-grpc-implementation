package server

import (
	"context"
	"fmt"

	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
)

func (s *MoviesServer) ListMovies(ctx context.Context, req *moviesv1.ListMoviesRequest) (*moviesv1.ListMoviesResponse, error) {
	result, err := s.listMoviesService.Call()
	if err != nil {
		return nil, err
	}

	fmt.Printf("result: %v\n", result)

	res := &moviesv1.ListMoviesResponse{
		Data: []*moviesv1.Movie{},
	}

	for _, item := range result.Movies {
		row := &moviesv1.Movie{
			Id:      item.ID,
			Title:   item.Title,
			Summary: item.Summary,
			Rating:  uint32(item.Rating),
		}

		res.Data = append(res.Data, row)
	}

	return res, nil
}
