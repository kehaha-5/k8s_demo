# Makefile

# Variables
IMAGE_NAME := http-client:v1.0
DOCKERFILE_PATH := ./Dockerfile

# Default target
all: build import_kind clean

# Build Docker image
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH) .

import_kind:
	kind load docker-image $(IMAGE_NAME) --name my-testing

# Clean Docker image
clean:
	docker rmi $(IMAGE_NAME)

.PHONY: all build import_kind clean