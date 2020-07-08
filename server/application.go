package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/virgilio/travel-route/shortestpath"
	"github.com/virgilio/travel-route/storage"
)

// RFKEY Context Key
const RFKEY ContextKey = "routesFile"

// LoadError error loading flight data
const LoadError = "Error loading flight stored data"

// MissingCityError error when city not present on graph
const MissingCityError = "Requested city not found"

const methodError = "Method not allowed"
const parsingError = "Error parsing flight data"
const appendError = "Error adding new Flight"

// TravelPlan ad
type TravelPlan struct {
	Cost   int
	Cities []string
}

// AppError standard error structure
type AppError struct {
	Message string    `json:"error"`
	Cause   string    `json:"cause"`
	Time    time.Time `json:"time"`
}

func (e AppError) Error() string {
	response, _ := json.Marshal(&e)
	return string(response)
}

// BestRouteHandler function handler
func BestRouteHandler(w http.ResponseWriter, r *http.Request) {
	rfName, _ := r.Context().Value(RFKEY).(string)
	flights, loadErr := storage.LoadFlights(rfName)
	if loadErr != nil {
		fmt.Fprintf(w, AppError{LoadError, loadErr.Error(), time.Now()}.Error())
		return
	}
	var flight shortestpath.Flight
	decErr := json.NewDecoder(r.Body).Decode(&flight)
	if decErr != nil {
		fmt.Fprintf(w, AppError{parsingError, decErr.Error(), time.Now()}.Error())
		return
	}

	graph := flights.CreateRouteGraph()

	if _, ok := graph[flight.From]; !ok {
		fmt.Fprintf(w, AppError{MissingCityError, flight.From, time.Now()}.Error())
		return
	}
	if _, ok := graph[flight.To]; !ok {
		fmt.Fprintf(w, AppError{MissingCityError, flight.To, time.Now()}.Error())
		return
	}

	cost, cities := graph[flight.From].ShortestPath(graph[flight.To], graph)

	plan := &TravelPlan{cost, []string{}}
	for _, city := range cities {
		plan.Cities = append(plan.Cities, city.Name)
	}

	response, _ := json.Marshal(plan)
	fmt.Fprintf(w, string(response))
}

// AddRouteHandler function handler
func AddRouteHandler(w http.ResponseWriter, r *http.Request) {
	rfName, _ := r.Context().Value(RFKEY).(string)
	if r.Method != "POST" {
		fmt.Fprintf(w, AppError{methodError, r.Method, time.Now()}.Error())
	} else {
		var flight shortestpath.Flight
		decErr := json.NewDecoder(r.Body).Decode(&flight)
		if decErr != nil {
			fmt.Fprintf(w, AppError{parsingError, decErr.Error(), time.Now()}.Error())
			return
		}
		if appErr := storage.AppendRoute(rfName, flight); appErr != nil {
			fmt.Fprintf(w, AppError{appendError, appErr.Error(), time.Now()}.Error())
		}

		response, _ := json.Marshal(&flight)
		fmt.Fprintf(w, string(response))
	}
}
