package main

import (
	"fmt"

	"container/list"
	"math/rand"
)

func doMergeSortOnSingleList() {
	fmt.Println("Starting merge sort on 1 list...")

	node1 := createList(25)
	printList(node1)

	node2, node3, _ := splitList(node1)

	node2 = sortList(node2)
	printList(node2)
	node3 = sortList(node3)
	printList(node3)

	node1 = mergeLists(node2, node3)
	printList(node1)

	fmt.Println("Finished merge sort on 1 list...")
}

func doMergeSortTwoLists() {
	fmt.Println("Starting merge sort on 2 lists...")

	node1 := createSortedList(5)
	printList(node1)
	node2 := createSortedList(10)
	printList(node2)
	mergedList := mergeLists(node1, node2)
	printList(mergedList)

	fmt.Println("Finished merge sort on 2 lists...")
}

func printList(lnode *list.List) {
	fmt.Print("{")
	for node := lnode.Front(); node != nil; node = node.Next() {
		fmt.Printf("%d,", node.Value)
	}
	fmt.Println("}")
}

func splitList(node *list.List) (*list.List, *list.List, error) {

	lSize := node.Len()
	idx := 0
	n1 := list.New()
	n2 := list.New()

	for n := node.Front(); n != nil; n = n.Next() {
		if idx < lSize/2 {
			n1.PushBack(n.Value)
		} else {
			n2.PushBack(n.Value)
		}
		idx += 1
	}

	return n1, n2, nil
}

func createList(num int) *list.List {
	node := list.New()

	for i := 0; i < num; i++ {
		node.PushBack(rand.Intn(10))
	}

	return node

}

func createSortedList(num int) *list.List {
	node := list.New()

	for i := 0; i < num; i++ {
		node.PushBack(rand.Intn(10))
	}

	node = sortList(node)

	return node
}

func sortList(lNode *list.List) *list.List {

	if lNode.Front() == nil {
		fmt.Println("Error: first node has nil Front")
		return nil
	}

	currentNode := lNode.Front()

	for currentNode != nil {
		nextNode := currentNode.Next()

		for nextNode != nil {
			// if current value is greater than next node value, swap
			if currentNode.Value.(int) > nextNode.Value.(int) {
				temp := currentNode.Value
				currentNode.Value = nextNode.Value
				nextNode.Value = temp
			}
			nextNode = nextNode.Next()
		}
		currentNode = currentNode.Next()
	}

	return lNode
}

func mergeLists(lNode1 *list.List, lNode2 *list.List) *list.List {

	resultNode := list.New()
	n1 := lNode1.Front()
	n2 := lNode2.Front()

	for n1 != nil && n2 != nil {
		if n1.Value.(int) < n2.Value.(int) {
			resultNode.PushBack(n1.Value)
			n1 = n1.Next()
		} else {
			resultNode.PushBack(n2.Value)
			n2 = n2.Next()
		}
	}

	// add the remaining nodes in the lists, if not empty
	for n1 != nil {
		resultNode.PushBack(n1.Value)
		n1 = n1.Next()
	}
	for n2 != nil {
		resultNode.PushBack(n2.Value)
		n2 = n2.Next()
	}

	return resultNode
}

func main() {
	doMergeSortTwoLists()
	doMergeSortOnSingleList()
}
