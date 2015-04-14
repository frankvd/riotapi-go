package riotapi

import (
	"log"
	"net/http"
	"strings"
)

// Service struct which is embedded in all services
type Service struct {
	Parser Parser
	Client Client
}

// Struct which groups all services
type Api struct {
	Champion     ChampionService
	Game         GameService
	League       LeagueService
	StaticData   StaticDataService
	Status       StatusService
	Match        MatchService
	MatchHistory MatchHistoryService
	Stats        StatsService
	Summoner     SummonerService
	Team         TeamService
}

// Initialize all services
func NewApi(apikey string) *Api {
	client := NewClient()
	client.ApiKey = apikey
	parser := new(JsonParser)

	api := Api{
		Champion:     ChampionService{Service: &Service{Client: client, Parser: parser}},
		League:       LeagueService{Service: &Service{Client: client, Parser: parser}},
		Summoner:     SummonerService{Service: &Service{Client: client, Parser: parser}},
		MatchHistory: MatchHistoryService{Service: &Service{Client: client, Parser: parser}},
		Match:        MatchService{Service: &Service{Client: client, Parser: parser}},
	}

	return &api
}

type Client interface {
	Call(endpoint string, params ...string) *http.Response
}

// Http client
type HttpClient struct {
	Endpoints map[string]string
	ApiKey    string
	Region    string
	BaseUrl   string
}

// Initialize a new client
func NewClient() *HttpClient {
	client := HttpClient{
		BaseUrl: "https://euw.api.pvp.net",
		ApiKey:  "apikey",
		Region:  "euw",
		Endpoints: map[string]string{
			"champions":        "/api/lol/{region}/v1.2/champion",
			"champion":         "/api/lol/{region}/v1.2/champion/{param}",
			"recent-games":     "/api/lol/{region}/v1.1/game/by-summoner/{param}/recent",
			"league":           "/api/lol/{region}/v2.5/league/by-summoner/{param}",
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
		},
	}

	return &client
}

// Call an API endpoint
func (client *HttpClient) Call(endpoint string, params ...string) *http.Response {
	url := client.createUrl(endpoint, params)
	log.Printf("Calling: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

// Compose an URL
func (client *HttpClient) createUrl(endpoint string, params []string) string {
	resourceUrl := client.Endpoints[endpoint]
	resourceUrl = strings.Replace(resourceUrl, "{region}", client.Region, 1)
	for _, value := range params {
		resourceUrl = strings.Replace(resourceUrl, "{param}", value, 1)
	}
	return client.BaseUrl + resourceUrl + "?api_key=" + client.ApiKey
}
