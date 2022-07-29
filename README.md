# echo-server
Echo server in GO for protobuffers implementation. It just prints though console what's being received. In this case it's being using a specific proto file for events.

### To deploy locally in your machine: 

Install GO: https://go.dev/doc/install

#### To install protobuffer compiler. If you are using another platform please go here: https://grpc.io/docs/languages/go/quickstart/
```
brew install protobuf
```

#### To install Go dependencies: 
```
go install http://google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install http://google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Project has a Makefile but is not yet working completely: 

#### To generate protobufers (not needed if there are no new changes) this command still fails. You can maybe run it locally.
```
make gen-protobuf
```
#### To build 
```
make build-osx 
```

#### To run the server locally
```
make run-server
```

#### For cleaning build folder
```
make clean
```

Example running on CTR: 

![image](https://user-images.githubusercontent.com/91892711/181766826-12263a39-5593-42f5-92b8-436d48672bc1.png)
