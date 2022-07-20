package main

import (
	"flag"

	"github.com/adammy/memepen-services/pkg/httpapi"
	"github.com/adammy/memepen-services/pkg/template"
)

func main() {
	env := flag.String("env", "local", "the environment for the application")
	flag.Parse()

	configWrapper, err := httpapi.LoadConfig[template.ConfigWrapper]("./configs", *env, template.DefaultConfig)
	if err != nil {
		panic(err)
	}
	config := configWrapper.Template

	s, err := template.NewServer(config.Type, &config)
	if err != nil {
		panic(err)
	}

	if err := s.Start(); err != nil {
		panic(err)
	}
}
