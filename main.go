package	main

import (
	"fmt"
	"log"
	route2 "go/application/route"
)


func main() {
	route := route2.Route{
		ID: "1",
		ClientID: "1"
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[1])
}