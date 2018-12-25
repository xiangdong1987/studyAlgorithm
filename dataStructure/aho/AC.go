package aho

import (
	"fmt"
)

type AcNode struct {
	data     int32
	isEnd    bool
	children []*AcNode
	length   int
	fail     *AcNode
}

type AcTrie struct {
	root       *AcNode
	dictionary map[int32]int
	dicLength  int
}

func TestAcTrie() {
	arrList := []string{"冬天", "夏天", "秋天", "春天"}
	acTrie := AcTrie{}
	acTrie.dictionary = make(map[int32]int)
	acTrie.initDictionary(arrList)
	acTrie.root = &AcNode{}
	acTrie.root.children = make([]*AcNode, acTrie.dicLength)
	//acTrie.root.fail=acTrie.root
	for _, value := range arrList {
		acTrie.addWord(value)
	}
	//初始错误指针
	acTrie.initFailPoint()
	//fmt.Println(acTrie)
	acTrie.match("春天在哪里，春天在哪里，在小鸟的肚子里，冬天不错。")
}

//初始化字典
func (ac *AcTrie) initDictionary(wordList []string) {
	i := 0
	for _, value := range wordList {
		for _, c := range value {
			if _, ok := ac.dictionary[c]; ok {
				continue
			} else {
				ac.dictionary[c] = i
				i++
				continue
			}
		}
	}
	ac.dicLength = i
}

//取出切片第一个
func pop(list []*AcNode) (*AcNode, []*AcNode) {
	if len(list) > 0 {
		a := list[0]
		b := list[1:]
		return a, b
	} else {
		return &AcNode{}, list
	}
}

//推入切片
func push(list []*AcNode, value *AcNode) []*AcNode {
	result := append(list, value)
	return result
}

//构建trie树
func (ac *AcTrie) addWord(word string) {
	//fmt.Println(ac.dictionary)
	nowNode := ac.root
	i := 1
	for _, c := range word {
		if nowNode.children[ac.dictionary[c]] != nil {
			nowNode = nowNode.children[ac.dictionary[c]]
		} else {
			newNode := &AcNode{}
			newNode.children = make([]*AcNode, ac.dicLength)
			nowNode.children[ac.dictionary[c]] = newNode
			nowNode.data = c
			nowNode = newNode
		}
		i++
		if i == len([]rune(word)) {
			//fmt.Println(i)
			nowNode.isEnd = true
			nowNode.length = i
		}
	}
}

//初始化错误指针
func (ac *AcTrie) initFailPoint() {
	var queue []*AcNode
	queue = push(queue, ac.root)
	var p *AcNode
	for len(queue) > 0 {
		p, queue = pop(queue)
		for i := 0; i < len(p.children); i++ {
			pc := p.children[i]
			if pc == nil {
				continue
			}
			if pc == ac.root {
				p.fail = ac.root
			} else {
				q := p.fail
				for q != nil {
					qc := q.children[ac.dictionary[pc.data]]
					if qc != nil {
						pc.fail = qc
						break
					}
					q = q.fail
				}
				if q == nil {
					pc.fail = ac.root
				}
			}
			push(queue, pc)
		}
	}
}

func (ac *AcTrie) match(str string) {
	p := ac.root
	i := 1
	for _, c := range str {
		//先判断字符是否在词典中
		if _, ok := ac.dictionary[c]; !ok {
			//从头再来
			i++
			p = ac.root
			continue
		}
		for p.children[ac.dictionary[c]] == nil && p != ac.root {
			p = p.fail
		}
		p = p.children[ac.dictionary[c]]
		//fmt.Println(p)
		if p == nil {
			//从头再来
			p = ac.root
		}
		tmp := p
		for tmp != ac.root && tmp != nil {
			if tmp.isEnd == true {
				pos := i - tmp.length + 1
				fmt.Println("Word is mach, pos is", pos, "length is", tmp.length)
			}
			tmp = tmp.fail
		}
		i++
	}
}
