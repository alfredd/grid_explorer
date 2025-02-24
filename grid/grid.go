package grid

import (
	"fmt"
	"math/rand"
)

type Grid struct {
	Width  int
	Height int
	state  [][]float64

	Start Pair[int, int]
	End   Pair[int, int]
}

type Pair[A, B int] struct {
	First  A
	Second B
}

func newGrid(width, height int) *Grid {

	g := &Grid{
		width, height,
		make([][]float64, height),
		Pair[int, int]{0, 0},
		Pair[int, int]{width - 1, height - 1},
	}
	for i := 0; i < height; i++ {
		g.state[i] = make([]float64, width)
	}
	return g
}

func (g *Grid) GetNeighbours(x, y int) []Pair[int, int] {
	cx, cy := x-1, y-1
	neighbours := []Pair[int, int]{}
	for i := cx; i < cx+3; i++ {
		for j := cy; j < cy+3; j++ {
			if i < 0 || j < 0 || i >= g.Width || j >= g.Height {
				continue
			}
			if i == x && j == y {
				continue
			}
			if g.state[j][i] >= 0 {
				neighbours = append(neighbours, Pair[int, int]{i, j})
			}

		}
	}
	return neighbours
}

func (g *Grid) GetStart() Pair[int, int] {
	return g.Start
}

func (g *Grid) PrintGrid() {
	fmt.Println("Grid: start: ", g.Start, " end: ", g.End)
	for i := range g.Height {
		for j := range g.Width {
			fmt.Printf("({%d, %d} %f ) ", i, j, g.state[i][j])
		}
		fmt.Println()
	}
}

func (g *Grid) setStart(x, y int) {
	g.Start = Pair[int, int]{x, y}
}

func (g *Grid) setEnd(x, y int) {
	g.End = Pair[int, int]{x, y}
}

func (g *Grid) setObstacle(x, y int) {
	g.state[x][y] = -1
}

func (g *Grid) IsEnd(x, y int) bool {
	return x == g.End.First && y == g.End.Second
}

func CreateGridWithObstacles(width, height, obstacleCount int) *Grid {
	g := newGrid(width, height)
	for range obstacleCount {
		x := rand.Intn(width)
		y := rand.Intn(height)
		g.setObstacle(x, y)
	}

	g.setStart(0, 0)
	if g.state[0][0] == -1 {
		g.state[0][0] = 0
	}
	g.setEnd(width-1, height-1)
	if g.state[height-1][width-1] == -1 {
		g.state[height-1][width-1] = 0
	}
	return g
}
