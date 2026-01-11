# Proto Makefile - Protocol Buffer definitions for Wilee
#
# Targets:
#   all      - Lint and build proto files (default)
#   lint     - Validate proto files with buf
#   build    - Build proto binary image
#   push     - Push to buf.build registry (runs lint and build first)
#   generate-go - Generate Go code from proto files
#
# Variables:
#   BUF      - buf command (default: buf)
#   OUT      - Output directory for generate-go (default: .)
#   V        - Verbose mode (V=1 shows commands)

.PHONY: all lint build push generate-go

BUF ?= buf
OUT ?= .

V = 0
Q = $(if $(filter 1,$V),,@)

all: lint build

lint:
	$Q $(BUF) lint

build:
	$Q $(BUF) build

push: lint build
	$Q $(BUF) push

generate-go: buf.gen-go.yaml
	$Q mkdir -p $(OUT)
	$Q find $(OUT) -name '*.pb.go' -delete
	$Q $(BUF) generate --template buf.gen-go.yaml -o $(OUT)
