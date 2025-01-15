FROM golang:1.23.4-alpine

ENV GO111MODULE=on

WORKDIR '/app'

COPY go.mod go.sum ./

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o main ./cmd/server

CMD [ "./main" ]