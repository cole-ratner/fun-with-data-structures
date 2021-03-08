package main

import (
	"errors"
	"fmt"
)

// node is a subcomponent of the primary linkedList structure
type node struct {
	data int
	next *node
}

// linkedList is the primary data structure
type linkedList struct {
	head *node
	length int
}

// adds a node to the head of the linked list
func (l *linkedList) prepend(n *node) {
	prev := l.head
	l.head = n
	l.head.next = prev
	l.length++
}

// helper func that gets the last node in the list
func (l *linkedList) getLast() *node{
	var last *node

	for node := l.head; node != nil; node = node.next {
		if node.next == nil {
			last = node
		}
	}
	return last
}

// adds a node to the end of the list
func (l *linkedList) append(n *node) {
	last := l.getLast()
	last.next = n
	l.length++
}

// iterates through the entire list and prints all the nodes
func (l * linkedList) printAll() {
	for node := l.head; node != nil; node = node.next {
		fmt.Printf("Found a Node! {data: %d, next: %v}\n", node.data, node.next) 
	}
}

// looks for a node with a specific node.data value
func (l *linkedList) findNodeWithValue(v int) (*node, error) {
	var wanted *node
	
	for node := l.head; node != nil; node = node.next {
		if node.data == v {
			wanted = node
			return wanted, nil
		}
	}
	return nil, errors.New("couldn't find a node with the requested value")
}

// inserts a new node after a node with a specific node.data value
func (l *linkedList) insertAfterNodeWithValue(v int, n *node) error {
	prev, err := l.findNodeWithValue(v)
	if err != nil {
		return err
	}
	
	if prev.next == nil {
		prev.next = n
	}

	n.next = prev.next
	prev.next = n
	l.length++
	return nil
}


func main() {
	//instantiate a new linkedList
	l := linkedList{}

// 	adding nodes to the linkedlist in various ways
	l.prepend(&node{data:1})
	l.prepend(&node{data:3})
	l.append(&node{data:15})
	
	err := l.insertAfterNodeWithValue(1, &node{data:33})
	if err != nil {
		fmt.Println(err)
	}

	l.append(&node{data:29})
//	//////
	
	// getting & printing the last node
	last := l.getLast()
	fmt.Printf("Last Node: %v\n\n", last)
	
	// getting a node with a value
	ok, err := l.findNodeWithValue(1)
	if err != nil {fmt.Println(err)}
	if 	ok != nil {fmt.Printf("Requested Node: %v\n\n", ok)}

	// printing all nodes
	fmt.Println("Now printing all nodes:")
	l.printAll()
}