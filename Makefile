trim_right_slash = $(patsubst %/,%,$(1))

go_files = $(filter-out %.gen.go, $(shell find . -name *.go -print))

test_files = $(filter %_test.go, $(go_files))

generics = $(filter %_generic.go, $(go_files))

generated = $(generics:.go=.go.gen.go)

mocks = $(shell grep -o "^[^\#%]\+go.mock.go:" Makefile | sed 's/\([^:]\+\)\:/\1/' | paste -sd ' ')

generated += $(mocks)

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

%generic.go.gen.go: %generic.go
	@gen_types=$(shell grep '//go:generate .*genny' $< | head -n 1 | sed 's/^.\+ gen \(".\+"\)$$/\1/');\
	echo genny -out $@ -in $< gen "$${gen_types}";\
	genny -out $@ -in $< gen "$${gen_types}"

%.go.mock.go:
	@mkdir -p $(@D)
	mockgen -source $< -destination $@ -package $(notdir $(@D))

test: $(generated)
	#go test $(pkgs_with_test)
	ginkgo $(pkgs_with_test)

fmt:
	gofmt -w -l .

tidy:
	go mod tidy

clean:
	rm -rf $(executable) $(generated)
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
endef
