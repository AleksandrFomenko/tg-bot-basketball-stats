package external

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"tg-bot/config"
	"time"
)

type balldontlieAPI struct {
	client http.Client
	apiKey string
}

const (
	curl = "https://api.balldontlie.io/v1/games"
)

func New() *balldontlieAPI {
	return &balldontlieAPI{
		client: http.Client{},
		apiKey: config.LoadConfig().Key,
	}
}

func (b *balldontlieAPI) getGameByDate(ctx context.Context, teamId int, date time.Time) {
	urlParams := url.Values{
		"team_ids": []string{strconv.Itoa(teamId)},
		"dates[]":  []string{date.Format("2006-01-02")},
	}
	req, err := http.NewRequest(http.MethodGet, curl+"?"+urlParams.Encode(), nil)
	req.Header.Add("Authorization", b.apiKey)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
	}
}
