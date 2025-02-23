package main

import  "github.com/alfredd/grid_explorer/grid"

func main() {
	g := grid.CreateGridWithObstacles(10,10,20)
	g.PrintGrid()
}
