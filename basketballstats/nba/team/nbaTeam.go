package team

import "errors"

type NBATeams struct {
	teamsMap map[string]NbaTeam
}
type NbaTeam struct {
	id int
}

func New() NBATeams {
	return NBATeams{teamsMap()}
}

func teamsMap() map[string]NbaTeam {
	mapIdTeams := make(map[string]NbaTeam)
	mapIdTeams["atlantahawk"] = NbaTeam{id: 1}
	mapIdTeams["bostonceltics"] = NbaTeam{id: 2}
	mapIdTeams["brooklynnets"] = NbaTeam{id: 3}
	mapIdTeams["charlottehornets"] = NbaTeam{id: 4}
	mapIdTeams["chicagobulls"] = NbaTeam{id: 5}
	mapIdTeams["clevelandcavaliers"] = NbaTeam{id: 6}
	mapIdTeams["dallasmavericks"] = NbaTeam{id: 7}
	mapIdTeams["denvernuggets"] = NbaTeam{id: 8}
	mapIdTeams["detroitpistons"] = NbaTeam{id: 9}
	mapIdTeams["goldenstatewarriors"] = NbaTeam{id: 10}
	mapIdTeams["houstonrockets"] = NbaTeam{id: 11}
	mapIdTeams["indianapacers"] = NbaTeam{id: 12}
	mapIdTeams["laclippers"] = NbaTeam{id: 13}
	mapIdTeams["losangeleslakers"] = NbaTeam{id: 14}
	mapIdTeams["memphisgrizzlies"] = NbaTeam{id: 15}
	mapIdTeams["miamiheat"] = NbaTeam{id: 16}
	mapIdTeams["milwaukeebucks"] = NbaTeam{id: 17}
	mapIdTeams["minnesotatimberwolves"] = NbaTeam{id: 18}
	mapIdTeams["neworleanspelicans"] = NbaTeam{id: 19}
	mapIdTeams["newyorkknicks"] = NbaTeam{id: 20}
	mapIdTeams["oklahomacitythunder"] = NbaTeam{id: 21}
	mapIdTeams["orlandomagic"] = NbaTeam{id: 22}
	mapIdTeams["philadelphia76ers"] = NbaTeam{id: 23}
	mapIdTeams["phoenixsuns"] = NbaTeam{id: 24}
	mapIdTeams["portlandtrailblazers"] = NbaTeam{id: 25}
	mapIdTeams["sacramentokings"] = NbaTeam{id: 26}
	mapIdTeams["sanantoniospurs"] = NbaTeam{id: 27}
	mapIdTeams["torontoraptors"] = NbaTeam{id: 28}
	mapIdTeams["utahjazz"] = NbaTeam{id: 29}
	mapIdTeams["washingtonwizards"] = NbaTeam{id: 30}
	mapIdTeams["chicagostags"] = NbaTeam{id: 37}
	mapIdTeams["stlouisbombers"] = NbaTeam{id: 38}
	mapIdTeams["clevelandrebels"] = NbaTeam{id: 39}
	mapIdTeams["detroitfalcons"] = NbaTeam{id: 40}
	mapIdTeams["torontohuskies"] = NbaTeam{id: 41}
	mapIdTeams["washingtoncapitols"] = NbaTeam{id: 42}
	mapIdTeams["providencesteamrollers"] = NbaTeam{id: 43}
	mapIdTeams["pittsburghironmen"] = NbaTeam{id: 44}
	mapIdTeams["baltimorebullets"] = NbaTeam{id: 45}
	mapIdTeams["indianapolisjets"] = NbaTeam{id: 46}
	mapIdTeams["andersonpackers"] = NbaTeam{id: 47}
	mapIdTeams["waterloohawks"] = NbaTeam{id: 48}
	mapIdTeams["indianapolisolympians"] = NbaTeam{id: 49}
	mapIdTeams["sheboyganredskins"] = NbaTeam{id: 51}
	return mapIdTeams
}

func (n NBATeams) FoundIDs(name string) (int, error) {
	if _, ok := n.teamsMap[name]; !ok {
		return 0, errors.New("Teams not found")
	}
	return n.teamsMap[name].id, nil
}
