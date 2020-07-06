package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/virgilio/travel-route/storage"
)

// Node represents a single city node
type Node struct {
	name string
}

// Route represents a route between two city nodes
type Route struct {
	from  *Node
	to    *Node
	price int
}

func (r *Route) csvLine() string {
	return fmt.Sprintf("%s,%s,%d\n", r.from.name, r.to.name, r.price)
}

func readInputRouteFile() *Node {
	return nil
}

func addRouteToFile(irf string, route *Route) {

}

func bestRoute(irf, node *Node) {

}

func bestRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.RawQuery)
	fmt.Fprintf(w, "Find best route from s to s")
	fmt.Fprintf(w, "PATH: %s", r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "%s", r.Method)
}

func addRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Method Not Allowed")
	} else {
		fmt.Fprintf(w, "Call add new Route")
	}
}

func main() {
	storage.AppendRoute()
	inputRoutesFile := os.Args[1]
	_, irfErr := os.Stat(inputRoutesFile)
	if irfErr != nil {
		log.Fatal(fmt.Sprintf("Input file %s with routes does not exists", inputRoutesFile))
	}

	http.HandleFunc("/addRoute", addRouteHandler)
	http.HandleFunc("/bestRoute", bestRouteHandler)
	go http.ListenAndServe(":8080", nil)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the route: ")
	for scanner.Scan() {
		route := scanner.Text()
		fmt.Printf("best route: %s\n", route)
		fmt.Println("please enter the route: ")
	}
}
