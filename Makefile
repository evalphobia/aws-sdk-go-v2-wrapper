.PHONY: init lint test-coverage send-coverage __setup_test

GO111MODULE=on
LINT_OPT := -E gofmt \
            -E golint \
            -E govet \
            -E gosec \
            -E unused \
            -E gosimple \
            -E structcheck \
            -E varcheck \
            -E ineffassign \
            -E deadcode \
            -E typecheck \
            -E misspell \
            -E whitespace \
            -E errcheck \
            --exclude '(comment on exported (method|function|type|const|var)|should have( a package)? comment|comment should be of the form)' \
            --timeout 5m


init:
	go mod download

lint:
	@type golangci-lint > /dev/null || go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint $(LINT_OPT) run ./...

test-coverage: __setup_test
	go test -race -covermode atomic -coverprofile=gotest.cov ./...

send-coverage:
	@type goveralls > /dev/null || go get -u github.com/mattn/goveralls
	goveralls -coverprofile=gotest.cov -service=github

__setup_test:
	@mkdir -p ./.local/s3_data
	@chmod 777 ./.local/s3_data
	docker-compose up -d