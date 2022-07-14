package server

import (
	"github.com/adammy/memepen-services/pkg/meme"
	"github.com/adammy/memepen-services/pkg/meme/font"
	"github.com/adammy/memepen-services/pkg/meme/image"
	templateRepositoryPkg "github.com/adammy/memepen-services/pkg/template/repository"
	"github.com/adammy/memepen-services/pkg/uploader"
	"github.com/spf13/viper"
)

// Config defines configuration for the Server.
type Config struct {
	// Port defines the port the server runs on.
	Port uint `mapstructure:"port"`

	// ServeLocalAssets defines if the Server should serve the local assets.
	ServeLocalAssets bool `mapstructure:"serve_local_assets"`

	// FontRepositoryType defines the RepositoryType for fonts.
	FontRepositoryType font.RepositoryType `mapstructure:"font_repository_type"`

	// ImageRepositoryType defines the RepositoryType for template images.
	ImageRepositoryType image.RepositoryType `mapstructure:"image_repository_type"`

	// MemeRepositoryType defines the RepositoryType for memes.
	MemeRepositoryType meme.RepositoryType `mapstructure:"meme_repository_type"`

	// TemplateRepositoryType defines the RepositoryType for templates.
	TemplateRepositoryType templateRepositoryPkg.Type `mapstructure:"template_repository_type"`

	// UploaderType defines the UploaderType.
	UploaderType uploader.Type `mapstructure:"uploader_type"`
}

type configWrapper struct {
	Meme Config `mapstructure:"meme"`
}

func LoadConfig(path, env string) (*Config, error) {
	var (
		config configWrapper
	)

	viper.SetDefault("meme.port", 8080)
	viper.SetDefault("meme.serve_local_assets", false)
	viper.SetDefault("meme.font_repository_type", font.InMemory)
	viper.SetDefault("meme.image_repository_type", image.Local)
	viper.SetDefault("meme.meme_repository_type", meme.InMemory)
	viper.SetDefault("meme.template_repository_type", templateRepositoryPkg.InMemory)
	viper.SetDefault("meme.uploader_type", uploader.Local)

	viper.AddConfigPath(path)
	viper.SetConfigName(env)
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config.Meme, nil
}
