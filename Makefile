
GIT_TAG := $(shell git describe)

.PHONY: build
build:
	docker build --rm \
		-t mx51io/version-json-tagging-action:latest \
		-t mx51io/version-json-tagging-action:${GIT_TAG} \
		.

.PHONY: push
push:
	docker push mx51io/version-json-tagging-action:${GIT_TAG}
