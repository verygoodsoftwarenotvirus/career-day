GOPATH     := $(GOPATH)
GIT_HASH   := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

.PHONY: lambda
lambda:
	GOOS=linux go build -o main github.com/verygoodsoftwarenotvirus/career-day/lambda
	zip deployment.zip main
	rm main

.PHONY: stress
stress:
	go run stress_test/main.go

.PHONY: more_stress
more_stress:
	go run stress_test/main.go 100ms

server:
	go build github.com/verygoodsoftwarenotvirus/career-day/server -o server

run: server
	nohup sudo ./server > log.out &