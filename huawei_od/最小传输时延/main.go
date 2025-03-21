package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
*
A --1-- B --2-- C

	\      |      /
	 3     1     4
	  \    |    /
	    -- D --

a
b d
d c d
d c

3 3
1 2 11
2 3 13
1 3 50
1 3

6 9
1 2 2
1 3 5
2 3 1
2 4 4
3 4 2
3 5 7
4 5 3
4 6 8
5 6 2
1 6

5 7
1 2 1
1 3 3
2 3 1
2 4 5
3 4 2
3 5 7
4 5 1
1 5

漫步香港1999
*/
func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	g := NewGraph(n)

	for i := 0; i < m; i++ {
		u, v, dis := 0, 0, 0
		fmt.Scan(&u, &v, &dis)
		g.AddEdge(u, v, dis)
	}

	start, end := 0, 0
	fmt.Scan(&start, &end)

	dist := Djikstra(g, start, n)

	fmt.Println(dist, dist[end])
}

func Djikstra(g *Graph, start, n int) []int {

	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	heap.Push(&priorityQueue, &Node{
		Vertex:   start,
		Distance: 0,
	})

	dist := make([]int, 0)
	for i := 0; i < n+1; i++ {
		dist = append(dist, math.MaxInt32)
	}
	dist[start] = 0

	for len(priorityQueue) > 0 {
		node := heap.Pop(&priorityQueue).(*Node)
		dis := node.Distance
		vertex := node.Vertex

		if dis > dist[node.Vertex] {
			continue
		}

		for _, neighbor := range g.adj[vertex] {
			if dis+neighbor.Distance < dist[neighbor.Vertex] {
				dist[neighbor.Vertex] = dis + neighbor.Distance
				heap.Push(&priorityQueue, &Node{
					Vertex:   neighbor.Vertex,
					Distance: dist[neighbor.Vertex],
				})
			}
		}
	}

	return dist
}

type Graph struct {
	adj [][]Node
	n   int
}

func NewGraph(n int) *Graph {
	adj := make([][]Node, n+1)
	for k, _ := range adj {
		adj[k] = make([]Node, 0)
	}
	return &Graph{
		adj: adj,
		n:   n,
	}
}

func (g Graph) AddEdge(u, v, dis int) {
	g.adj[u] = append(g.adj[u], Node{
		Vertex:   v,
		Distance: dis,
	})
}

type Node struct {
	Vertex   int
	Distance int
	Index    int
}
type PriorityQueue []*Node

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].Distance < p[j].Distance
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].Index, p[j].Index = p[j].Index, p[i].Index
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*p = old[0 : n-1]
	return item
}

func (p *PriorityQueue) Push(n any) {
	item := n.(*Node)
	item.Index = len(*p)
	*p = append(*p, item)
}
