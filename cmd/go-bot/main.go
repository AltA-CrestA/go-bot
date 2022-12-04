package main

import (
	"flag"
	eventconsumer "github.com/AltA-CrestA/go-bot/internal/consumer/event-consumer"
	"github.com/AltA-CrestA/go-bot/internal/events/telegram"
	"github.com/AltA-CrestA/go-bot/internal/storage/files"
	tgClient "github.com/AltA-CrestA/go-bot/pkg/clients/telegram"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventProcessor, eventProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service has stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
