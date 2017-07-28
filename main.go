package main

func main() {
	grid := Grid{0, 0}
	roverOne := Rover{Coords{0, 0}, ""}
	roverTwo := Rover{Coords{0, 0}, ""}

	grid.Build()
	grid.AddRover(&roverOne)
	grid.AddRover(&roverTwo)
}
