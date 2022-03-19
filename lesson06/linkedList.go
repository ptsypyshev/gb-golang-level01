package main

import (
	"fmt"
)

const (
	addFirst  = "addFirst"
	addMiddle = "addMiddle"
	addLast   = "addLast"
)

type Node struct {
	val  interface{}
	next *Node
}

type SingleList struct {
	length int
	head   *Node
}

func initSingleList() *SingleList {
	return &SingleList{}
}

func (list *SingleList) isEmpty() bool {
	return list.length == 0
}

func (list *SingleList) adder(newVal interface{}, operation string, n int) {
	node := &Node{val: newVal}
	switch {
	case list.isEmpty():
		list.head = node
	case operation == addFirst:
		node.next, list.head = list.head, node
	case operation == addLast:
		step := list.head
		for step.next != nil {
			step = step.next
		}
		step.next = node
	case operation == addMiddle:
		step := list.head
		for i := 0; i < n; i++ {
			step = step.next
		}
		step.next, node.next = node, step.next
	}
	list.length++
}

func (list *SingleList) addFirst(newVal interface{}) {
	list.adder(newVal, addFirst, 0)
}

func (list *SingleList) addLast(newVal interface{}) {
	list.adder(newVal, addLast, 0)
}

func (list *SingleList) addMiddle(newVal interface{}, n int) error {
	if n < 0 {
		return fmt.Errorf("you cannot use negative index of list")
	}
	if n > list.length {
		return fmt.Errorf("you cannot insert %v after %d because original list is too small", newVal, n)
	}
	list.adder(newVal, addMiddle, n)
	return nil
}

func (list *SingleList) getHead() (interface{}, error) {
	if list.isEmpty() {
		return "", fmt.Errorf("single List is empty")
	}
	return list.head.val, nil
}

func (list *SingleList) printList() {
	if list.isEmpty() {
		fmt.Println("Single List is empty")
		return
	}
	fmt.Printf("{ ")
	step := list.head
	for step != nil {
		fmt.Printf("%v ", step.val)
		step = step.next
	}
	fmt.Printf("}\n")
}

func (list *SingleList) reverseList() {
	step := list.head
	var prev *Node

	for step != nil {
		next := step.next
		step.next = prev
		prev = step
		step = next
	}
	list.head = prev
}

func main() {
	llist := initSingleList()
	llist.printList()
	fmt.Printf("Is your list empty? %t\n", llist.isEmpty())
	llist.addLast("last")
	fmt.Printf("Is your list empty? %t\n", llist.isEmpty())
	llist.addFirst(0)
	err1 := llist.addMiddle(1, 0)
	err2 := llist.addMiddle("middle", 1)
	err3 := llist.addMiddle("middle2", 5)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("You have some errors", err1, err2, err3)
	}

	res, err := llist.getHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	llist.printList()

	llist.reverseList()
	llist.printList()
}
