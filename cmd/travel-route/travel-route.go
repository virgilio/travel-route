package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/virgilio/travel-route/server"
	"github.com/virgilio/travel-route/shortestpath"
	"github.com/virgilio/travel-route/storage"
)

func main() {
	_, irfErr := os.Stat(os.Args[1])
	if irfErr != nil {
		log.Fatal(fmt.Sprintf("Input file %s with routes does not exists", os.Args[1]))
	}
	routesFile := os.Args[1]

	addRouteHandler := http.HandlerFunc(server.AddRouteHandler)
	bestRouteHandler := http.HandlerFunc(server.BestRouteHandler)

	http.Handle("/addRoute", server.ContextMiddleware(server.RFKEY, routesFile, addRouteHandler))
	http.Handle("/bestRoute", server.ContextMiddleware(server.RFKEY, routesFile, bestRouteHandler))
	go http.ListenAndServe(":8080", nil)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the route (FROM-TO):")
	for scanner.Scan() {
		route := strings.Split(scanner.Text(), "-")
		if len(route) != 2 {
			fmt.Println("Invalid Route, ")
		} else {
			flights, loadErr := storage.LoadFlights(routesFile)
			if loadErr != nil {
				fmt.Println(server.AppError{server.LoadError, loadErr.Error(), time.Now()}.Error())
				break
			}
			flight := shortestpath.Flight{route[0], route[1], 0}
			graph := flights.CreateRouteGraph()

			if _, ok := graph[flight.From]; !ok {
				fmt.Println(server.AppError{server.MissingCityError, flight.From, time.Now()}.Error())
				break
			}
			if _, ok := graph[flight.To]; !ok {
				fmt.Println(server.AppError{server.MissingCityError, flight.To, time.Now()}.Error())
				break
			}

			cost, cities := graph[flight.From].ShortestPath(graph[flight.To], graph)
			citiesNames := []string{}
			for _, city := range cities {
				citiesNames = append(citiesNames, city.Name)
			}
			tickCost := "$" + strconv.Itoa(cost)
			fmt.Println("best route:", strings.Join(citiesNames, " - "), ">", tickCost)
		}
		fmt.Println("please enter the route (FROM-TO):")
	}
}
