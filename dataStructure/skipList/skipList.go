package skipList

import "fmt"

type Node struct {
	Data      int
	NextPoint *Node
	PrePoint  *Node
	NextLevel *Node
}

type LinkedList struct {
	Head    *Node
	Current *Node
	Tail    *Node
	Length  int
	Level   int
}

type SkipList struct {
	List        LinkedList
	FirstIndex  LinkedList
	SecondIndex LinkedList
	//ThirdIndex  LinkedList
}

func InitSkipList() {
	data := []int{11, 12, 13, 19, 21, 31, 33, 42, 51, 62}
	sl := SkipList{}
	sl.initSkip(data)
	//sl.find(19)
	sl.add(11)
	showSkipList(sl)
}

func showSkipLinkedList(link LinkedList, name int) {
	var currentNode *Node
	currentNode = link.Head
	for {
		i := 1
		fmt.Print(name, "-Node:", currentNode.Data)
		if currentNode.NextPoint == nil {
			break
		} else {
			currentNode = currentNode.NextPoint
		}
		if name == 1 {
			fmt.Print("-------->")
		} else if name == 2 {
			for i <= 3 {
				fmt.Print("-------->")
				i++
			}
		} else {
			for i <= 7 {
				fmt.Print("-------->")
				i++
			}
		}
	}
	fmt.Println("")
}

func (sl *SkipList) initSkip(list []int) {
	sl.List = LinkedList{}
	sl.FirstIndex = LinkedList{}
	sl.SecondIndex = LinkedList{}
	var currentNode *Node
	for i := 0; i < len(list); i++ {
		currentNode = new(Node)
		currentNode.Data = list[i]
		addNode(sl, currentNode)
		//insertToLink(&sl.List, currentNode)
	}
	showSkipList(*sl)
}
func (sl *SkipList) find(x int) {
	var current *Node
	current = sl.SecondIndex.Head
	if current.Data == x {
		fmt.Println(current.Data)
		return
	}
	if x < current.Data {
		panic("No exist in skipList")
		return
	}
	for {
		if x > current.Data {
			fmt.Println(current.Data)
			current = current.NextPoint
		} else if x < current.Data {
			//下到底层索引
			fmt.Println(current.Data)
			current = current.PrePoint.NextLevel.NextPoint
		} else {
			fmt.Println(current.Data)
			return
		}
	}
}
func (sl *SkipList) add(x int) {
	var current *Node
	current = sl.SecondIndex.Head
	if current.Data == x {
		panic("Had existed in skipList")
		return
	}
	if x < current.Data {
		newNode2 := new(Node)
		newNode2.Data = x
		newNode2.NextPoint = sl.SecondIndex.Head
		sl.SecondIndex.Head.PrePoint = newNode2
		sl.SecondIndex.Head = newNode2
		newNode1 := new(Node)
		newNode1.Data = x
		newNode1.NextPoint = sl.FirstIndex.Head
		sl.FirstIndex.Head.PrePoint = newNode1
		sl.FirstIndex.Head = newNode1
		newNode := new(Node)
		newNode.Data = x
		newNode.NextPoint = sl.List.Head
		sl.SecondIndex.Head.PrePoint = newNode
		sl.List.Head = newNode
		return
	}
	for {
		if x > current.Data {
			if current.NextPoint == nil {
				if current.NextLevel != nil {
					current = current.NextLevel
				} else {
					//插入
					newNode := new(Node)
					newNode.Data = x
					current.NextPoint = newNode
					newNode.PrePoint = current
					return
				}
			} else {
				fmt.Println(current.Data)
				current = current.NextPoint
			}

		} else if x < current.Data {
			//向下去寻找第一个大于x的值
			fmt.Println(current.Data)
			if current.PrePoint.NextLevel != nil {
				current = current.PrePoint.NextLevel.NextPoint
			} else {
				//插入
				newNode := new(Node)
				newNode.Data = x
				current.PrePoint.NextPoint = newNode
				newNode.NextPoint = current
				current.PrePoint = newNode
				return
			}
		} else {
			fmt.Println(current.Data)
			return
		}
	}
}
func showSkipList(sl SkipList) {
	showSkipLinkedList(sl.SecondIndex, 3)
	fmt.Println("")
	showSkipLinkedList(sl.FirstIndex, 2)
	fmt.Println("")
	showSkipLinkedList(sl.List, 1)
}
func addNode(skipList *SkipList, node *Node) {
	insertToLink(&skipList.List, node)
	if skipList.FirstIndex.Length == 0 || ((skipList.List.Length-1)%2 == 0 && skipList.List.Length > 2) {
		newNode := new(Node)
		newNode.Data = node.Data
		newNode.NextLevel = node
		insertToLink(&skipList.FirstIndex, newNode)
		if skipList.SecondIndex.Length == 0 || ((skipList.FirstIndex.Length-1)%2 == 0 && skipList.FirstIndex.Length > 2) {
			newNode2 := new(Node)
			newNode2.Data = node.Data
			newNode2.NextLevel = newNode
			insertToLink(&skipList.SecondIndex, newNode2)
		}
	}
}

func insertToLink(link *LinkedList, node *Node) {
	if link.Head == nil {
		link.Head = node
		link.Tail = node
		link.Current = node
	} else {
		link.Tail.NextPoint = node
		node.PrePoint = link.Tail
		link.Tail = node
	}
	link.Length++
}
