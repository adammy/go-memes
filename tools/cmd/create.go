package main

import (
	"flag"
	"github.com/adammy/memepen-services/tools/create"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()

	cmdTypeStr := args[0]
	cmdType, ok := create.ToCommandType(cmdTypeStr)
	if !ok {
		log.Fatal("type should be valid")
	}

	name := args[1]

	templateFuncs := map[create.CommandType]create.CommandFunc{
		create.APICommand: func() error {
			return create.CreateAPI(name)
		},
		create.ServerCommand: func() error {
			return create.CreateNewServer(name)
		},
	}

	templateFuncs[cmdType]()
}
