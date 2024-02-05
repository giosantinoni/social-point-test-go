package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test/internal/platform/server/handler/health"
	"test/internal/platform/server/handler/scores"
	"test/kit/command"
	"test/kit/query"
)

type Server struct {
	httpAddr   string
	engine     *mux.Router
	commandBus command.Bus
	queryBus   query.Bus
}

func New(host string, port uint, commandBus command.Bus, queryBus query.Bus) Server {
	srv := Server{
		engine:     mux.NewRouter(),
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		commandBus: commandBus,
		queryBus:   queryBus,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return http.ListenAndServe(s.httpAddr, s.engine)
}

func (s *Server) registerRoutes() {
	s.engine.HandleFunc("/health", health.CheckHandler()).Methods(http.MethodGet)
	s.engine.HandleFunc("/user/{userId}/score", scores.ScoreHandler(s.commandBus)).Methods(http.MethodPost)
	s.engine.HandleFunc("/ranking", scores.RankingHandler(s.queryBus)).Methods(http.MethodGet)

}
