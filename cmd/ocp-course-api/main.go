package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-akka/configuration"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-course-api/internal/api"
	"github.com/ozoncp/ocp-course-api/internal/repo/impl"
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

	dsn := defConfig.GetString("database.data-source-name")
	if len(dsn) == 0 {
		log.Fatal().Msg("Data Source Name shouldn't be empty")
	}

	ctxInterrupted, stop1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop1()

	ctxTerminated, stop2 := signal.NotifyContext(ctxInterrupted, syscall.SIGTERM)
	defer stop2()
	g, ctx := errgroup.WithContext(ctxTerminated)

	var db *sqlx.DB

	err := backoff.Retry(func() error {
		var err error
		db, err = sqlx.Open("pgx", dsn)
		if err != nil {
			log.Debug().Err(err).Msg("Attempt to open connection to DB failed")
			return err
		}
		err = db.Ping()
		if err != nil {
			log.Debug().Err(err).Msg("Attempt to connect to DB failed")
		}
		return err
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))

	if err != nil {
		log.Error().Err(err).Msg("Couldn't connect to DB")
		return
	}

	repo := impl.NewRepoCourse(ctx, db)

	g.Go(func() error {
		return api.RunGrpc(ctx, serverConfig, func(s grpc.ServiceRegistrar) {
			pb.RegisterOcpCourseApiServer(s, api.NewOcpCourseApi(repo))
		})
	})

	g.Go(func() error {
		return api.RunHttp(ctx, serverConfig, pb.RegisterOcpCourseApiHandlerFromEndpoint)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error occurs: %v\n", err)
	}
}
