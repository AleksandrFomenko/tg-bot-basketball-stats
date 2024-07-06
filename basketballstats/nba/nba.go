package nba

import (
	"encoding/json"
	"fmt"
	"io"
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
	usa, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	timeUsa := currentTime.In(usa)

	param.Set("team_ids[]", strconv.Itoa(teamId))

	param.Set("dates[]", timeUsa.Format("2006-01-02"))
	if err != nil {
		e.Wrap("Ошибка при выполнении запроса:", err)
	}

	client := http.Client{}
	maxAttempts := 20
	attempts := 0
	for attempts < maxAttempts {
		req, err := http.NewRequest(http.MethodGet, urlReq+"?"+param.Encode(), nil)
		req.Header.Add("Authorization", "d1a5a405-1265-4bce-8214-7b39d6dec2d6")
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return fmt.Sprintf("Ошибка при чтении ответа: %v", err)
		}

		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return fmt.Sprintf("Ошибка при чтении ответа: %v", err)
		}
		fmt.Println("Длина данных:", len(response.Data))
		fmt.Println("Ответ сервера:", response)
		fmt.Println(currentTime.Format("2006-01-02"))
		r, _ := io.ReadAll(resp.Body)
		//fmt.Println(string(r))
		if err := json.Unmarshal(r, &response); err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
		}
		if len(response.Data) == 0 {
			timeUsa = timeUsa.Add(-24 * time.Hour)
			param.Set("dates[]", timeUsa.Format("2006-01-02"))
			attempts++

			continue
		}

		if response.Data[0].HomeTeamScore == 0 {
			response.Data = []Game{}
			timeUsa = timeUsa.Add(-24 * time.Hour)
			param.Set("dates[]", timeUsa.Format("2006-01-02"))
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Ошибка при выполнении запроса. Код статусаa: %v", resp.StatusCode)
		}
		break
	}
	if len(response.Data) == 0 {
		return "команда не играла последние 20 дней"
	}

	text := fmt.Sprintf("Дата: %s\nСезон: %d\nКоманда дома: %s (%s), конференция: %s, дивизион: %s \nКоманда в гостях: %s (%s),  конференция: %s, дивизион: %s\nCчет: %d:%d",
		response.Data[0].Date,
		response.Data[0].Season,
		response.Data[0].HomeTeam.FullName,
		response.Data[0].HomeTeam.Abbreviation,
		response.Data[0].HomeTeam.Conference,
		response.Data[0].HomeTeam.Division,
		response.Data[0].VisitorTeam.FullName,
		response.Data[0].VisitorTeam.Abbreviation,
		response.Data[0].VisitorTeam.Conference,
		response.Data[0].VisitorTeam.Division,
		response.Data[0].HomeTeamScore,
		response.Data[0].VisitorTeamScore)

	return text

}
