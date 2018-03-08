package main

import (
	"fmt"
	"sort"
)

// Original Code:
// https://golang.org/pkg/sort/
// Example (Sortkeys)

// Raj's Custom Modification
// Shared on Playground
// https://play.golang.org/p/TN-Otr5VQ0

// A couple of type definitions to make the units clear.
type earthMass float64
type au float64

// A Planet defines the properties of a solar system object.
type Planet struct {
	name     string
	mass     earthMass
	distance au
}

type PlanetSort struct {
	planets []Planet
	sortby  string
}

var planetslice = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

// How do we know what to sorty by
func (s *PlanetSort) SortPlanetsBy(sortstring string) {
	s.sortby = sortstring
}

// Master Sort Calling Function
func (s *PlanetSort) Sort() {
	sort.Sort(s)
}

// Implement the Interface Contracts: Less (i,j), Len(), Swap (i,j)

// Implementing the Len() Interface Function
func (s *PlanetSort) Len() int {
	return len(s.planets)
}

// Implementing the Len() Interface Function
func (s *PlanetSort) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Implementing the Len() Interface Function
func (s *PlanetSort) Less(i, j int) bool {
	switch s.sortby {
	case "name":
		return s.planets[i].name < s.planets[j].name
	case "mass":
		return s.planets[i].mass < s.planets[j].mass
	case "distance":
		return s.planets[i].distance < s.planets[j].distance

	}
	return false
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func main() {
	// Sort planets by name
	var planetsortinstance PlanetSort

	planetsortinstance = PlanetSort{planetslice, "name"}
	planetsortinstance.Sort()
	fmt.Println("By name:", planetsortinstance.planets)

	planetsortinstance = PlanetSort{planetslice, "mass"}
	planetsortinstance.Sort()
	fmt.Println("By name:", planetsortinstance.planets)

	planetsortinstance = PlanetSort{planetslice, "distance"}
	planetsortinstance.Sort()
	fmt.Println("By name:", planetsortinstance.planets)

}
