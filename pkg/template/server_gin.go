package template

import (
	http2 "github.com/adammy/memepen-services/pkg/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var _ http2.Server = (*ginServer)(nil)

type ginServer struct {
	config *Config
	router *gin.Engine
}

// NewGinServer returns a new Server utilizing the gin framework.
func NewGinServer(cfg *Config) (*ginServer, error) {
	r := gin.New()

	return &ginServer{
		config: cfg,
		router: r,
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
	v1 := s.router.Group("/v1")
	{
		v1.GET("/ping", s.pingHandler)
	}

	return nil
}

func (s *ginServer) pingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
