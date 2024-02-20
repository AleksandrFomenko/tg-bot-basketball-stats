package nba

type Game struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	Season           int    `json:"season"`
	Status           string `json:"status"`
	Period           int    `json:"period"`
	Time             string `json:"time"`
	Postseason       bool   `json:"postseason"`
	HomeTeamScore    int    `json:"home_team_score"`
	VisitorTeamScore int    `json:"visitor_team_score"`
	HomeTeam         struct {
		ID           int    `json:"id"`
		Conference   string `json:"conference"`
		Division     string `json:"division"`
		City         string `json:"city"`
		Name         string `json:"name"`
		FullName     string `json:"full_name"`
		Abbreviation string `json:"abbreviation"`
	} `json:"home_team"`
	VisitorTeam struct {
		ID           int    `json:"id"`
		Conference   string `json:"conference"`
		Division     string `json:"division"`
		City         string `json:"city"`
		Name         string `json:"name"`
		FullName     string `json:"full_name"`
		Abbreviation string `json:"abbreviation"`
	} `json:"visitor_team"`
}
