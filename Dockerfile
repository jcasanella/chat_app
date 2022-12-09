FROM golang:1.19.4-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY config/ config/
COPY server/ server/

RUN go mod download && go build -o chat_app

EXPOSE 8080

CMD [ "./chat_app" ]