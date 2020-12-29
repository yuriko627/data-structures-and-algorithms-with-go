package main

import "fmt"

func main() {
	initial()
	output(head)
	removeNode(2)
	fmt.Printf("Delete a node Berkeley at index = 2:\n")
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

	var nodeOakland *Node = &Node{data: "Oakland", next: nil}
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

func removeNode(removedPosition int) {
	var p = head //p as in pointer
	var i = 0    //i as in index

	//move the pointer to a node which is right in front of the removed node
	for {
		if p.next == nil || i >= removedPosition-1 { //pointed node is tail or removedPosition is less than 1 (meaning removed node is head)
			break
		}
		p = p.next
		i++
	}

	//temporally store the removed node
	var removedNode = p.next
	p.next = p.next.next
	removedNode.next = nil
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
	fmt.Printf("End\n")
}
