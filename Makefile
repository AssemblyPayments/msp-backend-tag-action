
.PHONY: build
build:
	# TODO: replace sandyleo26 with mx51 account name
	docker build --rm \
		-f dockerfile.version-json-tagging \
		-t sandyleo26/version-json-tagging:latest \
		.

.PHONY: push
push:
	# TODO: replace sandyleo26 with mx51 account name
	docker push sandyleo26/version-json-tagging:latest
