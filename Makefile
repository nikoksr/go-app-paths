fmt:
	@goimports -w .
	@golines --shorten-comments -m 120 -w .
	@gofumpt -w -l .
	@gci write -s standard -s default -s "prefix(github.com/nikoksr/go-app-paths)" .
.PHONY: fmt

lint:
	@golangci-lint run ./...
.PHONY: lint

test:
	@go test -v ./...
.PHONY: test

