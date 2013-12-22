package riotapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	parser    = new(JsonParser)
	apiKey    string
	region    string
	baseUrl   = "http://prod.api.pvp.net"
	endpoints = map[string]string{
		"champion":         "/api/lol/{region}/v1.1/champion",
		"recent-games":     "/api/lol/{region}/v1.1/game/by-summoner/{param}/recent",
		"league":           "/api/lol/{region}/v2.2/league/by-summoner/{param}",
		"summary":          "/api/lol/{region}/v1.2/stats/by-summoner/{param}/summary",
		"ranked-stats":     "/api/lol/{region}/v1.2/stats/by-summoner/{param}/ranked",
		"masteries":        "/api/lol/{region}/v1.2/summoner/{param}/masteries",
		"runes":            "/api/lol/{region}/v1.2/summoner/{param}/runes",
		"summoner-by-name": "/api/lol/{region}/v1.1/summoner/by-name/{param}",
		"summoner-by-id":   "/api/lol/{region}/v1.2/summoner/{param}",
		"summoner-names":   "/api/lol/{region}/v1.2/summoner/{param}/name",
		"team":             "/api/lol/{region}/v2.2/team/by-summoner/{param}",
	}
)

func SetApiKey(k string) {
	apiKey = k
}
func SetRegion(r string) {
	region = r
}
func SetParser(p Parser) {
	parser = p
}

func Call(endpoint string, params ...string) interface{} {
	resp, _ := http.Get(createUrl(endpoint, params))
	body, _ := ioutil.ReadAll(resp.Body)

	return parser.Parse(resp)
}

func createUrl(endpoint string, params []string) string {
	resourceUrl := endpoints[endpoint]
	resourceUrl = strings.Replace(resourceUrl, "{region}", region, 1)
	for _, value := range params {
		resourceUrl = strings.Replace(resourceUrl, "{param}", value, 1)
	}
	return baseUrl + resourceUrl + "?api_key=" + apiKey
}
