# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.4-alpine3.17 AS build

ENV CHAT_APP=chat_app
WORKDIR /app

COPY . .

RUN go mod download && go build -o ${CHAT_APP}

## Deploy
FROM alpine:3.17

ENV CHAT_APP=chat_app
WORKDIR /app

COPY --from=build /app/${CHAT_APP} .
COPY --from=build /app/config/*.yaml config/local.yaml

RUN addgroup -S ${CHAT_APP} && adduser -S ${CHAT_APP} -G ${CHAT_APP}
USER ${CHAT_APP}

EXPOSE 8080

ENTRYPOINT ./${CHAT_APP}