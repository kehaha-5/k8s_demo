# Makefile

# Variables
IMAGE_NAME := https-client:v1.0
DOCKERFILE_PATH := ./Dockerfile

# Default target
all: build clean_kind_image import_kind clean

# Build Docker image
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH) .

import_kind:
	kind load docker-image $(IMAGE_NAME) --name my-testing

# Clean Docker image
clean:
	docker rmi $(IMAGE_NAME)

clean_kind_image:
	docker exec -it my-testing-worker2 crictl rmi $(IMAGE_NAME) 
	docker exec -it my-testing-worker crictl rmi $(IMAGE_NAME) 
	docker exec -it my-testing-control-plane crictl rmi $(IMAGE_NAME) 

build_first: build import_kind clean 

.PHONY: all build import_kind clean clean_kind_image