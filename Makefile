TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
BIN_NAME?=sakura-iot-echo-server

default: test vet

run:
	go run $(CURDIR)/cmd/echo_server.go $(ARGS)

clean:
	rm -Rf $(CURDIR)/bin/*

build: clean vet
	govendor build -ldflags "-s -w -X `go list ./version`.Revision=`git rev-parse --short HEAD 2>/dev/null`" -o $(CURDIR)/bin/$(BIN_NAME) $(CURDIR)/cmd/echo_server.go

build-x: clean vet
	sh -c "'$(CURDIR)/scripts/build.sh' '$(BIN_NAME)'"

test: vet
	govendor test $(TEST) $(TESTARGS) -v -timeout=30m -parallel=4 ;

vet: golint
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

golint: fmt
	golint $(go list ./... | grep -v vendor)
fmt:
	gofmt -s -l -w $(GOFMT_FILES)

docker-run: 
	sh -c "'$(CURDIR)/scripts/build_docker_image.sh' '$(BIN_NAME)'" ; \
	sh -c "'$(CURDIR)/scripts/run_on_docker.sh' '$(BIN_NAME)'"

docker-daemon:
	sh -c "'$(CURDIR)/scripts/build_docker_image.sh' '$(BIN_NAME)'" ; \
	sh -c "'$(CURDIR)/scripts/run_on_docker_daemon.sh' '$(BIN_NAME)'"

docker-logs:
	docker logs -f $(BIN_NAME)

docker-rm:
	docker rm -f $(BIN_NAME)

docker-test: 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

docker-build: clean 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"


.PHONY: default test vet fmt lint
