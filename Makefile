GO = CGO_ENABLED=0 go
BINARY_NAME=analytics-server

gen-protobuf:
	protoc --proto_path=proto/ proto/*.proto --go_out=. --go-grpc_out=.

build-osx: 
	GOOS=darwin GOARCH=amd64 $(GO) build -o build/$(BINARY_NAME)-osx main.go

build-windows:
	GOARCH=386 GOOS=window go build -o ${BINARY_NAME}-windows main.go

build-heroku:
	go build -o bin/${BINARY_NAME} -v main.go

run-server:
	./build/$(BINARY_NAME)-osx

clean:
	rm -rf build
