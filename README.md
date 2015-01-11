Riot Games API client
=======================

Example
----------
	api := riotapi.NewApi("API_KEY")

	champions := api.Champion.All()

	for _, c := range champions {
		fmt.Printf("%d,", c.Id)
	}