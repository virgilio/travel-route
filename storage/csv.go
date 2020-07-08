// Package storage to deal with data
package storage

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/virgilio/travel-route/shortestpath"
)

// AppendRoute function appends a route line to csv file
func AppendRoute(routesFileName string, flight shortestpath.Flight) error {
	routesFile, fileErr := os.OpenFile(routesFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if fileErr != nil {
		return fileErr
	}
	defer routesFile.Close()

	writer := csv.NewWriter(routesFile)
	writeErr := writer.WriteAll([][]string{flight.CSVRecord()})
	if writeErr != nil {
		return writeErr
	}

	return routesFile.Close()
}

// LoadFlights function loads all Flights from file
func LoadFlights(routesFileName string) (*shortestpath.Flights, error) {
	flights := shortestpath.Flights(make([]*shortestpath.Flight, 0))

	routesFile, fileErr := os.Open(routesFileName)
	if fileErr != nil {
		return nil, fileErr
	}

	records, loadErr := csv.NewReader(routesFile).ReadAll()
	if loadErr != nil {
		return nil, loadErr
	}

	for _, record := range records {
		cost, costErr := strconv.Atoi(record[2])
		if costErr != nil {
			return nil, costErr
		}
		flight := shortestpath.Flight{record[0], record[1], cost}
		flights = append(flights, &flight)
	}

	return &flights, routesFile.Close()
}
