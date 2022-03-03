	teamRepository := repository.NewTeamRepository(database)
	competitionRepository := repository.NewCompetitionRepository(database)
	matchRepository := repository.NewMatchRepository(database)
	apiClient := client.NewApiFootballClient()

	// _ := service.NewUpdateService(teamRepository, competitionRepository, apiClient)
	updateService := service.NewUpdateService(teamRepository, competitionRepository, matchRepository, apiClient)
	// updateService.ImportCompetitionsAndSeasons()
	// updateService.ImportTeamsAndStadiums()
	updateService.ImportMatches("Brazil", 2020, true)

	defer destroy()
	//image.ConvertSvgToPngWithChrome("https://s.glbimg.com/es/sde/f/organizacoes/2020/02/12/botsvg.svg", "./assets/test.png")
