package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var incomingRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "incoming_requests", // metric name
		Help: "Number of incoming update requests",
	},
	[]string{"operation"},
)

var incomingRequestsSuccess = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "incoming_requests_success", // metric name
		Help: "Number of successfully incoming update requests",
	},
	[]string{"operation"},
)

func registerMetrics() error {
	if err := prometheus.Register(incomingRequests); err != nil {
		return err
	}
	if err := prometheus.Register(incomingRequestsSuccess); err != nil {
		return err
	}
	return nil
}

func IncIncomingRequests(operation string) {
	incomingRequests.With(prometheus.Labels{"operation": operation}).Inc()
}

func IncIncomingRequestsSuccess(operation string) {
	incomingRequests.With(prometheus.Labels{"operation": operation}).Inc()
}

func RunMetricsServer(ctx context.Context) error {
	ch := make(chan struct{})
	defer close(ch)

	if err := registerMetrics(); err != nil {
		log.Error().Err(err).Msg("Failed to register metrcis")
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s := http.Server{Addr: "0.0.0.0:9100", Handler: mux}

	shutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("Failed to shutdown")
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

	log.Info().Msg("Metrcis server listening on 0.0.0.0:9100")
	if err := s.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("Failed to start metrics server")
		return err
	}
	return nil
}
