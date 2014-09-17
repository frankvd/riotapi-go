package riotapi

import (
	"net/http"
	"strconv"
	"strings"
)

var (
	parser    = new(JsonParser)
	apiKey    string
	region    string
	baseUrl   = "http://euw.api.pvp.net"
	endpoints = map[string]string{
		"champion":         "/api/lol/{region}/v1.1/champion",
		"recent-games":     "/api/lol/{region}/v1.1/game/by-summoner/{param}/recent",
		"league":           "/api/lol/{region}/v2.2/league/by-summoner/{param}",
		"summary":          "/api/lol/{region}/v1.2/stats/by-summoner/{param}/summary",
		"ranked-stats":     "/api/lol/{region}/v1.2/stats/by-summoner/{param}/ranked",
		"masteries":        "/api/lol/{region}/v1.4/summoner/{param}/masteries",
		"runes":            "/api/lol/{region}/v1.4/summoner/{param}/runes",
		"summoner-by-name": "/api/lol/{region}/v1.4/summoner/by-name/{param}",
		"summoner-by-id":   "/api/lol/{region}/v1.4/summoner/{param}",
		"summoner-names":   "/api/lol/{region}/v1.4/summoner/{param}/name",
		"team":             "/api/lol/{region}/v2.2/team/by-summoner/{param}",
		"match-history":    "/api/lol/{region}/v2.2/matchhistory/{param}",
		"match":            "/api/lol/{region}/v2.2/match/{param}",
	}
)

func SetApiKey(k string) {
	apiKey = k
}
func SetRegion(r string) {
	region = r
}
func SetParser(p *JsonParser) {
	parser = p
}

func SummonerByName(name string) Summoner {
	resp := Call("summoner-by-name", name)
	summoner := map[string]Summoner{}
	parser.Parse(resp.Body, &summoner)

	return summoner[name]
}

func SummonerMasteries(id int) Masteries {
	textId := strconv.Itoa(id)
	resp := Call("masteries", textId)
	masteries := map[string]Masteries{}
	parser.Parse(resp.Body, &masteries)

	return masteries[textId]
}

func SummonerRunes(id int) Runes {
	textId := strconv.Itoa(id)
	resp := Call("runes", textId)
	runes := map[string]Runes{}
	parser.Parse(resp.Body, &runes)

	return runes[textId]
}

func Call(endpoint string, params ...string) *http.Response {
	resp, err := http.Get(createUrl(endpoint, params))
	if err != nil {
		panic(err)
	}
	return resp
}

func createUrl(endpoint string, params []string) string {
	resourceUrl := endpoints[endpoint]
	resourceUrl = strings.Replace(resourceUrl, "{region}", region, 1)
	for _, value := range params {
		resourceUrl = strings.Replace(resourceUrl, "{param}", value, 1)
	}
	return baseUrl + resourceUrl + "?api_key=" + apiKey
}
