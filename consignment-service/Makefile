GOPATH := $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)

build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/rscarvalho/shipgo/consignment-service \
		proto/consignment/consignment.proto

	docker build -t consignment-service .

run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		consignment-service