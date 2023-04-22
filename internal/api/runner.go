package api

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozoncp/ocp-course-api/internal/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serveSwagger(swaggerFile string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "api.swagger.json") {
			log.Error().Msgf("Swagger JSON not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		log.Info().Msgf("Serving %s, return %v", r.URL.Path, swaggerFile)
		http.ServeFile(w, r, swaggerFile)
	}
}

func RunHttp(ctx context.Context, config *Config, registrator func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error) error {
	ch := make(chan struct{})
	defer close(ch)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := registrator(ctx, gwmux, config.Grpc.Address(), opts)
	if err != nil {
		return err
	}

	mux.HandleFunc("/swagger/", serveSwagger(config.SwaggerFile))
	mux.Handle("/", gwmux)

	s := http.Server{
		Addr:    config.Http.Address(),
		Handler: utils.TracingWrapper(mux),
	}

	shutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("failed to shutdown")
		}
	}

	go func() {
		for {
			select {
			case <-ch:
				shutdown()
				return
			case <-ctx.Done():
				shutdown()
				return
			}
		}
	}()

	log.Info().Msgf("Server listening on %s", config.Http.Address())
	if err := s.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return err
	}

	return nil
}

func RunGrpc(ctx context.Context, config *Config, registrator func(grpc.ServiceRegistrar)) error {
	ch := make(chan struct{})
	defer close(ch)
	listen, err := net.Listen("tcp", config.Grpc.Address())

	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return err
	}

	s := grpc.NewServer()

	go func() {
		for {
			select {
			case <-ch:
				s.GracefulStop()
				return
			case <-ctx.Done():
				s.GracefulStop()
				return
			}
		}
	}()

	registrator(s)

	log.Info().Msgf("Server listening on %s", config.Grpc.Address())
	if err := s.Serve(listen); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return err
	}

	return nil
}
