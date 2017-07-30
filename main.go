package main

import "fmt"

func main() {
	grid := Grid{0, 0}
	roverOne := Rover{Coords{0, 0}, "", 0}
	roverTwo := Rover{Coords{0, 0}, "", 0}

	grid.Build()
	grid.AddRover(&roverOne)
	grid.AddRover(&roverTwo)

	fmt.Printf("Rover One Position Report: %v %v %v\n", roverOne.Position.X,
		roverOne.Position.Y, roverOne.Direction)
	fmt.Printf("Rover Two Position Report: %v %v %v\n", roverTwo.Position.X,
		roverTwo.Position.Y, roverTwo.Direction)
}
