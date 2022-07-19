package template

import (
	"fmt"

	"github.com/adammy/memepen-services/pkg/image"
	"github.com/spf13/viper"
)

// Config defines configuration for the Server.
type Config struct {
	// Port defines the port the server runs on.
	Port uint `mapstructure:"port"`

	// ImageRepositoryType defines the GetterType for template images.
	ImageRepositoryType image.GetterType `mapstructure:"image_repository_type"`

	// TemplateRepositoryType defines the GetterType for templates.
	TemplateRepositoryType RepositoryType `mapstructure:"template_repository_type"`

	// UploaderType defines the UploaderType.
	UploaderType image.UploaderType `mapstructure:"uploader_type"`
}

type configWrapper struct {
	Template Config `mapstructure:"template"`
}

func LoadConfig(path, env string) (*Config, error) {
	var config configWrapper

	viper.SetDefault("template.port", 8080)
	viper.SetDefault("template.image_repository_type", image.LocalGetter)
	viper.SetDefault("template.template_repository_type", InMemoryRepository)
	viper.SetDefault("template.uploader_type", image.LocalUploader)

	viper.AddConfigPath(path)
	viper.SetConfigName(env)
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config.Template, nil
}
