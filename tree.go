package tree

type Node struct {
	Value string
	Child map[string]*Node
}

type Tree struct {
	Root *Node
}

// BuildTree 通过给定敏感词数组建树
func BuildTree(words []string) (*Tree, error) {
	root := &Node{Child: map[string]*Node{}}
	for _, word := range words {
		runes := []rune(word)
		node := root
		for _, r := range runes {
			if node.Child == nil {
				node.Child = map[string]*Node{}
			}
			if v, ok := node.Child[string(r)]; ok{
				node = v
			} else {
				temp := &Node{Value: string(r)}
				node.Child[string(r)] = temp
				node = temp
			}
		}
	}
	return &Tree{Root: root}, nil
}

func (t *Tree)FindFirstMatchedWord(text string) (ok bool, word string) {
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if ok, word := t.findFirstMatchedWord(runes[i:]); ok {
			return ok, word
		}
	}
	return false, ""
}

func (t *Tree)findFirstMatchedWord(runes []rune) (ok bool, word string) {
	cur := t.Root
	end := 0
	for i, r := range runes {
		if cur.Child == nil {
			end = i
			break
		}
		if next := cur.Child[string(r)]; next != nil {
			cur = next
		} else {
			return false, ""
		}
	}
	if end != 0 {
		return true, string(runes[:end])
	}
	return false, ""
}