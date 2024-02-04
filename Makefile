port = 8080
buildFlag = -ldflags="-X 'main.buildVersion=`git describe --tags --abbrev=0`' -X 'main.buildDate=`date`'"

.PHONY: runServer
runServer: build
	./cmd/server/mac_server

.PHONY: runClient
runClient: build
	./cmd/client/mac_client

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

.PHONY: build
build:
	GOOS=darwin GOARCH=arm64 go build $(buildFlag) -o ./cmd/server/mac_server ./cmd/server
	GOOS=windows GOARCH=amd64 go build $(buildFlag) -o ./cmd/server/win_server.exe ./cmd/server
	GOOS=linux GOARCH=amd64 go build $(buildFlag) -o ./cmd/server/linux_server ./cmd/server
	chmod +x ./cmd/server/mac_server
	GOOS=darwin GOARCH=arm64 go build $(buildFlag) -o ./cmd/client/mac_client ./cmd/client
	GOOS=windows GOARCH=arm64 go build $(buildFlag) -o ./cmd/client/win_client.exe ./cmd/client
	GOOS=linux GOARCH=amd64 go build $(buildFlag) -o ./cmd/client/linux_client ./cmd/client
	chmod +x ./cmd/client/mac_client

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

.PHONY: genMock
genMock:
	mockgen -source=internal/server/repositories/interface.go \
    -destination=internal/server/repositories/mocks/mock_repository.go && \
    mockgen -source=internal/server/usecase/interface.go \
    -destination=internal/server/usecase/mocks/mock_usecase.go