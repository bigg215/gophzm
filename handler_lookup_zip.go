package main

import (
	"context"
	"fmt"
)

func handlerLookupZip(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <zip>", cmd.Name)
	}

	zoneRecord, err := s.db.GetZipZone(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("error retrieving record: %w", err)
	}

	fmt.Printf("Zipcode:\t%s\n", zoneRecord.Zipcode)
	fmt.Printf("Zone:\t\t%s\n", zoneRecord.Zone)
	fmt.Printf("Temp Range:\t%s\n", zoneRecord.Temprange)
	fmt.Printf("Dataset:\t%d\n", zoneRecord.Year)

	return nil
}
