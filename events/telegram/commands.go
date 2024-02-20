package telegram

import (
	"log"
	"strings"
	"tg-bot/basketballstats/nba"
)

const (
	selectTeam = "/selectteam"
	helpCmd    = "/help"
	start      = "/start"
)

func (p *Processor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command %s from %s", text, username)

	switch text {
	case selectTeam:
		return p.sendMess(chatId)
	case helpCmd:
		return p.sendHelp(chatId)
	case start:
		return p.sendHello(chatId, username)
	default:
		return p.tg.SendMassage(chatId, msgUnknownMessage)
	}
}

func (p *Processor) sendHelp(chatId int) error {
	return p.tg.SendMassage(chatId, msgHelp)
}

func (p *Processor) sendHello(chatId int, username string) error {
	return p.tg.SendMassage(chatId, "Привет, "+username+"!"+msgHelp)
}

func (p *Processor) sendMess(chatId int) error {
	nbaprocessor := nba.New()
	text := nbaprocessor.GetLastGameByTeam("boston")
	return p.tg.SendMassage(chatId, text)
}
