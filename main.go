package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	moviesv1 "github.com/teguh-satriya/go-grpc-implementation/proto/movies/v1"
	repositories "github.com/teguh-satriya/go-grpc-implementation/repository"
	"github.com/teguh-satriya/go-grpc-implementation/server"
	"github.com/teguh-satriya/go-grpc-implementation/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	logger             = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":8080", "gRPC server endpoint") // NOTE: grpc server endpoint options
)

func main() {
	flag.Parse()

	repo := repositories.NewMoviesRepository()

	listMoviesService := services.NewListMoviesService(repo)
	getMoviesService := services.NewGetMoviesService(repo)
	createMoviesService := services.NewCreateMoviesService(repo)
	updateMoviesService := services.NewUpdateMoviesService(repo)
	deleteMoviesService := services.NewDeleteMoviesService(repo)

	moviesServiceServer := server.NewMoviesServer(
		server.WithListMoviesService(listMoviesService),
		server.WithGetMoviesService(getMoviesService),
		server.WithCreateMoviesService(createMoviesService),
		server.WithUpdateMoviesService(updateMoviesService),
		server.WithDeleteMoviesService(deleteMoviesService),
	)

	// NOTE: Initialize gRPC Dial Option
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// NOTE: Initialize TCP Connection
	tcp, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		logger.Fatalf("net.Listen: cannot initialize tcp connection")
	}

	// NOTE: Create gRPC Server
	srv := grpc.NewServer()

	// NOTE: Create Mux Handler
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				AllowPartial:    true,
				EmitUnpopulated: false,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	grpclog.SetLoggerV2(logger)

	// NOTE: Setup context, so the requets can be cancelled
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// NOTE: Run grpc server as go routine
	go func() {
		// NOTE: Register internal servers
		moviesv1.RegisterMoviesServiceServer(srv, moviesServiceServer)

		srv.Serve(tcp)
	}()

	// NOTE: Start HTTP server (and proxy calls to gRPC server endpoint)
	// NOTE: Regsiter request servers
	err = moviesv1.RegisterMoviesServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, dialOptions)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	httpServer.ListenAndServe()
}
