module github.com/ozoncp/ocp-course-api

go 1.16

require (
	github.com/cheekybits/genny v1.0.0
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-course-api/pkg/ocp-course-api v0.0.0-00010101000000-000000000000
	github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/goleak v1.1.10
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20210607140030-00d4fb20b1ae // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/ozoncp/ocp-course-api/pkg/ocp-course-api => ./pkg/ocp-course-api

replace github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api => ./pkg/ocp-lesson-api
