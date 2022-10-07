package server

import (
	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
	"github.com/teguh-satriya/go-grpc-implementation/services"
)

type MoviesServer struct {
	moviesv1.UnimplementedMoviesServiceServer
	listMoviesService   services.ListMoviesService
	createMoviesService services.CreateMoviesService
	updateMoviesService services.UpdateMoviesService
	getMoviesService    services.GetMoviesService
	deleteMoviesService services.DeleteMoviesService
}

type MoviesServerSetter func(server *MoviesServer)

func NewMoviesServer(setters ...MoviesServerSetter) *MoviesServer {
	s := &MoviesServer{}

	for _, set := range setters {
		set(s)
	}

	return s
}

func WithListMoviesService(listMoviesService services.ListMoviesService) MoviesServerSetter {
	return func(as *MoviesServer) {
		as.listMoviesService = listMoviesService
	}
}

func WithUpdateMoviesService(updateMoviesService services.UpdateMoviesService) MoviesServerSetter {
	return func(as *MoviesServer) {
		as.updateMoviesService = updateMoviesService
	}
}

func WithGetMoviesService(getMoviesService services.GetMoviesService) MoviesServerSetter {
	return func(as *MoviesServer) {
		as.getMoviesService = getMoviesService
	}
}

func WithCreateMoviesService(createMoviesService services.CreateMoviesService) MoviesServerSetter {
	return func(as *MoviesServer) {
		as.createMoviesService = createMoviesService
	}
}

func WithDeleteMoviesService(deleteMoviesService services.DeleteMoviesService) MoviesServerSetter {
	return func(as *MoviesServer) {
		as.deleteMoviesService = deleteMoviesService
	}
}
