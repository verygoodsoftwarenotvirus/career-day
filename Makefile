GOPATH     := $(GOPATH)
GIT_HASH   := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

.PHONY: lambda
lambda:
	GOOS=linux go build -o main github.com/verygoodsoftwarenotvirus/career-day/cmd/lambda
	zip deployment.zip main
	rm main

.PHONY: stress
stress:
	go run cmd/stress_test/main.go

.PHONY: more_stress
more_stress:
	go run cmd/stress_test/main.go 100ms

.PHONY: server
server:
	go build -o server github.com/verygoodsoftwarenotvirus/career-day/cmd/server

.PHONY: run
run: server
	nohup sudo ./server > log.out &