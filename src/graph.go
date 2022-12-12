package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// GRAPH ( not oriented )

type Graph struct {
	nodes map[int]*Node
}

func newGraph() Graph {
	return Graph{nodes: make(map[int]*Node)}
}

func (g *Graph) addNode(id int) {
	g.nodes[id] = &Node{id: id}
}

func (g *Graph) getNode(id int) *Node {
	return g.nodes[id]
}

func (g *Graph) addLink(id1, id2 int) {
	g.nodes[id1].addLink(id2)
	g.nodes[id2].addLink(id1)
}

func (g *Graph) delLink(id1, id2 int) {
	g.nodes[id1].delLink(id2)
	g.nodes[id2].delLink(id1)
}

func (g *Graph) hasLink(id1, id2 int) bool {
	return g.nodes[id1].hasLink(id2)
}

// bfs
func (g *Graph) getDistances(start int) map[int]int {
	distances := map[int]int{start: 0}
	queue := []int{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range g.nodes[current].childs {
			if _, ok := distances[child]; !ok {
				distances[child] = distances[current] + 1
				queue = append(queue, child)
			}
		}
	}

	return distances
}

func (g *Graph) getShorterPath(start, end int) []int {

	path := []int{}
	distances := g.getDistances(start)

	if currentDist, ok := distances[end]; ok {
		path = append(path, end)
		current := end

		for currentDist > 0 {
			min := -1

			// todo: other method ??
			// min > s[i]? min = s[i] : min
			for _, child := range g.nodes[current].childs {
				if distances[child] < currentDist {
					min = child
					currentDist = distances[child]
				}
			}

			path = append([]int{min}, path...)
			current = min
		}
	}
	return path

}

// NODE

type Node struct {
	id     int
	childs []int
}

func (n *Node) addLink(child int) {
	n.childs = append(n.childs, child)
}

func (n *Node) delLink(child int) {
	index := slices.Index(n.childs, child)
	n.childs = slices.Delete(n.childs, index, index+1) // todo: native method
}

func (n *Node) hasLink(id int) bool {
	for _, c := range n.childs {
		if c == id {
			return true
		}
	}
	return false
}

func main() {

	g := newGraph()

	g.addNode(1)
	g.addNode(2)
	g.addNode(3)
	g.addNode(4)
	g.addNode(5)
	g.addNode(6)
	g.addNode(7)
	g.addNode(8)

	g.addLink(1, 2)
	g.addLink(1, 3)
	g.addLink(1, 4)
	g.addLink(2, 5)
	g.addLink(3, 4)
	g.addLink(3, 6)
	g.addLink(5, 6)
	g.addLink(6, 7)

	/*g.addNode(1)
	g.addNode(2)
	g.addNode(3)
	g.addNode(4)

	g.addLink(1, 2)
	g.addLink(1, 4)
	g.addLink(2, 3)*/

	fmt.Printf("Graph created %v\n", g)

	fmt.Printf("Has link %d -> %d ? %v\n", 1, 2, g.hasLink(1, 2))
	fmt.Printf("Has link %d -> %d ? %v\n", 1, 3, g.hasLink(1, 3))

	fmt.Printf("Distances from 1: %v\n", g.getDistances(1))

	fmt.Printf("Shorter path from 1 to 7: %v\n", g.getShorterPath(1, 7))

}
