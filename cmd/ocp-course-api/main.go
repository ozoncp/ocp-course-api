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
	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/event_producer"
	"github.com/ozoncp/ocp-course-api/internal/repo/impl"
	"github.com/ozoncp/ocp-course-api/internal/utils"
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
	uc "github.com/ozoncp/ocp-course-api/internal/utils/config"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

func eventsReader(
	ctx context.Context,
	events <-chan model.CourseEvent,
	p event_producer.CourseEventProducer,
) error {
	for {
		select {
		case <-ctx.Done():
			log.Debug().Msg("Events reader is finished")
			return nil
		case evt, ok := <-events:
			if !ok {
				return nil
			}
			err := p.SendEvent(&evt)
			if err != nil {
				log.Error().Err(err).Msg("Can't send course event")
			}
		}
	}
}

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Println("Project: ocp-course-api")
	fmt.Println("Author: Aleksei Shashev")
	fmt.Println("Site: https://github.com/ozoncp/ocp-course-api")

	var defConfig *configuration.Config
	if cfg, err := uc.LoadDefault(); err != nil {
		log.Error().Err(err).Msg("Couldn't read config")
		return 1
	} else {
		defConfig = cfg
		log.Info().Msgf("read config:\n%v\n", defConfig)
	}

	var serverConfig *api.Config
	if scfg, err := api.FromHoconConfig(defConfig, "server"); err != nil {
		log.Error().Err(err).Msg("Couldn't extract server config")
		return 1
	} else {
		serverConfig = scfg
	}

	dsn := defConfig.GetString("database.data-source-name")
	if len(dsn) == 0 {
		log.Error().Msg("Data Source Name shouldn't be empty")
		return 1
	}

	batchSize, err := commons.NewNaturalInt(int(defConfig.GetInt32("service.batch-size")))
	if err != nil {
		log.Error().Err(err).Msg("Wrong batch size")
		return 1
	}

	ctxInterrupted, stop1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop1()

	ctxTerminated, stop2 := signal.NotifyContext(ctxInterrupted, syscall.SIGTERM)
	defer stop2()
	g, ctx := errgroup.WithContext(ctxTerminated)

	db, err := utils.ConnectToPostgres(ctx, dsn)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't connect to DB")
		return 1
	}

	repo := impl.NewRepoCourse(ctx, db)

	eventProducer, err := event_producer.MakeCourseEventProducerWithRetry(
		ctx,
		defConfig.GetString("kafka.topic"),
		defConfig.GetStringList("kafka.brokers"),
	)

	if err != nil {
		log.Error().Err(err).Msg("Couldn't create producer")
		return 1
	}

	events := make(chan model.CourseEvent, defConfig.GetInt32("kafka.buffer", 16))
	defer close(events)

	g.Go(func() error { return eventsReader(ctx, events, eventProducer) })

	g.Go(func() error {
		return api.RunGrpc(ctx, serverConfig, func(s grpc.ServiceRegistrar) {
			pb.RegisterOcpCourseApiServer(s, api.NewOcpCourseApi(repo, events, batchSize))
		})
	})

	g.Go(func() error {
		return api.RunHttp(ctx, serverConfig, pb.RegisterOcpCourseApiHandlerFromEndpoint)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error occurs: %v\n", err)
	}

	return 0
}
