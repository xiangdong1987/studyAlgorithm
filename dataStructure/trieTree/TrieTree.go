package trieTree

import "fmt"

type TrieNode struct {
	value      int
	dictionary [26]*TrieNode
}

type TrieTree struct {
	root *TrieNode
}

func TestTrieTree() {
	tree := createTree()
	//fmt.Println(tree)
	flag := tree.findWord("her")
	fmt.Println(flag)
	flag = tree.findWord("x")
	fmt.Println(flag)
}

func createTree() TrieTree {
	arrList := []string{"how", "hi", "her", "hello", "so", "see"}
	myTree := TrieTree{}
	//添加跟节点
	myTree.root = &TrieNode{}
	for _, value := range arrList {
		myTree.addWord(value)
	}
	return myTree
}
func (t *TrieTree) addWord(word string) {
	fmt.Println(word)
	nowNode := t.root
	a := int('a')
	var char int
	for i := 0; i < len(word); i++ {
		char = int(word[i])
		if nowNode.dictionary[char-a] != nil {
			nowNode = nowNode.dictionary[char-a]
			continue
		} else {
			newNode := &TrieNode{}
			nowNode.dictionary[char-a] = newNode
			nowNode = newNode
			continue
		}
	}
}

func (t *TrieTree) findWord(word string) int {
	nowNode := t.root
	a := int('a')
	var char int
	for i := 0; i < len(word); i++ {
		char = int(word[i])
		if nowNode.dictionary[char-a] != nil {
			return 0
		} else {
			nowNode = nowNode.dictionary[char-a]
		}
		if i == len(word)-1 {
			return 1
		}
	}
	return 0
}
