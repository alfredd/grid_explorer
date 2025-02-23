package grid

import "fmt"

type Grid struct {
	Width  int
	Height int
	state [][]float64

	Start Pair[int, int]
	End Pair[int, int]
}

type Pair[A, B int] struct {
	First A
	Second B
}

func NewGrid(width, height int) *Grid {

	g := &Grid{
		width, height,
		make([][]float64, height),
		Pair[int, int]{0, 0},
		Pair[int, int]{width-1, height-1},
	}
	for i:=0; i<height; i++ {
		g.state[i] = make([]float64, width)
	}
	return g
}

func (g *Grid) GetNeighbours(x,y int) []Pair[int, int] {
	if x <0 {
		x = 0
	}
	if x >= g.Width {
		x = g.Width - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= g.Height {
		y = g.Height - 1
	}
	neighbours := []Pair[int, int]{
		{x-1, y},
		{x+1, y},
		{x, y-1},
		{x, y+1},
	}
	return neighbours
}

func (g *Grid) PrintGrid() {
	for i:=0; i<g.Height; i++ {
		for j:=0; j<g.Width; j++ {
			fmt.Printf("%f ",g.state[i][j])
		}
		fmt.Println()
	}
}

func (g *Grid) SetStart(x, y int) {
	g.Start = Pair[int, int]{x, y}
}

func (g *Grid) SetEnd(x, y int) {
	g.End = Pair[int, int]{x, y}
}

func (g *Grid) Obstacles(x, y int) {
	g.state[y][x] = -1
}