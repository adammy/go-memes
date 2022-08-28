package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/adammy/memepen-services/pkg/httpapi"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	config  *Config
	router  *chi.Mux
	service *Service
}

func NewServer(config *Config) *Server {
	repository := NewRepository(config.RepositoryType)
	service := NewService(repository)
	return &Server{
		config:  config,
		router:  chi.NewRouter(),
		service: service,
	}
}

func (s *Server) Start() error {
	openapi, err := GetSwagger()
	if err != nil {
		return err
	}
	openapi.Servers = nil

	logger := httpapi.NewLogger(s.config.Logger)

	s.router.Use(oapimiddleware.OapiRequestValidatorWithOptions(openapi, &oapimiddleware.Options{
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			httpapi.SendErrorJSON(w, statusCode, errors.New(message))
		},
	}))
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(s.config.RequestTimeout))
	s.router.Use(httplog.RequestLogger(logger))

	HandlerFromMux(s, s.router)

	logger.Info().Msgf("server starting on port %d", s.config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router); err != nil {
		return err
	}

	return nil
}

func (s *Server) GetTemplate(w http.ResponseWriter, r *http.Request, templateID string) {
	template, err := s.service.GetTemplate(r.Context(), templateID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusNotFound, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}
	httpapi.SendJSON(w, http.StatusOK, template)
}

func (s *Server) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	var create CreateTemplate
	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
		return
	}

	template, err := s.service.CreateTemplate(r.Context(), create)
	if err != nil {
		if errors.Is(err, httpapi.ErrBadRequest) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}
	httpapi.SendJSON(w, http.StatusCreated, template)
}

func (s *Server) UpdateTemplate(w http.ResponseWriter, r *http.Request, templateID string) {
	var create CreateTemplate
	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
		return
	}

	template, err := s.service.UpdateTemplate(r.Context(), templateID, create)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusNotFound, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}

	httpapi.SendJSON(w, http.StatusOK, template)
}

func (s *Server) DeleteTemplate(w http.ResponseWriter, r *http.Request, templateID string) {
	if err := s.service.DeleteTemplate(r.Context(), templateID); err != nil {
		if errors.Is(err, ErrNotFound) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusNotFound, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}
	httpapi.SendJSON(w, http.StatusNoContent, nil)
}
