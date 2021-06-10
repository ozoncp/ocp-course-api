package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-akka/configuration"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-course-api/internal/api"
	uc "github.com/ozoncp/ocp-course-api/internal/utils/config"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

func main() {
	fmt.Println("Project: ocp-course-api")
	fmt.Println("Author: Aleksei Shashev")
	fmt.Println("Site: https://github.com/ozoncp/ocp-course-api")

	var defConfig *configuration.Config
	if cfg, err := uc.LoadDefault(); err != nil {
		log.Fatal().Err(err).Msg("Couldn't read config")
	} else {
		defConfig = cfg
		log.Info().Msgf("read config:\n%v\n", defConfig)
	}

	var serverConfig *api.Config
	if scfg, err := api.FromHoconConfig(defConfig, "server"); err != nil {
		log.Fatal().Err(err).Msg("Couldn't extract server config")
	} else {
		serverConfig = scfg
	}

	ctxInterrupted, stop1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop1()

	ctxTerminated, stop2 := signal.NotifyContext(ctxInterrupted, syscall.SIGTERM)
	defer stop2()
	g, ctx := errgroup.WithContext(ctxTerminated)

	g.Go(func() error {
		return api.RunGrpc(ctx, serverConfig, func(s grpc.ServiceRegistrar) {
			pb.RegisterOcpCourseApiServer(s, api.NewOcpCourseApi())
		})
	})

	g.Go(func() error {
		return api.RunHttp(ctx, serverConfig, pb.RegisterOcpCourseApiHandlerFromEndpoint)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error occurs: %v\n", err)
	}
}
