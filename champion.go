package riotapi

import (
	"strconv"
)

// Champion service
type ChampionService struct {
	*Service
}

// Champion object
type Champion struct {
	Active            bool
	BotEnabled        bool
	BotMmEnabled      bool
	FreeToPlay        bool
	Id                int
	RankedPlayEnabled bool
}

// Returns all champions
func (service *ChampionService) All() []Champion {
	resp := service.Client.Call("champions", nil, nil)

	var champions map[string][]Champion
	service.Parser.Parse(resp.Body, &champions)
	return champions["champions"]
}

// Returns a single champion
func (service *ChampionService) One(id int) Champion {
	resp := service.Client.Call("champion", []string{strconv.Itoa(id)}, nil)

	var champion Champion
	service.Parser.Parse(resp.Body, &champion)

	return champion
}
