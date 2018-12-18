package graph

import "fmt"

//无向图
type graph struct {
	vertex int
	list   map[int][]int
	found  bool
}

func TestGraph() {
	g := NewGraph(8)
	g.addVertex(1, 2)
	g.addVertex(1, 4)
	g.addVertex(2, 3)
	g.addVertex(2, 5)
	g.addVertex(4, 5)
	g.addVertex(3, 6)
	g.addVertex(5, 6)
	g.addVertex(5, 7)
	g.addVertex(6, 8)
	g.addVertex(7, 8)
	g.bfs(1, 7)
}

//创建图
func NewGraph(v int) *graph {
	g := new(graph)
	g.vertex = v
	g.found = false
	g.list = map[int][]int{}
	i := 0
	for i < v {
		g.list[i] = make([]int, 0)
		i++
	}
	return g
}

//添加边
func (g *graph) addVertex(t int, s int) {
	g.list[t] = push(g.list[t], s)
	g.list[s] = push(g.list[s], t)
}

//取出切片第一个
func pop(list []int) (int, []int) {
	if len(list) > 0 {
		a := list[0]
		b := list[1:]
		return a, b
	} else {
		return -1, list
	}
}

//推入切片
func push(list []int, value int) []int {
	result := append(list, value)
	return result
}

//广度优先搜索
func (g *graph) bfs(s int, t int) {
	if s == t {
		return
	}
	visited := make([]bool, g.vertex+1)
	var queue []int
	queue = append(queue, s)
	prev := make([]int, g.vertex+1)
	i := 0
	for i < len(prev) {
		prev[i] = -1
		i++
	}
	for len(queue) != 0 {
		var w int
		w, queue = pop(queue)
		for j := 0; j < len(g.list[w]); j++ {
			q := g.list[w][j]
			fmt.Println(q)
			if !visited[q] {
				prev[q] = w
				if q == t {
					fmt.Println(prev)
					printPath(prev, s, t)
					return
				}
				visited[q] = true
				queue = append(queue, q)
			}
		}
	}
}

//深度优先搜索
func (g *graph) dfs(s int, t int) {
	prev := make([]int, g.vertex+1)
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	visit := make([]bool, g.vertex+1)
	g.recurDsf(s, t, prev, visit)
	fmt.Println(prev)
	printPath(prev, s, t)
}

func (g *graph) recurDsf(w int, t int, prev []int, visited []bool) {
	if g.found {
		return
	}
	if w == t {
		g.found = true
		return
	}
	visited[w] = true
	for i := 0; i < len(g.list[w]); i++ {
		q := g.list[w][i]
		if !visited[q] {
			prev[q] = w
			g.recurDsf(g.list[w][i], t, prev, visited)
		}
	}
}

//深度优先搜做
func printPath(prev []int, s int, t int) {
	if prev[t] != -1 && s != t {
		printPath(prev, s, prev[t])
	}
	fmt.Println(t, "  ")
}
