package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X int
	Y int
}

type Grid struct {
	width  int
	height int
}

func (grid *Grid) Build() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Mars Rover remote control program.")
	fmt.Println("Please enter the grid size 'X Y': ")
	text, _ := reader.ReadString('\n')

	words := strings.Fields(text)
	grid.width, _ = strconv.Atoi(words[0])
	grid.height, _ = strconv.Atoi(words[1])

}

func (grid Grid) AddRover(rover *Rover) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter initial position: ")
	position, _ := reader.ReadString('\n')
	rover.Place(position)

	fmt.Println("Please enter movement sequence: ")
	sequence, _ := reader.ReadString('\n')
	rover.Move(sequence, &grid)
}
