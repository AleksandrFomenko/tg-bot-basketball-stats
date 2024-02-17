package nbaapi

import (
	"io"
	"net/http"
)

func GetTeam() string {

	url := "https://basketapi1.p.rapidapi.com/api/basketball/search/kevin"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "9440d7dfcfmsh5b5348b61c08b8bp10474fjsn4ab01c6dde22")
	req.Header.Add("X-RapidAPI-Host", "basketapi1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	return string(body)

}
