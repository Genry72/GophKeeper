port = 8080
buildFlag = -ldflags="-X 'main.buildVersion=`git describe --tags --abbrev=0`' -X 'main.buildDate=`date`'"

.PHONY: runServer
runServer: build
	./cmd/server/server

.PHONY: runClient
runClient: build
	./cmd/client/client

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: build
build:
	go build $(buildFlag) -o ./cmd/server/ ./cmd/server/
	chmod +x ./cmd/server
	go build $(buildFlag) -o ./cmd/client/ ./cmd/client/
	chmod +x ./cmd/client

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: genProto
genProto:
	protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      proto/*.proto
#	protoc --go_out=./gen/go --go_opt=paths=source_relative \
#      --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
#      proto/users.proto