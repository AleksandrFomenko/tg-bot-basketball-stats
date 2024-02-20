package nba

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"tg-bot/basketballstats/nba/team"
	"tg-bot/lib/e"
	"time"
)

type NbaProcessor struct {
}

var response struct {
	Data []Game `json:"data"`
}

func New() NbaProcessor {
	return NbaProcessor{}
}

func (n NbaProcessor) GetLastGameByTeam(teamName string) string {

	urlReq := "http://api.balldontlie.io/v1/games"
	param := url.Values{}
	teamId, err := team.New().FoundIDs(teamName)

	if err != nil {
		e.Wrap("команда не найдена", err)
	}
	currentTime := time.Now()

	param.Set("team_ids[]", strconv.Itoa(teamId))
	param.Add("dates[]", currentTime.Format("2006-01-02"))

	if err != nil {
		e.Wrap("Ошибка при выполнении запроса:", err)
		return fmt.Sprintf("Ошибка при выполнении запроса: %v", err)
	}

	client := http.Client{}
	maxAttempts := 20 // Максимальное количество попыток
	attempts := 0
	for attempts < maxAttempts {
		req, err := http.NewRequest(http.MethodGet, urlReq, nil)
		req.Header.Add("Authorization", "d1a5a405-1265-4bce-8214-7b39d6dec2d6")
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return fmt.Sprintf("Ошибка при чтении ответа: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return fmt.Sprintf("Ошибка при чтении ответа: %v", err)
		}
		defer resp.Body.Close()
		if len(response.Data) == 0 {
			currentTime = currentTime.Add(-24 * time.Hour)
			param.Set("dates[]", currentTime.Format("2006-01-02"))
			attempts++
			continue
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Ошибка при выполнении запроса. Код статуса: %v", resp.StatusCode)
		}

		break
	}
	return string(currentTime.Format("2006-01-02"))
}
