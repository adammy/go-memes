package server

import (
	"net/http"
	"strconv"

	memePkg "github.com/adammy/memepen-services/pkg/meme"
	"github.com/gin-gonic/gin"
)

var _ Server = (*ginServer)(nil)

type ginServer struct {
	config  *Config
	router  *gin.Engine
	service *memePkg.Service
}

// NewGinServer returns a new Server utilizing the gin framework.
func NewGinServer(cfg *Config, svc *memePkg.Service) (*ginServer, error) {
	r := gin.New()

	return &ginServer{
		config:  cfg,
		router:  r,
		service: svc,
	}, nil
}

func (s *ginServer) Start() error {
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	err := s.router.SetTrustedProxies([]string{})
	if err != nil {
		return err
	}

	if err := s.registerRoutes(); err != nil {
		return err
	}

	if err := s.router.Run(":" + strconv.Itoa(int(s.config.Port))); err != nil {
		return err
	}
	return nil
}

func (s *ginServer) registerRoutes() error {
	if s.config.ServeLocalAssets {
		s.router.Static("/assets", "./assets")
	}

	v1 := s.router.Group("/v1")
	{
		v1.GET("/ping", s.pingHandler)
		v1.POST("/templates/:templateID/memes", s.createMemeFromTemplateIDHandler)
	}

	return nil
}

func (s *ginServer) pingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (s *ginServer) createMemeFromTemplateIDHandler(ctx *gin.Context) {
	templateID := ctx.Param("templateID")
	if templateID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "templateID required",
		})
		return
	}

	var createMeme memePkg.CreateMemeFromTemplate
	if err := ctx.ShouldBind(&createMeme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	meme, err := s.service.CreateMemeAndUploadFromTemplateID(templateID, createMeme.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, meme)
}
