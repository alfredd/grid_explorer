package main

import (
	"fmt"
	"github.com/alfredd/grid_explorer/grid"
)

func dfs(g *grid.Grid, x, y int, discovered map[grid.Pair[int, int]]bool,path *[]grid.Pair[int, int]) bool{
	p := grid.Pair[int, int]{First:x, Second:y}
	discovered[p] = true
	if g.IsEnd(x, y) {
		fmt.Println("Found path to destination.")
		return true
	}
	for _, n := range g.GetNeighbours(x,y) {
		if _, ok := discovered[n]; !ok {
			found:=dfs(g, n.First, n.Second, discovered, path)
			if found { 
				*path = append(*path, n)
				return true
			}
		}
	}
	return false
}

func find_path(g *grid.Grid) {
	discovered := make (map[grid.Pair[int, int]]bool)
	path:= &[]grid.Pair[int, int]{}
	found:=dfs(g, g.GetStart().First, g.GetStart().Second, discovered, path)
	if found {
		fmt.Println("Path: ", *path)
	} else {
		fmt.Println("No path found.")
	}
	
}

func main() {
	g := grid.CreateGridWithObstacles(10,10,20)
	g.PrintGrid()
	find_path(g)
}
