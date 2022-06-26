FROM golang:1.16 as base

FROM base as dev

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN apt-get update
RUN apt-get install librdkafka-dev -y
CMD ["air"]