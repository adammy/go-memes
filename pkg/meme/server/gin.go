package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ Server = (*ginServer)(nil)

type ginServer struct {
	engine *gin.Engine
}

// NewGinServer returns a new Server utilizing the gin framework.
func NewGinServer() (*ginServer, error) {
	engine := gin.Default()

	if err := initRoutes(engine); err != nil {
		return nil, err
	}

	return &ginServer{
		engine: engine,
	}, nil
}

func (s *ginServer) Start() error {
	if err := s.engine.Run(); err != nil {
		return err
	}
	return nil
}

func initRoutes(e *gin.Engine) error {
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	return nil
}
