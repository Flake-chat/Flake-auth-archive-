.PHONY: build
build:
	go build -v ./cmd/main.go

.DEFAUIT_GOAL := build