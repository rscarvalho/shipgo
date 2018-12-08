GOPATH := $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)

build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/rscarvalho/shipgo/consignment-service \
		proto/consignment/consignment.proto