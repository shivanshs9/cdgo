package pkg

import (
	"github.com/perlin-network/noise"
)

type Node struct {
	*noise.Node
	Nick string
}

func NewNode(node *noise.Node, nick string) *Node {
	resNode := &Node{node, nick}
	resNode.Nick = nick
	return resNode
}
