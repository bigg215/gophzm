package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const phzmCsvHeaders = "zipcode,zone,trange,zonetitle\n"

func handlerCombine(s *state, cmd command) error {
	files, err := os.ReadDir(inputDirectory)

	if err != nil {
		return fmt.Errorf("invalid directory: %w", err)
	}

	csvCombined := phzmCsvHeaders

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

		scanner := bufio.NewScanner(csvFile)
		scanner.Scan()

		for scanner.Scan() {
			csvCombined += fmt.Sprintln(scanner.Text())
		}
	}

	outputFilePath := filepath.Join(outputDirectory, outputFileName)

	outputFile, err := os.Create(outputFilePath)

	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outputFile.Close()

	bytesWritten, err := outputFile.Write([]byte(csvCombined))

	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}
	fmt.Printf("Combine completed:\n")
	fmt.Printf("Output:\t%s\n", outputFileName)
	fmt.Printf("Bytes:\t%d\n", bytesWritten)
	return nil
}
