.PHONY:

build:
	go build -o bot cmd/bot_app/main.go

run: build
	./bot

build-image:
	docker build -t telegram-bot:v0.1 .

start-container:
	docker run --name=telegram-bot -p 60:60 --env-file=.env telegram-bot:v0.1