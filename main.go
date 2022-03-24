package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	value int
	next  *node
}

const upperBoundRdmNum = 20
const sizeList = 9

func main() {
	l := genRandomLinkedList(sizeList)
	printLinkedList(l)
	l = reverse(l)
	printLinkedList(l)
}

func reverse(list *node) *node {
	start := list
	current := start
	next := current.next
	var aux *node
	for next.next != nil {
		aux = next.next
		next.next = current
		current = next
		next = aux
	}
	next.next = current
	start.next = nil
	return next
}

func genRandomLinkedList(n int) *node {
	if n == 0 {
		return nil
	}
	var list = &node{}
	current := list
	for i := 0; i < n; i++ {
		newNode := &node{}
		newNode.value = rand.Int() % upperBoundRdmNum
		current.next = newNode
		current = newNode
	}
	return list.next
}

func printLinkedList(list *node) {
	current := list
	fmt.Print("[")
	for current != nil {
		fmt.Printf("%v, ", current.value)
		current = current.next
	}
	fmt.Print("\b\b]\n")
}
