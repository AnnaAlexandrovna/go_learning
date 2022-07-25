package app_test

import (
	"go-learning/internal/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Add(t *testing.T) {
	list := app.InitList()
	assert.Equal(t, 0, list.Length)
	list.Add("123")
	assert.Equal(t, 1, list.Length)
	list.Add("345")
	assert.Equal(t, 2, list.Length)
}

func TestList_Delete(t *testing.T) {
	list := app.InitList()
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

func TestList_Find(t *testing.T) {
	list := app.InitList()
	result := list.Find("123")
	assert.Equal(t, (*app.Node)(nil), result)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Key)

	list.Add("123")
	result = list.Find("123")
	assert.Equal(t, "123", result.Key)
}
