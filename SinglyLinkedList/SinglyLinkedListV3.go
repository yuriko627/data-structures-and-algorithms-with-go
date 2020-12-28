package main

import "fmt"

// Initialize the list, insert a new node at index=2
func main() {
	initial()
	fmt.Printf("Insert a new node Walnut at index=2: \n")
	insert(2, "Walnut")
	output(head)
}

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.next = nil

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

	tail.data = "Fremont"
	tail.next = nil
	nodeBerkeley.next = tail

}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0

	//insert a new node in the poisition defined with insertPosition
	for {
		if p.next == nil || i > insertPosition-1 {
			break
		}

		p = p.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = p.next
	p.next = newNode
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
