package appInterface_test

import (
	appInterface "go-learning/interfaceImplementation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Add(t *testing.T) {
	list := appInterface.NewLinkedList()
	assert.Equal(t, 0, list.Length)
	list.Add("123")
	assert.Equal(t, 1, list.Length)
	list.Add("345")
	assert.Equal(t, 2, list.Length)
}

func TestList_Delete(t *testing.T) {
	list := appInterface.NewLinkedList()
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
	list := appInterface.NewLinkedList()
	result := list.Find("123")
	assert.Equal(t, (*appInterface.Node)(nil), result)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Value)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Value)
}

func TestList_Find_Int(t *testing.T) {
	list := appInterface.NewLinkedList()
	result := list.Find(123)
	assert.Equal(t, (*appInterface.Node)(nil), result)

	list.Add(123)
	result = list.Find(123)
	assert.Equal(t, 123, result.Value)

	list.Add(123)
	result = list.Find(123)
	assert.Equal(t, 123, result.Value)
}
