module github.com/ozoncp/ocp-course-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/Shopify/sarama v1.29.0
	github.com/cenkalti/backoff/v4 v4.1.1
	github.com/cheekybits/genny v1.0.0
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/containerd/continuity v0.1.0 // indirect
	github.com/go-akka/configuration v0.0.0-20200606091224-a002c0330665
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/moby/term v0.0.0-20210610120745-9d4ed1856297 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/opencontainers/runc v1.0.0-rc95 // indirect
	github.com/ory/dockertest/v3 v3.6.5
	github.com/ozoncp/ocp-course-api/pkg/ocp-course-api v0.0.0-00010101000000-000000000000
	github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v0.9.3
	github.com/rs/zerolog v1.22.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0
	go.uber.org/goleak v1.1.10
	golang.org/x/net v0.0.0-20210610132358-84b48f89b13b // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	google.golang.org/genproto v0.0.0-20210607140030-00d4fb20b1ae // indirect
	google.golang.org/grpc v1.38.0
)

replace github.com/ozoncp/ocp-course-api/pkg/ocp-course-api => ./pkg/ocp-course-api

replace github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api => ./pkg/ocp-lesson-api
