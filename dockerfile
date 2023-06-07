# Dockerfile

FROM golang:1.20.3-alpine3.17 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o benchmark ./client/cmd/benchmark && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

FROM --platform=linux/arm64 alpine:3.17
COPY --from=build /app/benchmark /bin/benchmark

CMD ["benchmark"]