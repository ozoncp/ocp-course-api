trim_right_slash = $(patsubst %/,%,$(1))

go_files = $(foreach \
		f,\
		$(shell find . -name *.go -print),\
		$(if $(findstring /gen-, $f),,$f))

test_files = $(filter %_test.go, $(go_files))

generics = $(filter %_generic.go, $(go_files))

generated = $(foreach \
		f,\
		$(generics),\
		$(dir $f)gen-$(notdir $f))

pkgs_with_generics = $(call trim_right_slash, $(sort $(dir $(generics))))

pkgs_with_test = $(call trim_right_slash, $(sort $(dir $(test_files))))

cmds_dir = ./cmd/
pkgs_with_cmds = $(call trim_right_slash, $(dir $(wildcard $(cmds_dir)*/main.go)))

executable = $(notdir $(pkgs_with_cmds))

all: $(executable)

$(executable): $(generated) $(filter-out %_test.go, $(go_files)) go.sum
	go build $(cmds_dir)$@

$(generated) &: $(generics) go.sum
	go generate $(pkgs_with_generics)

go.sum: go.mod
	$(run-prepare)

test: $(generated)
	go test $(pkgs_with_test)

fmt:
	gofmt -w -l .

tidy:
	go mod tidy

clean:
	rm -rf $(executable) $(generated)

prepare:
	$(run-prepare)

.PHONY: test fmt tidy clean all prepare

define run-prepare =
go mod tidy
go mod download
go get github.com/cheekybits/genny
go mod tidy
endef
