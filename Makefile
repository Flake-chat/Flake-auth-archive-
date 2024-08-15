.PHONY: build
build:
	go build -v ./cmd/auth/main.go

.DEFAUIT_GOAL := build