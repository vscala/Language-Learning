package main

import (
	"fmt"
)

// Grapher interface specifies functions that can be performed on Graph objects
type Grapher interface {
	AddNode(value int)
	AddEdge(nodeA int, nodeB int, weight float64)
	//FindPath(nodeA int, nodeB int)
	ListNodes()
	ListEdges()
}

/* 
	The graph is implemented as a mapping from int -> int -> float.
	Edges of the graph are directed and weighted.
	The nodes map holds all the weighted edges of the map.
	
	- The first int represents the source node
	- The second int represents the destination node
	- The float64 represents the weight of the edge from source to destination
*/
type Graph struct {
	nodes map[int]map[int]float64
}

// AddNode is a function on Graph objects that adds a node given a value
func (graph Graph) AddNode(value int) {
	graph.nodes[value] = make(map[int]float64, 0.0)
}

// AddEdge is a function on Graph objects that adds an edge from two nodes 
// with a given weight 
func (graph Graph) AddEdge(nodeA int, nodeB int, weight float64) {
	graph.nodes[nodeA][nodeB] = weight
}

// ListNodes prints a list of all nodes 
func (graph Graph) ListNodes() {
	fmt.Println(" -- Nodes -- ")
	for node, _ := range graph.nodes {
		fmt.Println(node)
	}
	fmt.Println()
}

// ListEdges prints a list of all edges and their weights
func (graph Graph) ListEdges() {
	fmt.Println(" -- Edge List -- ")
	fmt.Println("Node 1 \t Node 2  Weight")
	for node, _ := range graph.nodes {
		for adj, _ := range graph.nodes[node] {
			fmt.Println(node, "\t", adj, "\t", graph.nodes[node][adj])
		}
	}
	fmt.Println()
}

func main() {
	graph := Graph{make(map[int]map[int]float64, 0)}
	graph.AddNode(5)
	graph.AddNode(7)
	graph.AddNode(512)
	graph.AddNode(51)
	graph.AddEdge(5,7,2)
	graph.AddEdge(512,51,5.5)
	graph.ListNodes()
	graph.ListEdges()
}
