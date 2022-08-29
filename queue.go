package main

import "fmt"

type Node struct {
	value string
	prev  *Node
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func newNode(value string, next *Node, prev *Node) *Node {
	return &Node{value: value, next: next, prev: prev}
}

func newQueue() *Queue {
	return &Queue{head: nil, tail: nil}
}
func (queue *Queue) enqueue(node *Node) {
	if node == nil {
		return
	}

	if queue.head == nil {
		queue.head = node
		queue.tail = node
		return
	}

	// Set node.next to head of queue
	node.next = queue.head
	node.prev = nil

	queue.head.prev = node

	// set head of queue to node
	queue.head = node
}

func (queue *Queue) dequeue() *Node {

	if queue.tail == nil {
		return nil
	}

	if queue.tail == queue.head {
		node := queue.head
		queue.head = nil
		queue.tail = nil
		return node
	}
	node := queue.tail

	queue.tail = node.prev

	queue.tail.next = nil

	(*node).prev = nil
	return node
}

func (queue *Queue) ToString() string {
	result := ""
	counter := 0
	currNode := queue.head
	for currNode != nil {
		result += fmt.Sprintf("\tNode %d: %s\n", counter, currNode.value)
		currNode = currNode.next
		counter++
	}
	return result
}
func testQueue() {
	queue := newQueue()

	queue.enqueue(newNode("ya", nil, nil))
	queue.enqueue(newNode("yeet", nil, nil))
	fmt.Println(queue.ToString())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.ToString())

	// dequeue empty
	fmt.Println(queue.dequeue())

	// one node
	queue.enqueue(newNode("ya", nil, nil))
	fmt.Println(queue.dequeue())
	fmt.Println(queue.ToString())
}
