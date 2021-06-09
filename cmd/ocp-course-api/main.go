package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-course-api/internal/api"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

var (
	listenInterface = flag.String("interface", "0.0.0.0", "listening interface")
	grpcPort        = flag.Int("grpc-port", 7002, "port for gRPC server endpoint")
	httpPort        = flag.Int("http-port", 7000, "port for HTTP server endpoint")
	swaggerFile     = flag.String("swagger", "swagger/ocp-course-api.swagger.json", "path to a file with the swagger definitions")
)

func main() {
	fmt.Println("Project: ocp-course-api")
	fmt.Println("Author: Aleksei Shashev")
	fmt.Println("Site: https://github.com/ozoncp/ocp-course-api")

	ctxInterrupted, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	g, ctx := errgroup.WithContext(ctxInterrupted)

	flag.Parse()

	config := api.NewConfig(*listenInterface, *grpcPort, *httpPort, *swaggerFile)

	g.Go(func() error {
		return api.RunGrpc(ctx, config, func(s grpc.ServiceRegistrar) {
			pb.RegisterOcpCourseApiServer(s, api.NewOcpCourseApi())
		})
	})

	g.Go(func() error {
		return api.RunHttp(ctx, config, pb.RegisterOcpCourseApiHandlerFromEndpoint)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error occurs: %v\n", err)
	}
}
