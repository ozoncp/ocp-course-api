package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func serveSwagger(swaggerFile string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "api.swagger.json") {
			log.Error().Msgf("Swagger JSON not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		log.Info().Msgf("Serving %s", r.URL.Path)
		http.ServeFile(w, r, swaggerFile)
	}
}

func RunHttp(config *Config, registrator func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := registrator(ctx, gwmux, config.Grpc.Address(), opts)
	if err != nil {
		return err
	}

	mux.HandleFunc("/swagger/", serveSwagger(config.SwaggerFile))
	mux.Handle("/", gwmux)

	fmt.Printf("Server listening on %s\n", config.Http.Address())
	return http.ListenAndServe(config.Http.Address(), mux)
}

func RunGrpc(config *Config, registrator func(grpc.ServiceRegistrar)) error {
	listen, err := net.Listen("tcp", config.Grpc.Address())
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return err
	}

	s := grpc.NewServer()
	registrator(s)

	fmt.Printf("Server listening on %s\n", config.Grpc.Address())
	if err := s.Serve(listen); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return err
	}

	return nil
}
