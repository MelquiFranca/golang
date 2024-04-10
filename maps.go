package main

import "fmt"

type PlayerFootball struct {
	name string
	year int
}
type VertexT struct {
	Lat, Long float64
}

var positions map[string][]PlayerFootball

func main() {
	positions = make(map[string][]PlayerFootball)
	positions["Goleiros"] = []PlayerFootball{
		PlayerFootball{name: "Rossi", year: 29},
		PlayerFootball{name: "Matheus Cunha", year: 29},
	}
	positions["Atacantes"] = []PlayerFootball{
		PlayerFootball{name: "Gabriel Barbosa", year: 26},
		PlayerFootball{name: "Pedro", year: 24},
	}
	fmt.Println(positions)
	for key, values := range positions {
		// fmt.Println(key, positions[key])
		//or
		fmt.Println(key, values)
	}

	googleKey := "Google"
	beelsKey := "Bell Labs"
	// Sem usar make
	var m = map[string]VertexT{
		"Bell Labs": {40.68433, -74.39967}, //Posso omitir o tipo, se for apenas o nome do tipo para o nível superior
		googleKey:   VertexT{37.42202, -122.08408},
	}
	fmt.Println(m)
	delete(m, googleKey)
	fmt.Println(m)
	elem, ok := m[googleKey]
	fmt.Println(googleKey, ":", elem, ok)
	elem, ok = m[beelsKey]
	fmt.Println(beelsKey, ":", elem, ok)
	m[googleKey] = VertexT{37.42202, -122.08408}
	fmt.Println(m)
}
