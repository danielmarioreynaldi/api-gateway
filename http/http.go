package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/danielmarioreynaldi/api-gateway/config"
	"github.com/go-chi/chi"
)

type HttpServer struct {
	Server *http.Server
	Router *chi.Mux
	Addr   string
}

func NewHttpServer(httpCfg config.HttpCfg) *HttpServer {
	addr := fmt.Sprintf("%s:%d", httpCfg.Host, httpCfg.Port)
	router := chi.NewRouter()

	httpServer := &HttpServer{
		Router: router,
		Addr:   addr,
	}

	return httpServer
}

func (h *HttpServer) Start() {
	h.Server = &http.Server{
		Addr:    h.Addr,
		Handler: h.Router,
	}

	if err := h.Server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Cant start serving HTTP server")
		}

	}
}

func (h *HttpServer) Stop() {
	c := context.Background()
	err := h.Server.Shutdown(c)
	if err != nil {
		log.Fatal("Failed to shutdown HTTP server")
	}
}
