package team

import "errors"

type NBATeam struct {
	teamsMap map[string]int
}

func New() NBATeam {
	return NBATeam{teamsMap()}
}

func teamsMap() map[string]int {
	mymap := make(map[string]int)
	mymap["boston"] = 2
	return mymap
}

func (n NBATeam) FoundIDs(name string) (int, error) {
	if _, ok := n.teamsMap[name]; !ok {
		return 0, errors.New("Teams not found")
	}
	return n.teamsMap[name], nil
}
