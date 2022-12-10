GO_WORKSPACE := $(GOPATH)/src
clean:
	rm pb/*

gen :
	protoc --proto_path=proto proto/*.proto \
	--go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo "finish run protoc compiler"

runtest:
	go test -timeout 30s gocpu/serializer -run ^TestFileSerializer$