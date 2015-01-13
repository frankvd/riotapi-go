package riotapi

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
	resp := service.Client.Call("champion")

	var champions map[string][]Champion
	service.Parser.Parse(resp.Body, &champions)
	return champions["champions"]
}
