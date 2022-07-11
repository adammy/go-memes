package server

import (
	"github.com/adammy/go-memes/pkg/meme"
	"github.com/google/uuid"
	"image/png"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var _ Server = (*ginServer)(nil)

type ginServer struct {
	engine *gin.Engine
}

// NewGinServer returns a new Server utilizing the gin framework.
func NewGinServer() (*ginServer, error) {
	r := gin.Default()

	err := r.SetTrustedProxies([]string{})
	if err != nil {
		return nil, err
	}

	if err := registerRoutes(r); err != nil {
		return nil, err
	}

	return &ginServer{
		engine: r,
	}, nil
}

func (s *ginServer) Start() error {
	if err := s.engine.Run(":8080"); err != nil {
		return err
	}
	return nil
}

func registerRoutes(r *gin.Engine) error {
	r.GET("/ping", pingHandler)
	r.POST("/memes", createMemeFromTemplateIDHandler)
	return nil
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func createMemeFromTemplateIDHandler(ctx *gin.Context) {
	var createMeme meme.CreateMemeFromTemplate
	if err := ctx.BindJSON(&createMeme); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	svc, err := meme.NewService("")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "this needs to go away lol",
		})
		return
	}

	img, err := svc.CreateMemeFromTemplateID(createMeme.TemplateID, createMeme.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	f, err := os.Create(uuid.NewString() + ".png")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "a+",
	})
}
