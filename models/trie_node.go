package models

type TrieNode struct {
	Children map[uint8]*TrieNode
	Word     string
	IsWord   bool
}

func NewTrieNode(word string, isWord bool) *TrieNode {
	return &TrieNode{
		Children: make(map[uint8]*TrieNode, 0),
		Word:     word,
		IsWord:   isWord,
	}
}

func (n *TrieNode) SplitNode(offset int) *TrieNode {
	if offset < len(n.Word) {
		afterOffset := n.Word[offset:]
		n.Word = n.Word[0:offset]
		newNode := NewTrieNode(afterOffset, n.IsWord)
		n.IsWord = false
		n.ModeChildrenTo(newNode)
		n.Children[afterOffset[0]] = newNode // add link to newly created node
	}
	return n
}

func (n *TrieNode) ModeChildrenTo(dest *TrieNode) {
	for k, v := range n.Children {
		dest.Children[k] = v
		delete(n.Children, k)
	}
}

func (n TrieNode) GetChild(s string) *TrieNode {
	if s == "" {
		return nil
	}
	node, ok := n.Children[s[0]]
	if !ok { // child node link doesn't exist
		return nil
	}
	return node
}

func (n *TrieNode) GetOrCreateChild(s string) *TrieNode {
	if s == "" {
		return nil
	}
	node, ok := n.Children[s[0]]
	if !ok { // child node link doesn't exist
		node = NewTrieNode(s, true) // create node
		n.Children[s[0]] = node     // create link
	}
	return node
}

func (n *TrieNode) GetStringsR(preVal string, acc *[]string) { // get strings recursive
	if n == nil {
		return
	}
	preVal += n.Word
	if n.IsWord {
		*acc = append(*acc, preVal)
		//fmt.Print(" | ", preVal)
	}
	for _, v := range n.Children {
		v.GetStringsR(preVal, acc)
	}
}
