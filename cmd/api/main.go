package main

import (
	"os"
	"strconv"
	"test/internal/app/score/command"
	"test/internal/app/score/ranking/query"
	inmemorybus "test/internal/platform/bus/inmemory"
	"test/internal/platform/server"
	inmemory "test/internal/platform/storage/inmemory"
)

func main() {

	var (
		host    = os.Getenv("HOST")
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	)
	var (
		userRepository = inmemory.NewUserRepository()
	)

	var (
		commandBus = inmemorybus.NewCommandBus()
		queryBus   = inmemorybus.NewQueryBus()
	)

	incrementScoreCommandHandler := command.NewUpdateUserScoreCommandHandler(userRepository)
	commandBus.Register(command.UpdateUserScoreCommandType, incrementScoreCommandHandler)

	newScoreCommandHandler := command.NewNewUserScoreCommandHandler(userRepository)
	commandBus.Register(command.NewUserScoreCommandType, newScoreCommandHandler)

	absoluteQueryHandler := query.NewAbsoluteRankingQueryHandler(userRepository)
	queryBus.Register(query.AbsoluteRankingQueryType, absoluteQueryHandler)

	relativeQueryHandler := query.NewRelativeRankingQueryHandler(userRepository)
	queryBus.Register(query.RelativeRankingQueryType, relativeQueryHandler)

	srv := server.New(host, uint(port), commandBus, queryBus)
	err := srv.Run()
	if err != nil {
		return
	}
}
