package topologicalSort

import (
	"fmt"
)

//有向图
type graph struct {
	vertex int           //顶点
	list   map[int][]int //连接表边
}

//添加边
func (g *graph) addVertex(t int, s int) {
	g.list[t] = push(g.list[t], s)
}

func TestKhanSort() {
	g := NewGraph(8)
	g.addVertex(2, 1)
	g.addVertex(3, 1)
	g.addVertex(7, 1)
	g.addVertex(4, 2)
	g.addVertex(5, 2)
	g.addVertex(8, 7)
	g.DfsSort()
}

//创建图
func NewGraph(v int) *graph {
	g := new(graph)
	g.vertex = v
	g.list = map[int][]int{}
	i := 0
	for i < v {
		g.list[i] = make([]int, 0)
		i++
	}
	return g
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

//添加边
func (g *graph) KhanSort() {
	var inDegree = make(map[int]int)
	var queue []int
	for i := 1; i <= g.vertex; i++ {
		for _, m := range g.list[i] {
			inDegree[m]++
		}
	}
	for i := 1; i <= g.vertex; i++ {

		if inDegree[i] == 0 {
			queue = push(queue, i)
		}
	}
	for len(queue) > 0 {
		var now int
		now, queue = pop(queue)
		fmt.Println("->", now)
		for _, k := range g.list[now] {
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = push(queue, k)
			}
		}
	}
}

func (g *graph) DfsSort() {
	inverseList := make(map[int][]int)
	//初始化逆向邻接表
	for i := 1; i <= g.vertex; i++ {
		for _, k := range g.list[i] {
			inverseList[k] = append(inverseList[k], i)
		}
	}
	visited := make([]bool, g.vertex+1)
	visited[0] = true
	for i := 1; i <= g.vertex; i++ {
		if visited[i] == false {
			visited[i] = true
			dfs(i, inverseList, visited)
		}
	}

}

func dfs(vertex int, inverseList map[int][]int, visited []bool) {
	for _, w := range inverseList[vertex] {
		if visited[w] == true {
			continue
		} else {
			visited[w] = true
			dfs(w, inverseList, visited)
		}
	}
	fmt.Println("->", vertex)
}
