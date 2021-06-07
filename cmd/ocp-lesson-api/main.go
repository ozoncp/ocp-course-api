package main

import (
	"flag"
	"fmt"

	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-course-api/internal/api"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

var (
	listenInterface = flag.String("interface", "0.0.0.0", "listening interface")
	grpcPort        = flag.Int("grpc-port", 7002, "port for gRPC server endpoint")
	httpPort        = flag.Int("http-port", 7000, "port for HTTP server endpoint")
	swaggerFile     = flag.String("swagger", "swagger/ocp-lesson-api.swagger.json", "path to a file with the swagger definitions")
)

func main() {
	fmt.Println("Project: ocp-lesson-api")
	fmt.Println("Author: Aleksei Shashev")
	fmt.Println("Site: https://github.com/ozoncp/ocp-course-api")

	flag.Parse()

	config := api.NewConfig(*listenInterface, *grpcPort, *httpPort, *swaggerFile)

	go func() {
		err := api.RunGrpc(config, func(s grpc.ServiceRegistrar) {
			pb.RegisterOcpLessonApiServer(s, api.NewOcpLessonApi())
		})
		if err != nil {
			panic(err)
		}
	}()

	err := api.RunHttp(config, pb.RegisterOcpLessonApiHandlerFromEndpoint)
	if err != nil {
		panic(err)
	}
}
