FROM golang:1.16 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/lgylgy/rinkgo

COPY go.mod .
COPY go.sum .

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -o pstat ./cmd/pstat/

FROM alpine:latest
ARG DBURL
ENV DBURL ${DBURL}

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/lgylgy/rinkgo/pstat .

EXPOSE 5000

CMD ./pstat --url ${DBURL}