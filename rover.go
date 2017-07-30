package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rover struct {
	Position       Coords
	Direction      string
	DirectionIndex int
}

var directions = map[string]int{
	"N": 1,
	"E": 2,
	"S": 3,
	"W": 4,
}

func (rover *Rover) Place(i string) {
	words := strings.Fields(i)
	rover.Position.X, _ = strconv.Atoi(words[0])
	rover.Position.Y, _ = strconv.Atoi(words[1])
	rover.Direction = words[2]
	rover.DirectionIndex = directions[rover.Direction]
}

func (r *Rover) Move(s string, grid *Grid) {
	fmt.Println("==========================")
	fmt.Printf("Rover recieved sequence: %s", s)
	s = strings.ToUpper(s)
	seq := []byte(s)

	for _, char := range seq {
		switch string(char) {
		case "L":
			r.rotate("left", grid)
		case "R":
			r.rotate("right", grid)
		case "M":
			r.forward(grid)
		default:
			break
		}
	}
}

func (r *Rover) forward(grid *Grid) {
	switch r.Direction {
	case "N":
		if r.Position.Y < grid.height {
			r.Position.Y += 1
		}
	case "E":
		if r.Position.X < grid.width {
			r.Position.X += 1
		}
	case "S":
		if r.Position.Y > 1 {
			r.Position.Y -= 1
		}
	case "W":
		if r.Position.X > 1 {
			r.Position.X -= 1
		}
	default:
		break
	}
}

func (r *Rover) rotate(dir string, grid *Grid) {
	switch dir {
	case "left":
		if r.DirectionIndex > 1 {
			r.DirectionIndex = directions[r.Direction] - 1
		}
	case "right":
		if r.DirectionIndex < 4 {
			r.DirectionIndex = directions[r.Direction] + 1
		}
	default:
		break
	}
}
