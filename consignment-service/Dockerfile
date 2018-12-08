FROM ubuntu:bionic

RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service

EXPOSE 50051

CMD ["./consignment-service"]