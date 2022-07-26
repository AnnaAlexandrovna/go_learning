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
}
