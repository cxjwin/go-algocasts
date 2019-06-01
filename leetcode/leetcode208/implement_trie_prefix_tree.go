package leetcode208

type TrieNode struct {
	children  [26]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.root = &TrieNode{endOfWord: false}
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	cur := this.root
	for _, v := range word {
		idx := int(v - 'a')
		if cur.children[idx] == nil {
			cur.children[idx] = &TrieNode{endOfWord: false}
		}
		cur = cur.children[idx]
	}
	cur.endOfWord = true
}

func (this *Trie) endNode(word string) *TrieNode {
	if len(word) == 0 {
		return this.root
	}

	cur := this.root
	for _, v := range word {
		if cur == nil {
			return cur
		}

		idx := int(v - 'a')
		cur = cur.children[idx]
	}
	return cur
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.endNode(word)
	return node != nil && node.endOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.endNode(prefix) != nil
}
