package main

import (
	"log"
	tgclient "tg-bot/clients/telegram"
	"tg-bot/config"
	eventconsumer "tg-bot/consumer/event-consumer"
	"tg-bot/events/telegram"
)

/*
	help - вызвать справку
	selectteam - выбрать команду
*/

const (
	tgHost    = "api.telegram.org"
	batchSize = 100
)

func main() {
	cfg := config.LoadConfig()
	if cfg.Token == "" {
		log.Fatal("некорректный токен")
	}
	eventsProcessor := telegram.New(tgclient.New(tgHost, cfg.Token))
	log.Print("service started")
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
