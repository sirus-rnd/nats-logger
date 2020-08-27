GO111MODULE=on
BINARY_NAME=nats-logger
BINARY_NAME_CROSS_LINUX=nats-logger
RIMRAF=rm -rf
PACKAGE_NAME=go.sirus.dev/sirus-rnd/nats-logger

# setup OS variables
ifeq ($(OS), Windows_NT)
	BINARY_NAME=nats-logger.exe
endif

.PHONY: all test docs

all:
	make init
	make build

init:
	make clean
	make install-dependency

build:
	go build -o $(BINARY_NAME) -v

run:
	go run $(PACKAGE_NAME)

help:
	go run $(PACKAGE_NAME) help

lint:
	revive -config revive.toml -formatter stylish $(PACKAGE_NAME) pkg/...

test:
	ginkgo -cover -outputdir=./coverage ./...

mock:
	mockgen -destination pkg/linimasa/mock/edge_central.go go.sirus.dev/sirus-rnd/nats-logger/pkg/linimasa IEdgeCentral

merge-coverage:
	@for f in `find . -name \*.coverprofile`; do tail -n +2 $$f >>_total; done
	@echo 'mode: atomic' >total.coverprofile
	@awk -f merge-profiles.awk <_total >>total.coverprofile

build-cross-linux:
	make init
	make lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME_CROSS_LINUX) -v

install-dependency:
	go get -v github.com/mgechev/revive
	go get -v github.com/onsi/ginkgo/ginkgo
	go get -v github.com/golang/mock/mockgen@v1.4.3
	go mod tidy

clean:
	$(RIMRAF) $(BINARY_NAME)
