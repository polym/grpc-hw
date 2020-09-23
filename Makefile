protoc:
	protoc -I  protoc/ hello.proto --go_out=plugins=grpc:pkg/helloworld/

app:
	GOOS=darwin GOARCH=amd64 go build -o build/client cmd/client/*.go
	go build -o build/server cmd/server/*.go

.PHONY: protoc
