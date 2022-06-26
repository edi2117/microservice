package route
import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID string  `json:"routeid"`
	ClientID string  `json:"clientid"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat  float64  `json:"lat"`
	Long float64  `json:"long"`
}

func(r *Route)LoadPositions() error{
	if r.ID == "" {
		return errors.new{text: "route id not informed"}
	}

	f, err := os.Open(name: "destinations/" + r.ID + ".text")
	if err != nil {
		return err
	}

	defer f.close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := string.Split(scanner.Text(), sep:",")
		lat, err := strconv.ParseFloat(data[0], biSize: 64)

		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1],biSize: 64)

		if err != nil {
			return nil
		}

		r.Position = append(r.Position, Position{
			Lat: 	lat, 
			Long: 	long
		})

		return nil
	}
}
	
	type PartialRoutePosition struct {
		ID string  `json:"routeid"`
		ClientID string `json:"clientid"`
		Position []float64 `json:"position"`
		Finished bool `json:"finished"`
	}

func(e *Route) ExportJsonPositions() ([][]string, error) {
		var route PartialRoutePosition
		var result []string
		total := len(r.Positions)

		for k, v := range r.Positions {
			route.id = r.ID
			route.ClientID = r.ClientID
			route.Position = []float64{ v.Lat, v.Long }
			route.Finished = false
			if total-1 == k {
				route.Finished = true
			}

			jsonRoute := json.Marshal(route)
			if err != nil {
				return nil, err
			}
			result = append(result, string(jsonRoute))
		}
	return result, nil
}