package app_generic_test

import (
	app_generic "go-learning/internal/app/generic_implementation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Add(t *testing.T) {
	list := app_generic.InitList[string]()
	assert.Equal(t, 0, list.Length)
	list.Add("123")
	assert.Equal(t, 1, list.Length)
	list.Add("345")
	assert.Equal(t, 2, list.Length)
}

func TestList_Delete(t *testing.T) {
	list := app_generic.InitList[string]()
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
	list := app_generic.InitList[string]()
	result := list.Find("123")
	assert.Equal(t, (*app_generic.Node[string])(nil), result)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Key)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Key)
}

func TestList_Find_Int(t *testing.T) {
	list := app_generic.InitList[int]()
	result := list.Find(123)
	assert.Equal(t, (*app_generic.Node[int])(nil), result)

	list.Add(123)
	result = list.Find(123)
	assert.Equal(t, 123, result.Key)

	list.Add(123)
	result = list.Find(123)
	assert.Equal(t, 123, result.Key)
}
