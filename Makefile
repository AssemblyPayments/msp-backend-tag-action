
GIT_TAG := $(shell git describe)
DOCKER_IMAGE := mx51io/version-json-tagging-action

.PHONY: build
build:
	docker build --rm \
		-t ${DOCKER_IMAGE}:latest \
		-t ${DOCKER_IMAGE}:${GIT_TAG} \
		.

.PHONY: push
push:
	docker push ${DOCKER_IMAGE}:${GIT_TAG}
