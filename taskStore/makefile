.PHONY: fmt test build install get swagger

fmt:
	@go fmt ./...

test:
	@go test ./...

build:
	@go build -o taskStore

install:
	@go install ./...

get:
	@go get ./...

swagger:
	@swag init -g server/server.go --parseDependency

mocks:
	@mockgen -source=$(SOURCE) -destination=$(DESTINATION) -package=$(PACKAGE)
