package basketballstats

import "tg-bot/basketballstats/nba"

type Worker interface {
	GetLastGameByTeam(string) string
}

type basketballstats struct {
	Worker Worker
}

func New() basketballstats {
	return basketballstats{
		nba.New(),
	}
}
