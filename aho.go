package aho

// Aho implements Aho-Corasick Automaton
type Aho struct {
	Built bool
	root  *node
}

func New() *Aho {
	return &Aho{false, newNode()}
}

func (a *Aho) AddWord(word string) {
	n := a.root
	for _, r := range word {
		child, ok := n.Children[r]
		if !ok {
			newNode := newNode()
			newNode.Parent = n
			newNode.ParentCh = r
			n.Children[r] = newNode
			n = n.Children[r]
		} else {
			n = child
		}
	}

	n.Word = &word
}

func (a *Aho) Build() {
	stack := NewStack()
	stack.Push(a.root)
	for !stack.Empty() {
		curNode := stack.Pop()
		for _, child := range curNode.Children {
			stack.Push(child)
		}

		if curNode == a.root {
			curNode.SuffixLink = a.root
			continue
		}

		if curNode.Parent == a.root {
			curNode.SuffixLink = a.root
			continue
		}

		if suffixLink, ok := curNode.Parent.SuffixLink.Children[curNode.ParentCh]; ok {
			curNode.SuffixLink = suffixLink
			continue
		}

		if suffixLink, ok := a.root.Children[curNode.ParentCh]; ok {
			curNode.SuffixLink = suffixLink
			continue
		}

		curNode.SuffixLink = a.root
	}

	a.Built = true
}

func (a *Aho) FindAll(s string) []string {
	a.assertBuilt()

	words := make(map[string]bool)
	n := a.root
	for _, r := range s {
		if n.Word != nil {
			words[*n.Word] = true
		}

		if child, ok := n.Children[r]; ok {
			n = child
			continue
		}

		for n.SuffixLink != n {
			n = n.SuffixLink
			if child, ok := n.Children[r]; ok {
				n = child
				break
			}
		}
	}

	wordList := make([]string, len(words))
	i := 0
	for w := range words {
		wordList[i] = w
		i++
	}

	return wordList
}

func (a *Aho) assertBuilt() {
	if !a.Built {
		panic("build first")
	}
}

type node struct {
	Parent     *node
	ParentCh   rune
	Children   map[rune]*node
	Word       *string
	SuffixLink *node
}

func newNode() *node {
	return &node{
		Children: make(map[rune]*node),
	}
}
