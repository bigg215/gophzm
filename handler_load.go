package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bigg215/gophzm/internal/database"
)

func handlerLoad(s *state, cmd command) error {
	files, err := os.ReadDir(inputDirectory)

	if err != nil {
		return fmt.Errorf("invalid directory: %w", err)
	}

	for _, file := range files {
		//skip directory guard
		if file.IsDir() {
			continue
		}
		//only csv files guard
		if file.Name()[len(file.Name())-3:] != "csv" {
			continue
		}

		filePath := filepath.Join(inputDirectory, file.Name())

		csvFile, err := os.Open(filePath)

		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
		defer csvFile.Close()

		r := csv.NewReader(csvFile)

		_, err = r.Read()

		if err != nil {
			return fmt.Errorf("error reading csv file: %w", err)
		}

		fmt.Printf("processing file:\t%s\n", filePath)

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("error reading csv file: %w", err)
			}
			_, err = s.db.AddZip(context.Background(), database.AddZipParams{
				Createdat: time.Now().UTC(),
				Updatedat: time.Now().UTC(),
				Zipcode:   record[0],
				Zone:      record[1],
				Temprange: record[2],
				Zonetitle: record[3],
				Year:      defaultDataYear,
			})

			if err != nil {
				return fmt.Errorf("error adding record: %w", err)
			}

			fmt.Println(record)
		}

	}
	return nil
}
