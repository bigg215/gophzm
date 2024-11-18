package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllZips(context.Background())

	if err != nil {
		return fmt.Errorf("error reseting database: %w", err)
	}

	return nil
}

func handlerLookupZip(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <zip>", cmd.Name)
	}
	zipInteger, err := strconv.Atoi(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("invalid zip: %w", err)
	}

	zoneRecord, err := s.db.GetZipZone(context.Background(), int32(zipInteger))

	if err != nil {
		return fmt.Errorf("error retrieving record: %w", err)
	}

	fmt.Printf("Zipcode:\t%d\n", zoneRecord.Zipcode)
	fmt.Printf("Zone:\t\t%s\n", zoneRecord.Zone)
	fmt.Printf("Temp Range:\t%s\n", zoneRecord.Temprange)
	fmt.Printf("Dataset:\t%d\n", zoneRecord.Year)

	return nil
}
