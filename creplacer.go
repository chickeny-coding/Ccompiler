package main

func commentFilter(n *Node) *Node {
	if n.c == nil {
		return n
	}
	if len(n.c[0].c) == 3 {
		return commentFilter(n.c[1])
	}
	n.c[1] = commentFilter(n.c[1])
	if len(n.c) == 4 {
		n.c[3] = commentFilter(n.c[3])
	}
	return n
}

func removeRedundancy(n *Node) *Node {
	if n.c == nil {
		return n
	}
	if len(n.c) == 4 && n.c[1].c == nil && n.c[3].c == nil {
		return &Node{BLOCK, nil}
	}
	n.c[1] = removeRedundancy(n.c[1])
	if len(n.c) == 4 {
		n.c[3] = removeRedundancy(n.c[3])
	}
	return n
}

func replacer(n *Node) *Node {
	if len(n.c) > 0 {
		n.c[0] = commentFilter(n.c[0])
		n.c[0] = removeRedundancy(n.c[0])
	}
	return n
}
