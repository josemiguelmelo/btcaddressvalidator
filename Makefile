.PHONY: go-get
go-get:
	@echo " > Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GO111MODULE=on go mod vendor

.PHONY: test
test:
	@echo " > Testing ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GO111MODULE=on go test -coverprofile=cover.out ./...

.PHONY: coverage
coverage:
	mkdir -p coverage
	gocover-cobertura < cover.out > coverage/coverage.xml

.PHONY: gofmt
gofmt:
	@echo " > Linting ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GO111MODULE=on gofmt -w .
