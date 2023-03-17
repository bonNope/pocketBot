FROM golang:1.19.5-alpine3.17 AS builder

COPY . /github.com/bonNope/pocketBot/
WORKDIR /github.com/bonNope/pocketBot/

RUN go mod download
RUN go build -o bot cmd/bot_app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/bonNope/pocketBot/bot .
COPY --from=0 /github.com/bonNope/pocketBot/configs configs/
COPY --from=0 /github.com/bonNope/pocketBot/.env .

EXPOSE 60

CMD ["./bot"]