package main

import "fmt"

func handlerStatus(s *state, cmd command) error {
	fmt.Println("Status OK")
	return nil
}
