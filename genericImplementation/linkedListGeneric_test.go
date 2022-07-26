package appGeneric_test

import (
	appGeneric "go-learning/genericImplementation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Add(t *testing.T) {
	list := appGeneric.NewLinkedList[string]()
	assert.Equal(t, 0, list.Length)
	list.Add("123")
	assert.Equal(t, 1, list.Length)
	list.Add("345")
	assert.Equal(t, 2, list.Length)
}

func TestList_Delete(t *testing.T) {
	list := appGeneric.NewLinkedList[string]()
	list.Delete("123")
	assert.Equal(t, 0, list.Length)

	list.Add("123")
	assert.Equal(t, 1, list.Length)
	list.Delete("123")
	assert.Equal(t, 0, list.Length)

	list.Add("123")
	list.Add("123")
	list.Delete("123")
	assert.Equal(t, 1, list.Length)

	list.Delete("567")
	assert.Equal(t, 1, list.Length)
}

func TestList_Find_String(t *testing.T) {
	list := appGeneric.NewLinkedList[string]()
	result := list.Find("123")
	assert.Equal(t, (*appGeneric.Node[string])(nil), result)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Value)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Value)
}

func TestList_Find_Int(t *testing.T) {
	list := appGeneric.NewLinkedList[int]()
	result := list.Find(123)
	assert.Equal(t, (*appGeneric.Node[int])(nil), result)

	list.Add(123)
	result = list.Find(123)
	assert.Equal(t, 123, result.Value)

	list.Add(123)
	result = list.Find(123)
	assert.Equal(t, 123, result.Value)
}

func TestList_FindCycle(t *testing.T) {
	list := appGeneric.NewLinkedList[int]()
	list.Add(123)
	list.Add(123)
	result := list.IsCycleInList()
	assert.Equal(t, false, result)
	head := list.Head
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = head
	result = list.IsCycleInList()
	assert.Equal(t, true, result)
}

func TestList_RevertList(t *testing.T) {
	list := appGeneric.NewLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	var listElementsArray []interface{}
	assert.Equal(t, 1, list.Head.Value)
	current := list.Head
	for current.Next != nil {
		listElementsArray = append(listElementsArray, current.Value)
		current = current.Next
	}
	listElementsArray = append(listElementsArray, current.Value)
	assert.Equal(t, 1, listElementsArray[0])
	assert.Equal(t, 2, listElementsArray[1])
	assert.Equal(t, 3, listElementsArray[2])
	list.RevertList()
	var revertElementsArray []interface{}
	current = list.Head
	for current.Next != nil {
		revertElementsArray = append(revertElementsArray, current.Value)
		current = current.Next
	}
	revertElementsArray = append(revertElementsArray, current.Value)
	assert.Equal(t, 3, revertElementsArray[0])
	assert.Equal(t, 2, revertElementsArray[1])
	assert.Equal(t, 1, revertElementsArray[2])
	assert.Equal(t, 3, list.Head.Value)
}
