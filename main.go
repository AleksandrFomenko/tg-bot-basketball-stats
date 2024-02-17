package main

import (
	"flag"
	"log"
	tgclient "tg-bot/clients/telegram"
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
	eventsProcessor := telegram.New(tgclient.New(tgHost, mustToken()))
	log.Print("service started")
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("token-bt", "", "токен для запуска")
	flag.Parse()
	if *token == "" {
		log.Fatal("некорректный токен")
	}
	return *token
}
