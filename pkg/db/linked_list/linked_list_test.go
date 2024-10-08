package linkedlist_test

import (
	"testing"

	linkedlist "github.com/PepsiKingIV/KeyValueDB/pkg/db/linked_list"
	"github.com/stretchr/testify/assert"
)

func Test_NewLinkedList(t *testing.T) {
	newNode := linkedlist.NewLinkedList()
	assert.Equal(t, newNode, &linkedlist.Node{})
}

func Test_AddNode(t *testing.T) {
	type TestCase struct {
		Name     string
		Key      string
		Value    string
		Expected error
	}
	tcs := []TestCase{
		{
			Name:     "Basic case",
			Key:      "Key",
			Value:    "Value",
			Expected: nil,
		},
		{
			Name:     "Invalid key",
			Key:      "",
			Value:    "Value",
			Expected: nil,
		},
		{
			Name:     "Invalid Value",
			Key:      "Key",
			Value:    "",
			Expected: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			newNode := linkedlist.NewLinkedList()
			newNode.NextNode = &linkedlist.Node{}
			err := linkedlist.Add(newNode, tc.Key, tc.Value)
			assert.Equal(t, tc.Expected, err)
		})
	}
}

func Test_DeleteNode(t *testing.T) {
	type TestCase struct {
		Name          string
		NodeForDelete linkedlist.Node
		Node2         linkedlist.Node
		Expected      error
	}
	tcs := []TestCase{
		{
			Name: "Basic case",
			NodeForDelete: linkedlist.Node{
				Key:      "ForDelete",
				Value:    "ForDelete",
				NextNode: nil,
			},
			Node2: linkedlist.Node{
				Key:      "Key2",
				Value:    "Value2",
				NextNode: nil,
			},
			Expected: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			newNode := linkedlist.NewLinkedList()
			newNode.NextNode = &tc.NodeForDelete
			tc.NodeForDelete.NextNode = &tc.Node2
			err := linkedlist.Delete(newNode, tc.NodeForDelete.Key)
			assert.Equal(t, tc.Expected, err)
			assert.Equal(t, *newNode.NextNode, tc.Node2)
		})
	}
}

func Test_GetNode(t *testing.T) {
	type TestCase struct {
		Name     string
		Key      string
		Value    string
		Node1    linkedlist.Node
		Node2    linkedlist.Node
		Expected error
	}
	tcs := []TestCase{
		{
			Name:  "Basic case",
			Key:   "Key1",
			Value: "Value1",
			Node1: linkedlist.Node{
				Key:      "Key1",
				Value:    "Value1",
				NextNode: nil,
			},
			Node2: linkedlist.Node{
				Key:      "Key2",
				Value:    "Value2",
				NextNode: nil,
			},
			Expected: nil,
		},
		{
			Name:  "Not found",
			Key:   "Key",
			Value: "Value",
			Node1: linkedlist.Node{
				Key:      "",
				Value:    "",
				NextNode: nil,
			},
			Node2: linkedlist.Node{
				Key:      "",
				Value:    "",
				NextNode: nil,
			},
			Expected: linkedlist.ErrNotFound,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			newNode := linkedlist.NewLinkedList()
			newNode.NextNode = &tc.Node1
			tc.Node2.NextNode = &tc.Node2
			node, err := linkedlist.Get(newNode, tc.Key)
			assert.Equal(t, tc.Expected, err)
			if err == nil {
				assert.Equal(t, tc.Node1.Value, node.Value)
			}
		})
	}
}
