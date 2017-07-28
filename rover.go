package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rover struct {
	Position  Coords
	Direction string
}

func (rover *Rover) Place(i string) {
	words := strings.Fields(i)
	rover.Position.X, _ = strconv.Atoi(words[0])
	rover.Position.Y, _ = strconv.Atoi(words[1])
	rover.Direction = words[2]
}

func (r *Rover) Move(seq string, grid *Grid) {
	fmt.Printf("Rover recieved sequence: %s", seq)
}
