package models

type Trie struct {
	Root *TrieNode
}

func NewCompressedTrie() *Trie {
	return &Trie{NewTrieNode("", false)}
}

func (t *Trie) GetStrings() []string {
	acc := make([]string, 0)
	t.Root.GetStringsR("", &acc)
	return acc
}

func (t *Trie) Search(s string) *TrieNode {
	var curr **TrieNode = &t.Root
	n := len(s)
	var searchNode *TrieNode = nil
	for i, k := 0, 0; i < n && curr != nil; {
		for i < n && k < len((*curr).Word) && s[i] == (*curr).Word[k] {
			k++
			i++
		}
		if k < len((*curr).Word) {
			break
		}
		if i < n {
			childNode := (*curr).GetChild(s[i:])
			if childNode == nil {
				break
			}
			curr = &childNode
			k = 0 // reset node word index
		} else {
			if (*curr).IsWord {
				searchNode = *curr
			}
		}
	}
	return searchNode
}

func (t *Trie) Add(s string) {
	var curr **TrieNode = &t.Root
	n := len(s)
	for i, k := 0, 0; i < n && curr != nil; {
		for i < n && k < len((*curr).Word) && s[i] == (*curr).Word[k] {
			k++
			i++
		}
		if k < len((*curr).Word) {
			splitNode := (*curr).SplitNode(k)
			curr = &splitNode
			i--
			k--
			continue
		}
		if i < n {
			childNode := (*curr).GetOrCreateChild(s[i:])
			curr = &childNode
			k = 0 // reset node word index
		} else {
			(*curr).IsWord = true
		}
	}
}
