.PHONY: go-get
go-get:
	@echo " > Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GO111MODULE=on go mod vendor

.PHONY: test
test:
	@echo " > Testing ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GO111MODULE=on go test ./...

.PHONY: gofmt
gofmt:
	@echo " > Linting ..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GO111MODULE=on gofmt -w .
