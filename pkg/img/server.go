package img

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"net/http"
	"strings"

	"github.com/adammy/memepen-services/pkg/httpapi"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	cfg    *Config
	router *chi.Mux
	svc    *Service
}

func NewServer(cfg *Config) *Server {
	repository := NewRepository(cfg.RepositoryType)
	uploader := NewUploader(cfg.UploaderType, cfg.UploaderBasePath)
	svc := NewService(repository, uploader, cfg.BaseURL, cfg.MaxWidth, cfg.MaxHeight)
	return &Server{
		cfg:    cfg,
		router: chi.NewRouter(),
		svc:    svc,
	}
}

func (s *Server) Start() error {
	openapi3filter.RegisterBodyDecoder(httpapi.ImagePNG, openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder(httpapi.ImageJPG, openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder(httpapi.ImageJPEG, openapi3filter.FileBodyDecoder)

	openapi, err := GetSwagger()
	if err != nil {
		return err
	}
	openapi.Servers = nil

	logger := httpapi.NewLogger(s.cfg.Logger)

	s.router.Use(oapimiddleware.OapiRequestValidatorWithOptions(openapi, &oapimiddleware.Options{
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			httpapi.SendErrorJSON(w, statusCode, errors.New(message))
		},
	}))
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(s.cfg.RequestTimeout))
	s.router.Use(httplog.RequestLogger(logger))

	HandlerFromMux(s, s.router)

	logger.Info().Msgf("server starting on port %d", s.cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), s.router); err != nil {
		return err
	}

	return nil
}

func (s *Server) GetImage(w http.ResponseWriter, r *http.Request, imageID string) {
	img, err := s.svc.Get(r.Context(), imageID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusNotFound, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}
	httpapi.SendJSON(w, http.StatusOK, img)
}

func (s *Server) CreateImage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, s.cfg.MaxSize)
	contentType := r.Header.Get(httpapi.ContentTypeHeader)

	if contentType == httpapi.ApplicationJson {
		s.createImageFromJSON(w, r)
		return
	} else if contentType == httpapi.ImagePNG || contentType == httpapi.ImageJPG || contentType == httpapi.ImageJPEG {
		s.createImageFromBinary(w, r)
		return
	} else if strings.HasPrefix(contentType, httpapi.MultipartFormData) {
		s.createImageFromForm(w, r)
		return
	}

	httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, ErrInvalidContentType)
}

func (s *Server) createImageFromJSON(w http.ResponseWriter, r *http.Request) {
	var create CreateImageJSON
	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
		return
	}

	img, err := s.svc.CreateFromRemote(r.Context(), create.URL)
	if err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}

	httpapi.SendJSON(w, http.StatusCreated, img)
}

func (s *Server) createImageFromBinary(w http.ResponseWriter, r *http.Request) {
	rawImg, rawImgType, err := image.Decode(r.Body)

	if rawImgType != PNG && rawImgType != JPG && rawImgType != JPEG {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, ErrInvalidImgFormat)
		return
	}

	if err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
		return
	}

	img, err := s.svc.CreateFromUpload(r.Context(), rawImg)
	if err != nil {
		if errors.Is(err, httpapi.ErrBadRequest) || errors.Is(err, ErrImgSizeTooLarge) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}

	httpapi.SendJSON(w, http.StatusCreated, img)
}

func (s *Server) createImageFromForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(s.cfg.MaxSize); err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
		return
	}
	rawImg, _, err := image.Decode(file)
	if err != nil {
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}

	img, err := s.svc.CreateFromUpload(r.Context(), rawImg)
	if err != nil {
		if errors.Is(err, ErrImgSizeTooLarge) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusBadRequest, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}

	httpapi.SendJSON(w, http.StatusCreated, img)
}

func (s *Server) DeleteImage(w http.ResponseWriter, r *http.Request, imageID string) {
	if err := s.svc.Delete(r.Context(), imageID); err != nil {
		if errors.Is(err, ErrNotFound) {
			httpapi.SendErrorJSONWithRequest(w, r, http.StatusNotFound, err)
			return
		}
		httpapi.SendErrorJSONWithRequest(w, r, http.StatusInternalServerError, err)
		return
	}
	httpapi.SendJSON(w, http.StatusNoContent, nil)
}
