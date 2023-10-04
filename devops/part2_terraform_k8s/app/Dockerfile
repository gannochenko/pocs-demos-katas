FROM golang:1.18-alpine as builder

WORKDIR /root

ARG GITHUB_USERNAME
ARG GITHUB_TOKEN

RUN apk add --no-cache --update git

RUN echo "machine github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" > ~/.netrc
RUN echo "machine api.github.com login ${GITHUB_USERNAME} password ${GITHUB_TOKEN}" >> ~/.netrc

ENV GO111MODULE "on"

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . /build
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o app ./cmd/main.go

FROM alpine:3.15.0 as base
COPY --from=builder /build/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
