.PHONY: build
build:
	go build -v ./cmd/sso

run:
	CONFIG_PATH="./config/local.yaml" ./sso

.DEFAULT_GOAL := build