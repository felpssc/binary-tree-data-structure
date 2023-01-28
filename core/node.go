package core

type Node struct {
	Data  int
	Right *Node
	Left  *Node
}

func NewTreeNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

func (n *Node) Insert(data int) {
	if n == nil {
		return
	}

	if data >= n.Data {
		if n.Right == nil {
			n.Right = NewTreeNode(data)
		} else {
			n.Right.Insert(data)
		}
	} else {
		if n.Left == nil {
			n.Left = NewTreeNode(data)
		} else {
			n.Left.Insert(data)
		}
	}
}

func (n *Node) SetData(data int) {
	n.Data = data
}

func (n *Node) SetLeft(node *Node) {
	n.Left = node
}

func (n *Node) SetRight(node *Node) {
	n.Right = node
}

func (n Node) GetData() int {
	return n.Data
}
