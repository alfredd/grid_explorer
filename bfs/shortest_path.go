package main

import (
	"fmt"

	"github.com/alfredd/grid_explorer/grid"
)

type Node struct {
	current grid.Pair[int, int]
	parent  *Node
}

func shortest_path(g *grid.Grid) {
	explored := make(map[grid.Pair[int, int]]bool)
	queue := []Node{}
	start := Node{g.GetStart(), nil}
	queue = append(queue, start)
	explored[start.current] = true
	var path Node
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if g.IsEnd(v.current.First, v.current.Second) {
			path = v
			break
		}
		for _, n := range g.GetNeighbours(v.current.First, v.current.Second) {
			if _, ok := explored[n]; !ok {
				explored[n] = true
				queue = append(queue, Node{n, &v})

			}
		}
	}
	fmt.Println("Shortest path: ")
	for path.parent != nil {
		fmt.Print(path.current)
		path = *path.parent
	}
	fmt.Println(path.current)

}

func main() {
	g := grid.CreateGridWithObstacles(10, 10, 20)
	g.PrintGrid()
	shortest_path(g)
}
