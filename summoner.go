package riotapi

import "strconv"

// Summoner service
type SummonerService struct {
	*Service
}

// Summoner
type Summoner struct {
	Name          string `json:name`
	ID            int
	SummonerLevel int
	ProfileIconID int
	RevisionDate  int
}

// Find a summoner by his/her name
func (service *SummonerService) ByName(name string) Summoner {
	resp := service.Client.Call("summoner-by-name", name)

	summoner := map[string]Summoner{}
	service.Parser.Parse(resp.Body, &summoner)

	return summoner[name]
}

// Returns the masteries of the summoner
func (service *SummonerService) Masteries(id int) Masteries {
	textId := strconv.Itoa(id)
	resp := service.Client.Call("masteries", textId)
	masteries := map[string]Masteries{}
	service.Parser.Parse(resp.Body, &masteries)

	return masteries[textId]
}

// Returns the runes of the summoner
func (service *SummonerService) Runes(id int) Runes {
	textId := strconv.Itoa(id)
	resp := service.Client.Call("runes", textId)
	runes := map[string]Runes{}
	service.Parser.Parse(resp.Body, &runes)

	return runes[textId]
}
