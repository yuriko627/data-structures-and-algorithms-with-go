package main

import "fmt"

//initialize the list and append a new node at the end of the list
func main() {
	initial()

	fmt.Printf("Append a new node named 'Walnut' to the end: \n")
	add("Walnut")

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

func add(data string) {
	var newNode *Node = &Node{
		data: data,
		next: nil,
	}
	tail.next = newNode
	tail = newNode
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
