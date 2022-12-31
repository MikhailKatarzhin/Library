pwd = $(PWD)

help: ## Show this help.
## --------------------------------------------------------------------------
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

ci: ##               Run CI local, needs docker.
	make lint_docker
	make test_with_docker

lint: ##             Run linters for project using golangci-lint on this machine.
	golangci-lint run -v

lint_docker: ##      Run linter in the docker.
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v

test: ##             Run tests.
	go test -race -bench=. -benchmem ./...

test_with_docker: ## Run tests for each needs docker.
## --------------------------------------------------------------------------
	go test -tags testdocker -race -bench=. -benchmem ./...

ginkgo_run: ##   Run ginkgo tests.
	ACK_GINKGO_DEPRECATIONS=2.6.1 ginkgo run \
		--tags testdocker \
		--succinct \
		--output-interceptor-mode=none -
		-trace -r -p

ginkgo_watch: ## Run ginkgo watch for tests.
	ACK_GINKGO_DEPRECATIONS=2.6.1 ginkgo watch \
		--tags testdocker \
		--vv \
		--output-interceptor-mode=none \
		--trace -r -p
