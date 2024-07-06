package telegram

import (
	"log"
	"strings"
	"tg-bot/basketballstats/nba"
	commandconst "tg-bot/events/telegram/commandConst"
)

func (p *Processor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command %s from %s", text, username)

	switch text {
	case commandconst.AtlantaHawk:
		return p.sendMess(chatId, text)
	case commandconst.BostonCeltics:
		return p.sendMess(chatId, text)
	case commandconst.BrooklynNets:
		return p.sendMess(chatId, text)
	case commandconst.CharlotteHornets:
		return p.sendMess(chatId, text)
	case commandconst.ChicagoBulls:
		return p.sendMess(chatId, text)
	case commandconst.ClevelandCavaliers:
		return p.sendMess(chatId, text)
	case commandconst.DallasMavericks:
		return p.sendMess(chatId, text)
	case commandconst.DenverNuggets:
		return p.sendMess(chatId, text)
	case commandconst.DetroitPistons:
		return p.sendMess(chatId, text)
	case commandconst.GoldenStateWarriors:
		return p.sendMess(chatId, text)
	case commandconst.HoustonRockets:
		return p.sendMess(chatId, text)
	case commandconst.IndianaPacers:
		return p.sendMess(chatId, text)
	case commandconst.LaClippers:
		return p.sendMess(chatId, text)
	case commandconst.LosAngelesLakers:
		return p.sendMess(chatId, text)
	case commandconst.MemphisGrizzlies:
		return p.sendMess(chatId, text)
	case commandconst.MiamiHeat:
		return p.sendMess(chatId, text)
	case commandconst.MilwaukeeBucks:
		return p.sendMess(chatId, text)
	case commandconst.MinnesotaTimberwolves:
		return p.sendMess(chatId, text)
	case commandconst.NewOrleansPelicans:
		return p.sendMess(chatId, text)
	case commandconst.NewYorkKnicks:
		return p.sendMess(chatId, text)
	case commandconst.OklahomaCityThunder:
		return p.sendMess(chatId, text)
	case commandconst.OrlandoMagic:
		return p.sendMess(chatId, text)
	case commandconst.Philadelphia76ers:
		return p.sendMess(chatId, text)
	case commandconst.PhoenixSuns:
		return p.sendMess(chatId, text)
	case commandconst.PortlandTrailBlazers:
		return p.sendMess(chatId, text)
	case commandconst.SacramentoKings:
		return p.sendMess(chatId, text)
	case commandconst.SanAntonioSpurs:
		return p.sendMess(chatId, text)
	case commandconst.TorontoRaptors:
		return p.sendMess(chatId, text)
	case commandconst.UtahJazz:
		return p.sendMess(chatId, text)
	case commandconst.WashingtonWizards:
		return p.sendMess(chatId, text)
	case commandconst.ChicagoStags:
		return p.sendMess(chatId, text)
	case commandconst.StLouisBombers:
		return p.sendMess(chatId, text)
	case commandconst.ClevelandRebels:
		return p.sendMess(chatId, text)
	case commandconst.DetroitFalcons:
		return p.sendMess(chatId, text)
	case commandconst.TorontoHuskies:
		return p.sendMess(chatId, text)
	case commandconst.WashingtonCapitols:
		return p.sendMess(chatId, text)
	case commandconst.ProvidenceSteamrollers:
		return p.sendMess(chatId, text)
	case commandconst.PittsburghIronmen:
		return p.sendMess(chatId, text)
	case commandconst.BaltimoreBullets:
		return p.sendMess(chatId, text)
	case commandconst.IndianapolisJets:
		return p.sendMess(chatId, text)
	case commandconst.AndersonPackers:
		return p.sendMess(chatId, text)
	case commandconst.WaterlooHawks:
		return p.sendMess(chatId, text)
	case commandconst.IndianapolisOlympians:
		return p.sendMess(chatId, text)
	case commandconst.SheboyganRedskins:
		return p.sendMess(chatId, text)

	case commandconst.HelpCmd:
		return p.sendHelp(chatId)
	case commandconst.Start:
		return p.sendHello(chatId, username)
	default:
		return p.tg.SendMessage(chatId, msgUnknownMessage)
	}
}

func (p *Processor) sendHelp(chatId int) error {

	return p.tg.SendMessage(chatId, msgHelp)
}

func (p *Processor) sendHello(chatId int, username string) error {
	return p.tg.SendMessage(chatId, "Привет, "+username+"!"+msgHelp)
}

func (p *Processor) sendMess(chatId int, Request string) error {
	req := strings.ReplaceAll(Request, "/", "")
	nbaprocessor := nba.New()
	text := nbaprocessor.GetLastGameByTeam(req)
	return p.tg.SendMessage(chatId, text)
}
