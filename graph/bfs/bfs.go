package bfs

import (
	"container/list"
	"fmt"
)

type Node struct {
	Id        int
	Neighbors map[int]*Node
}

func Graph() {
	//initializing 4 nodes
	node1 := &Node{Id: 1, Neighbors: make(map[int]*Node)}
	node2 := &Node{Id: 2, Neighbors: make(map[int]*Node)}
	node3 := &Node{Id: 3, Neighbors: make(map[int]*Node)}

	node2.Neighbors[node1.Id] = node1 //node2 is directly connect to node1
	node3.Neighbors[node2.Id] = node2 //node3 is directly connect to node2
	node1.Neighbors[node3.Id] = node3 //node1 is directly connect to node3

	node0 := &Node{Id: 0, Neighbors: make(map[int]*Node)} //initializing node 0 as root

	//initializing a graph
	g := make(map[int]*Node)
	g[node0.Id] = node0
	g[node0.Id].Neighbors[node1.Id] = node1
	g[node0.Id].Neighbors[node2.Id] = node2
	g[node0.Id].Neighbors[node3.Id] = node3

	nodes := Traversing(node0)
	for _, node := range nodes {
		fmt.Println(node.Id)
	}
}

func Traversing(n *Node) []*Node {

	visited := make(map[int]*Node) //initializing for maintaining nodes which will visit in below loop
	queue := list.New()            //it is linked list structure but we going to use as queue

	queue.PushBack(n)
	visited[n.Id] = n

	//traversing in queue value
	for queue.Len() > 0 {
		currNode := queue.Front()
		//iterate via all of the neighbors
		for id, node := range currNode.Value.(*Node).Neighbors {
			if _, ok := visited[id]; !ok {
				visited[id] = node
				queue.PushBack(node)
			}
		}
		queue.Remove(currNode)
	}

	nodes := make([]*Node, 0)
	for _, val := range visited {
		nodes = append(nodes, val)
	}
	return nodes
}
