package meme

import (
	"github.com/adammy/memepen-services/pkg/font"
	"github.com/adammy/memepen-services/pkg/image"
	templateRepositoryPkg "github.com/adammy/memepen-services/pkg/template"
	"github.com/spf13/viper"
)

// Config defines configuration for the Server.
type Config struct {
	// Port defines the port the server runs on.
	Port uint `mapstructure:"port"`

	// ServeLocalAssets defines if the Server should serve the local assets.
	ServeLocalAssets bool `mapstructure:"serve_local_assets"`

	// FontRepositoryType defines the Type for fonts.
	FontRepositoryType font.GetterType `mapstructure:"font_repository_type"`

	// ImageRepositoryType defines the Type for template images.
	ImageRepositoryType image.GetterType `mapstructure:"image_repository_type"`

	// MemeRepositoryType defines the Type for memes.
	MemeRepositoryType RepositoryType `mapstructure:"meme_repository_type"`

	// TemplateRepositoryType defines the Type for templates.
	TemplateRepositoryType templateRepositoryPkg.RepositoryType `mapstructure:"template_repository_type"`

	// UploaderType defines the UploaderType.
	UploaderType image.UploaderType `mapstructure:"uploader_type"`
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
	viper.SetDefault("meme.font_repository_type", font.LocalGetter)
	viper.SetDefault("meme.image_repository_type", image.LocalGetter)
	viper.SetDefault("meme.meme_repository_type", InMemoryRepository)
	viper.SetDefault("meme.template_repository_type", templateRepositoryPkg.InMemoryRepository)
	viper.SetDefault("meme.uploader_type", image.LocalUploader)

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
