trim_right_slash = $(patsubst %/,%,$(1))

pb_includes=-I./vendor.protogen

go_files = $(filter-out %.pb.go, \
            $(filter-out %.pb.gw.go, \
            $(filter-out %.pb.validate.go, \
            $(filter-out %.gen.go, \
            $(shell find . -name *.go -print | fgrep -v 'vendor.protogen') \
            ))))

test_files = $(filter %_test.go, $(go_files))

generics = $(filter %_generic.go, $(go_files))

generated = $(generics:.go=.go.gen.go)

mocks = $(shell grep -o "^[^\#%]\+go.mock.go:" Makefile | sed 's/\([^:]\+\)\:/\1/' | paste -sd ' ')
pbs = $(shell grep -o "^[^\#%]\+\(pb.go\|pb.gw.go\|pb.validate.go\):" Makefile | sed 's/\([^:]\+\)\:/\1/' | paste -sd ' ')

generated += $(mocks)
generated += $(pbs)
generated += swagger/ocp-course-api.swagger.json
generated += swagger/ocp-lesson-api.swagger.json

pkgs_with_generics = $(call trim_right_slash, $(sort $(dir $(generics))))

pkgs_with_test = $(call trim_right_slash, $(sort $(dir $(test_files))))

cmds_dir = ./cmd/
pkgs_with_cmds = $(call trim_right_slash, $(dir $(wildcard $(cmds_dir)*/main.go)))

executable = $(notdir $(pkgs_with_cmds))

all: $(executable)

$(executable): $(generated) $(filter-out %_test.go, $(go_files)) go.sum
	go build $(cmds_dir)$@

internal/flusher/flusher_generic.go.gen.go: internal/utils/slice/sliding_generic.go.gen.go

go.sum: go.mod
	$(run-prepare)

internal/mock_repo/repo.go.mock.go: internal/repo/repo_generic.go
internal/mock_flusher/flusher.go.mock.go: internal/flusher/flusher_generic.go
internal/mock_saver/flush_alarm.go.mock.go: internal/saver/flush_alarm.go

pkg/ocp-course-api/course_messages.pb.go: api/ocp-course-api/course_messages.proto
pkg/ocp-course-api/course_messages.pb.validate.go: api/ocp-course-api/course_messages.proto
pkg/ocp-course-api/course_service_grpc.pb.go: api/ocp-course-api/course_service.proto
pkg/ocp-course-api/course_service.pb.gw.go: api/ocp-course-api/course_service.proto

swagger/ocp-course-api.swagger.json:api/ocp-course-api/course_service.proto

pkg/ocp-lesson-api/lesson_messages.pb.go: api/ocp-lesson-api/lesson_messages.proto
pkg/ocp-lesson-api/lesson_messages.pb.validate.go: api/ocp-lesson-api/lesson_messages.proto
pkg/ocp-lesson-api/lesson_service_grpc.pb.go: api/ocp-lesson-api/lesson_service.proto
pkg/ocp-lesson-api/lesson_service.pb.gw.go: api/ocp-lesson-api/lesson_service.proto

swagger/ocp-lesson-api.swagger.json:api/ocp-lesson-api/lesson_service.proto

%generic.go.gen.go: %generic.go
	@gen_types=$(shell grep '//go:generate .*genny' $< | head -n 1 | sed 's/^.\+ gen \(".\+"\)$$/\1/');\
	echo genny -out $@ -in $< gen "$${gen_types}";\
	genny -out $@ -in $< gen "$${gen_types}"

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
	#go test $(pkgs_with_test)
	ginkgo $(pkgs_with_test)

fmt:
	gofmt -w -l .

tidy:
	go mod tidy

clean:
	rm -rf $(executable) $(filter-out pkg/%,$(generated))
	go clean -cache -testcache $(call trim_right_slash, $(sort $(dir $(go_files))))

prepare:
	$(run-prepare)

lint:
	golangci-lint run

generate: $(generated)

.PHONY: test fmt tidy clean all prepare lint generate

define run-prepare =
go mod download
go get github.com/cheekybits/genny
go get github.com/onsi/ginkgo/ginkgo
go get github.com/golang/mock/mockgen
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -d github.com/envoyproxy/protoc-gen-validate
go get github.com/golang/protobuf/descriptor@v1.5.2
go get github.com/golang/protobuf/proto@v1.5.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
mkdir -p vendor.protogen
@if [ ! -d vendor.protogen/google ]; then \
	git clone --depth=1 https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
	mkdir -p  vendor.protogen/google/ &&\
	mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
	rm -rf vendor.protogen/googleapis ;\
fi
@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
	mkdir -p vendor.protogen/github.com/envoyproxy &&\
	git clone --depth=1 https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
fi
endef
