package shortestPath

import "fmt"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Edge struct {
	Sid    int
	Tid    int
	Weight int
}

type Graph struct {
	Adj    map[int][]Edge
	vertex int
}

func (e *Graph) addEdge(s int, t int, w int) {
	edge := Edge{s, t, w}
	e.Adj[s] = append(e.Adj[s], edge)
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

//堆排序
func (h *PriorityQueue) SortHeap(heaps []Vertex) {
	length := len(heaps)
	length = length - 1
	if length == 1 {
		return
	}
	if length == 2 {
		h.AdjustHeap(length - 1)
	}
	for length > 0 {
		h.SliceNodeSwap(1, length)
		length--
		h.Heapfiy(length, 1)
	}
	//反序
	minPos := 1
	maxPos := h.count
	for minPos < maxPos {
		h.SliceNodeSwap(minPos, maxPos)
		minPos++
		maxPos--
	}
}

//自下而上调整
func (h *PriorityQueue) AdjustHeap(length int) {
	if length < 1 {
		return
	}
	if length == 2 {
		if h.nodes[length].Dist > h.nodes[length-1].Dist {
			h.SliceNodeSwap(length, length-1)
		}
		return
	}
	i := length
	for i/2 > 0 && h.nodes[i].Dist > h.nodes[i/2].Dist {
		h.SliceNodeSwap(i, i/2)
		i = i / 2
	}
	return
}

//输出heap
func heapShow(heaps []Vertex) {
	for one, value := range heaps {
		fmt.Println(one, value)
	}
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
	fmt.Println(length)
	h.Heapfiy(length, 1)
	heapShow(h.nodes)
	h.nodes = append(h.nodes[:length+1], h.nodes[length+2:]...)
	h.count--
	return top
}

func (g *Graph) Dijkstra(s int, t int) {
	predecessor := make([]int, g.vertex)
	vertexes := make([]Vertex, g.vertex)
	for i := 0; i < g.vertex; i++ {
		vertexes[i] = Vertex{i, MaxInt}
	}
	queue := PriorityQueue{}
	queue.count = g.vertex
	inqueue := make([]bool, g.vertex)
	vertexes[s].Dist = 0
	queue.InsertHeap(vertexes[s])
	inqueue[s] = true
	for len(queue.nodes) > 0 {
		minVertex := queue.GetTopHeap()
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
					queue.Heapfiy(len(queue.nodes), 1)
				} else {
					queue.InsertHeap(nextVertex)
					inqueue[nextVertex.Id] = true
				}
			}
		}
	}
	g.printPath(s, t, predecessor)
}
func (g *Graph) printPath(s int, t int, predecessor []int) {
	if s == t {
		return
	}
	g.printPath(s, predecessor[t], predecessor)
	println("->", t)
}
