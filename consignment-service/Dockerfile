FROM golang:1.10.0 as builder

WORKDIR /go/src/github.com/rscarvalho/shipgo/consignment-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder go/src/github.com/rscarvalho/shipgo/consignment-service/consignment-service .

EXPOSE 50051

CMD ["./consignment-service"]