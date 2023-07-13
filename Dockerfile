FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache git \
    && go mod download \
    && go build -o app .

CMD ["./app"]
