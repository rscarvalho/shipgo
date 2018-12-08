FROM ubuntu:bionic

RUN mkdir -p /app
WORKDIR /app

ADD ./fixtures/consignment.json /app/consignment.json
ADD ./consignment-cli /app/consignment-cli

CMD ["./consignment-cli"]
