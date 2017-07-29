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

func (r *Rover) Move(s string, grid *Grid) {
	fmt.Println("==========================")
	fmt.Printf("Rover recieved sequence: %s", s)
	s = strings.ToUpper(s)
	seq := []byte(s)

	for _, char := range seq {
		switch string(char) {
		case "L":
			if r.Position.X < 1 {
				break
			}
			r.Position.X -= 1
		case "R":
			if (r.Position.X + 1) > grid.width {
				break
			}
			r.Position.X += 1
		case "M":
			switch r.Direction {
			case "N":
				r.Position.Y += 1
			case "E":
				r.Position.X += 1
			case "S":
				r.Position.Y -= 1
			case "W":
				r.Position.X -= 1
			default:
				break
			}
		default:
			break
		}
	}
}
