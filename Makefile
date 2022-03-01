SHELL=/bin/bash

.PHONY: run_back
run_back:
	go run -race cmd/uri-one/main.go run -config=./configs/config.dev.yaml

.PHONY: build_back build_font
build_back:
	bash scripts/build.sh amd64

.PHONY: linter
linter:
	bash scripts/linter.sh

.PHONY: tests
tests:
	bash scripts/tests.sh

.PHONY: develop_up develop_down
develop_up:
	bash scripts/docker.sh docker_up
develop_down:
	bash scripts/docker.sh docker_down

.PHONY: ci
ci:
	bash scripts/ci.sh

.PHONY: deb
deb: 
	deb-builder build