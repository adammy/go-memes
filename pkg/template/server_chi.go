package template

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adammy/memepen-services/pkg/httpapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

var _ httpapi.Server = (*chiServer)(nil)

type chiServer struct {
	config *Config
	router *chi.Mux
}

// NewChiServer returns a new Server utilizing the gin framework.
func NewChiServer(cfg *Config) (*chiServer, error) {
	r := chi.NewRouter()

	return &chiServer{
		config: cfg,
		router: r,
	}, nil
}

func (s *chiServer) Start() error {
	logger := httpapi.NewLogger(s.config.Logger)

	s.router.Use(middleware.Heartbeat("/ping"))
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(20 * time.Second))
	s.router.Use(httplog.RequestLogger(logger))

	s.registerRoutes()

	if err := http.ListenAndServe(":"+strconv.Itoa(s.config.Port), s.router); err != nil {
		return err
	}

	return nil
}

func (s *chiServer) registerRoutes() {
	s.router.Route("/v1", func(r chi.Router) {
		r.Get("/hello", s.otherHandler)
	})
}

func (s *chiServer) otherHandler(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	oplog.Info().Msg("another msg with context??")
	_ = httpapi.SendJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
	//_ = httpapi.SendErrorJSON(w, http.StatusInternalServerError, errors.New("muh error"))
}
