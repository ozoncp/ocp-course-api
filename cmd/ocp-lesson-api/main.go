package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-akka/configuration"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-course-api/internal/api"
	"github.com/ozoncp/ocp-course-api/internal/repo/impl"
	"github.com/ozoncp/ocp-course-api/internal/utils"
	uc "github.com/ozoncp/ocp-course-api/internal/utils/config"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

func main() {
	fmt.Println("Project: ocp-lesson-api")
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

	dsn := defConfig.GetString("database.data-source-name")
	if len(dsn) == 0 {
		log.Fatal().Msg("Data Source Name shouldn't be empty")
	}

	ctxInterrupted, stop1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop1()

	ctxTerminated, stop2 := signal.NotifyContext(ctxInterrupted, syscall.SIGTERM)
	defer stop2()
	g, ctx := errgroup.WithContext(ctxTerminated)

	db, err := utils.ConnectToPostgres(ctx, dsn)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't connect to DB")
		return
	}

	repo := impl.NewRepoLesson(ctx, db)

	g.Go(func() error {
		return api.RunGrpc(ctx, serverConfig, func(s grpc.ServiceRegistrar) {
			pb.RegisterOcpLessonApiServer(s, api.NewOcpLessonApi(repo))
		})
	})

	g.Go(func() error {
		return api.RunHttp(ctx, serverConfig, pb.RegisterOcpLessonApiHandlerFromEndpoint)
	})

	if err := g.Wait(); err != nil {
		log.Info().Err(err)
	}
}
