DOCKER_REGISTRY = docker.io
# GIT_SHORT_SHA := $(shell git log -1 --format="%h")

# The image name is the name of the folder, e.g: ms-cards, ms-file-upload etc.
IMAGE_NAMESPACE = gust0
IMAGE_NAME = ulil-albab-be
COMMIT_ID ?= $(shell git rev-parse --short HEAD)

# Build the Docker image and add the needed tags.
build-docker:
	docker build \
		--build-arg=SERVICE_NAME=$(IMAGE_NAME) \
		-t=$(IMAGE_NAMESPACE)/$(IMAGE_NAME) .
	docker tag $(IMAGE_NAMESPACE)/$(IMAGE_NAME) $(IMAGE_NAMESPACE)/$(IMAGE_NAME):$(COMMIT_ID)


# Pushes to the configured registry.
push:
	docker push $(IMAGE_NAMESPACE)/$(IMAGE_NAME):$(COMMIT_ID)

build:
	go build -o ulil-albab-be.app src/project/main.go


# running
run:
	go run src/project/main.go


# testing
test:
	go test -v ./...

local-test:	
	go test -v ./... -coverprofile=cover.out
	go tool cover -html=cover.out


