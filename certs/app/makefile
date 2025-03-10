# Makefile

# Variables
IMAGE_NAME := http-client:v1.0
DOCKERFILE_PATH := ./Dockerfile

# Default target
all: build

# Build Docker image
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH) .

# Clean Docker image
clean:
	docker rmi $(IMAGE_NAME)

.PHONY: all build clean