package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllZips(context.Background())

	if err != nil {
		return fmt.Errorf("error reseting database: %w", err)
	}

	return nil
}
