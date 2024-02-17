# Build stage
FROM golang:1.22.0 AS builder

WORKDIR /build

COPY src .

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go mod download

RUN go build -o go-app


# Deploy
# FROM golang:alpine3.19 AS deploy
FROM alpine:3.19.1 AS deploy

# https://stackoverflow.com/a/59367690/3310235
RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY src/conf/ conf/

COPY src/views/ views/

COPY --from=builder /build/go-app .

RUN chmod +x go-app

EXPOSE 8010

CMD ["./go-app"]