package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/bigg215/gophzm/internal/database"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type state struct {
	db *database.Queries
}

const inputDirectory = "phzm-zips"

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	programState := state{
		db: dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//register commands
	cmds.register("status", handlerStatus)
	cmds.register("load", handlerLoad)

	//parse cli args
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args ...]")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//run the command
	err = cmds.run(&programState, command{
		Name: cmdName,
		Args: cmdArgs,
	})

	if err != nil {
		log.Fatal(err)
	}
}
