package shortestpath

import (
	"strconv"
)

// City represents a single city node
type City struct {
	Name          string
	departures    Flights
	pathFromStart []*City
	pathCost      int
}

// Flight represents a route from -> to and its cost
type Flight struct {
	From string
	To   string
	Cost int
}

// Flights represents a slice of flights
type Flights []*Flight

// CSVRecord returns a csv string line
func (f *Flight) CSVRecord() []string {
	return []string{f.From, f.To, strconv.Itoa(f.Cost)}
}

// CreateRouteGraph create the graph used on algo
func (flights Flights) CreateRouteGraph() map[string]*City {
	graph := make(map[string]*City)
	for _, flight := range flights {
		if _, exists := graph[flight.From]; !exists {
			graph[flight.From] = &City{flight.From, make([]*Flight, 0), nil, -1}
		}
		if _, exists := graph[flight.To]; !exists {
			graph[flight.To] = &City{flight.To, make([]*Flight, 0), nil, -1}
		}

		graph[flight.From].departures = append(graph[flight.From].departures, flight)
	}

	return graph
}

// ShortestPath calculates the shortest path given the start and end nodes
func (start *City) ShortestPath(end *City, graph map[string]*City) (int, []*City) {
	start.pathCost = 0
	start.pathFromStart = []*City{start}

	flights := Flights(make([]*Flight, 0))
	for _, departure := range start.departures {
		flights = append(flights, departure)
	}
	for len(flights) > 0 {
		flight := flights[0]
		flights = flights[1:]
		departCity := graph[flight.From]
		city := graph[flight.To]

		if city.pathCost < 0 || flight.Cost+departCity.pathCost < city.pathCost {
			city.pathCost = flight.Cost + departCity.pathCost
			city.pathFromStart = append(departCity.pathFromStart, []*City{city}...)
			for _, departure := range city.departures {
				flights = append(flights, departure)
			}
		}
	}

	return end.pathCost, end.pathFromStart
}
