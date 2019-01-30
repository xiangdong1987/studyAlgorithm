package shortestPath

import "fmt"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func TestShortestPath() {
	g := Graph{}
	g.Vertex = 6
	g.Adj = make(map[int][]Edge, 6)
	g.addEdge(0, 1, 10)
	g.addEdge(0, 4, 15)
	g.addEdge(1, 2, 15)
	g.addEdge(1, 3, 2)
	g.addEdge(3, 2, 1)
	g.addEdge(3, 5, 12)
	g.addEdge(2, 5, 5)
	g.addEdge(4, 5, 10)
	g.Dijkstra(0, 5)
}

type Edge struct {
	Sid    int
	Tid    int
	Weight int
}

type Graph struct {
	Adj    map[int][]Edge
	Vertex int
}

func (g *Graph) addEdge(s int, t int, w int) {
	edge := Edge{s, t, w}
	g.Adj[s] = append(g.Adj[s], edge)
}

type Vertex struct {
	Id   int
	Dist int
}

type PriorityQueue struct {
	nodes []Vertex
	count int
}

//插入堆
func (h *PriorityQueue) InsertHeap(one Vertex) {
	h.nodes = append(h.nodes, one)
	length := len(h.nodes)
	h.count = length - 1
	h.AdjustHeap(h.count)
}

//自下而上调整 (小顶堆)
func (h *PriorityQueue) AdjustHeap(length int) {
	if length < 1 {
		return
	}
	if length == 2 {
		if h.nodes[length].Dist < h.nodes[length-1].Dist {
			h.SliceNodeSwap(length, length-1)
		}
		return
	}
	i := length
	for i/2 > 0 && h.nodes[i].Dist < h.nodes[i/2].Dist {
		h.SliceNodeSwap(i, i/2)
		i = i / 2
	}
	return
}

//node slice交换
func (h *PriorityQueue) SliceNodeSwap(i int, j int) {
	x := h.nodes[i]
	h.nodes[i] = h.nodes[j]
	h.nodes[j] = x
}

//自上向下堆化
func (h *PriorityQueue) Heapfiy(length int, pos int) {
	for {
		maxPos := pos
		if pos*2 < length && h.nodes[pos].Dist < h.nodes[pos*2].Dist {
			maxPos = pos * 2
		}
		if pos*2+1 < length && h.nodes[maxPos].Dist < h.nodes[pos*2+1].Dist {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.SliceNodeSwap(pos, maxPos)
		pos = maxPos
	}
}

//获取堆顶
func (h *PriorityQueue) GetTopHeap() Vertex {
	if h.count == 0 {
		panic("Heap is empty")
	}
	top := h.nodes[1]
	//堆顶和堆底交换
	h.SliceNodeSwap(1, len(h.nodes)-1)
	length := len(h.nodes) - 2
	h.Heapfiy(length, 1)
	h.nodes = append(h.nodes[:length+1], h.nodes[length+2:]...)
	h.count--
	return top
}

func (g *Graph) Dijkstra(s int, t int) {
	predecessor := make([]int, g.Vertex)
	vertexes := make([]Vertex, g.Vertex)
	for i := 0; i < g.Vertex; i++ {
		vertexes[i] = Vertex{i, MaxInt}
	}
	queue := PriorityQueue{}
	queue.nodes = append(queue.nodes, Vertex{})
	inqueue := make([]bool, g.Vertex)
	vertexes[s].Dist = 0
	queue.InsertHeap(vertexes[s])
	inqueue[s] = true
	for len(queue.nodes) > 0 {
		fmt.Println(queue)
		minVertex := queue.GetTopHeap()
		fmt.Println(g.Adj[minVertex.Id])
		if minVertex.Id == t {
			break
		}
		for i := 0; i < len(g.Adj[minVertex.Id]); i++ {
			e := g.Adj[minVertex.Id][i]
			nextVertex := vertexes[e.Tid]
			if minVertex.Dist+e.Weight < nextVertex.Dist {
				nextVertex.Dist = minVertex.Dist + e.Weight
				predecessor[nextVertex.Id] = minVertex.Id
				if inqueue[nextVertex.Id] == true {
					queue.InsertHeap(nextVertex)
					queue.AdjustHeap(queue.count)
				} else {
					queue.InsertHeap(nextVertex)
					inqueue[nextVertex.Id] = true
				}
			}
		}
	}
	println("->", s)
	g.printPath(s, t, predecessor)
}

func (g *Graph) printPath(s int, t int, predecessor []int) {
	if s == t {
		return
	}
	g.printPath(s, predecessor[t], predecessor)
	println("->", t)
}
