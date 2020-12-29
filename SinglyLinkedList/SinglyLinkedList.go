package main

import "fmt"

//initialize and output the list
func main() {
	initial()
	output(head)
}

type Node struct {
	data string
	next *Node //define next
}

//head node; the first node
var head *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.next = nil

	// derefernce the node pointer from its memory addreess
	// to the currenct variable at a new memory address (&Node: new referenced address)
	var nodeOakland *Node = &Node{
		data: "Oakland",
		next: nil,
	}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{
		data: "Berkeley",
		next: nil,
	}
	nodeOakland.next = nodeBerkeley

	//tail node: the last node

	var tail *Node = &Node{
		data: "Fremont",
		next: nil,
	}
	nodeBerkeley.next = tail
}

func output(node *Node) {
	var p = node
	for {
		if p == nil {
			break
		}

		fmt.Printf("%s->", p.data)
		p = p.next
	}

	fmt.Printf("End\n\n")
}
