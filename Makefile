trim_right_slash = $(patsubst %/,%,$(1))

go_path = $(shell go env GOPATH)

protoc_gen_validate = github.com/envoyproxy/protoc-gen-validate@v1.0.2

pb_includes=-I./vendor.protogen
pb_includes+=-I$(go_path)/pkg/mod/$(protoc_gen_validate)

go_files = $(filter-out %.pb.go, \
            $(filter-out %.pb.gw.go, \
            $(filter-out %.pb.validate.go, \
            $(filter-out %.gen.go, \
            $(shell find . -name *.go -print | fgrep -v 'vendor.protogen') \
            ))))

test_files = $(filter %_test.go, $(go_files))

mocks = $(shell grep -o "^[^\#%]\+go.mock.go:" Makefile | sed 's/\([^:]\+\)\:/\1/' | paste -sd ' ')
pbs = $(shell grep -o "^[^\#%]\+\(pb.go\|pb.gw.go\|pb.validate.go\):" Makefile | sed 's/\([^:]\+\)\:/\1/' | paste -sd ' ')

generated = $(mocks)
generated += $(pbs)
generated += swagger/ocp-course-api.swagger.json
generated += swagger/ocp-lesson-api.swagger.json

pkgs_with_test = $(call trim_right_slash, $(sort $(dir $(test_files))))

cmds_dir = ./cmd/
pkgs_with_cmds = $(call trim_right_slash, $(dir $(wildcard $(cmds_dir)*/main.go)))

executable = $(notdir $(pkgs_with_cmds))

all: $(executable)

$(executable): $(generated) $(filter-out %_test.go, $(go_files))
	go build $(cmds_dir)$@

go.sum: go.mod
	$(run-prepare)

pkg/ocp-course-api/course_service.pb.go: api/ocp-course-api/course_service.proto
pkg/ocp-course-api/course_service.pb.validate.go: api/ocp-course-api/course_service.proto
pkg/ocp-course-api/course_service_grpc.pb.go: api/ocp-course-api/course_service.proto
pkg/ocp-course-api/course_service.pb.gw.go: api/ocp-course-api/course_service.proto

swagger/ocp-course-api.swagger.json:api/ocp-course-api/course_service.proto

pkg/ocp-lesson-api/lesson_service.pb.go: api/ocp-lesson-api/lesson_service.proto
pkg/ocp-lesson-api/lesson_service.pb.validate.go: api/ocp-lesson-api/lesson_service.proto
pkg/ocp-lesson-api/lesson_service_grpc.pb.go: api/ocp-lesson-api/lesson_service.proto
pkg/ocp-lesson-api/lesson_service.pb.gw.go: api/ocp-lesson-api/lesson_service.proto

swagger/ocp-lesson-api.swagger.json:api/ocp-lesson-api/lesson_service.proto

%.swagger.json:
	mkdir -p $(@D)
	protoc -I$(<D) $(pb_includes) \
        --swagger_out=$(@D) \
        --swagger_opt=allow_merge=true,merge_file_name=$(basename $(@F)):swagger \
		$<

%.go.mock.go:
	@mkdir -p $(@D)
	mockgen -source $< -destination $@ -package $(notdir $(@D))

%_grpc.pb.go:
	mkdir -p $(@D)
	protoc -I$(<D) $(pb_includes) --go-grpc_out=$(@D) --go-grpc_opt=paths=source_relative $<

%.pb.validate.go:
	mkdir -p $(@D)
	protoc -I$(<D) $(pb_includes) --validate_out lang=go:$(@D) --validate_opt=paths=source_relative $<

%.pb.go:
	mkdir -p $(@D)
	protoc -I$(<D) $(pb_includes) --go_out=$(@D) --go_opt=paths=source_relative $<

%pb.gw.go:
	mkdir -p $(@D)
	protoc -I$(<D) $(pb_includes) \
		--grpc-gateway_out=$(@D) \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=source_relative \
		$<

test: $(generated)
	go test $(pkgs_with_test)

ginkgo: $(generated)
	ginkgo $(pkgs_with_test)

fmt:
	gofmt -w -l .

tidy:
	go mod tidy

clean:
	#rm -rf $(executable) $(filter-out pkg/%,$(generated))
	rm -rf $(executable)
	go clean  $(call trim_right_slash, $(sort $(dir $(go_files))))
	go clean -cache
	go clean -testcache

clean_generate:
	rm -rf $(generated)

prepare:
	$(run-prepare)

lint:
	golangci-lint run

generate: $(generated)

.PHONY: test fmt tidy clean all prepare lint generate clean_generate

define run-prepare =
go mod download
go get github.com/onsi/ginkgo/ginkgo
go get github.com/golang/mock/mockgen
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -d github.com/envoyproxy/protoc-gen-validate
go get google.golang.org/protobuf/proto@v1.31.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go install $(protoc_gen_validate)
mkdir -p vendor.protogen
@if [ ! -d vendor.protogen/google ]; then \
	git clone --depth=1 https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
	mkdir -p  vendor.protogen/google/ &&\
	mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
	rm -rf vendor.protogen/googleapis ;\
fi
endef
