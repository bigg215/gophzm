package main

import (
	"fmt"
	"log"
	"os"
)

type state struct {
	userName string
}

const inputDirectory = "phzm-zips"
const outputDirectory = "phzm-combined"
const outputFileName = "phzmData.csv"

func main() {
	programState := state{
		userName: "system",
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//register commands
	cmds.register("status", handlerStatus)
	cmds.register("combine", handlerCombine)

	//parse cli args
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args ...]")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//run the command
	err := cmds.run(&programState, command{
		Name: cmdName,
		Args: cmdArgs,
	})

	if err != nil {
		log.Fatal(err)
	}
}
