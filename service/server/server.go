package server

import (
	"context"
	"fmt"
	"github/yudgxe/leadgen.market/service/api"
	"net/http"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
}

func New(host, port string) (*Server, error) {
	router, err := api.NewRouter()
	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: router,
	}

	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (s *Server) Start() {
	go func() {
		log.Debugf("Sever start at: %s", s.Addr)

		if err := s.ListenAndServe(); err != http.ErrServerClosed && err != nil {
			log.Errorf("Error on server listen - %s", err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	if err := s.Shutdown(context.Background()); err != nil {
		log.Errorf("Error on server shutdown - %s", err)
	}

	log.Debug("Server gracefully shutdown")
}
