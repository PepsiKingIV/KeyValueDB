package linkedlist

import (
	"runtime"
)

type LinkedList interface {
	Add(*Node, string, string) error
	Get(*Node, string) (*Node, error)
	Delete(*Node, string) error
}

type Node struct {
	Key      string
	Value    string
	NextNode *Node
}

func NewLinkedList() *Node {
	return &Node{}
}

// записывать надо в начало. Запись сделанная последней, с большей вероятностью понадобится
// К тому же это позволит сократить время записи => двойной профит Запись + Чтение
func Add(node *Node, key string, value string) error {
	// nextNode := node
	// for nextNode.NextNode != nil {
	// 	nextNode = nextNode.NextNode
	// }

	// nextNode.NextNode = &Node{
	// 	Key:      key,
	// 	Value:    value,
	// 	NextNode: nil}
	node.NextNode = &Node{
		Key:      key,
		Value:    value,
		NextNode: node.NextNode,
	}
	return nil
}

func Get(node *Node, key string) (*Node, error) {
	nextNode := node
	for nextNode.NextNode != nil && nextNode.Key != key {
		nextNode = nextNode.NextNode
	}
	if nextNode.Key != key {
		return nil, ErrNotFound
	}
	return nextNode, nil
}

func Delete(node *Node, key string) error {
	var nextNode *Node = node.NextNode
	var currentNode *Node = node
	if nextNode == nil {
		return ErrNotFound
	}
	for nextNode.Key != key {
		currentNode = nextNode
		nextNode = nextNode.NextNode
		if currentNode.NextNode == nil {
			return ErrNotFound
		}
	}
	currentNode.NextNode = nextNode.NextNode
	runtime.GC()
	return nil
}
