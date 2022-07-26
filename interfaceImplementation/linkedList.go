package appInterface

import "reflect"

type Node struct {
	Next  *Node
	Value interface{}
}
type List struct {
	Head   *Node
	Length int
}

func NewLinkedList() *List {
	return &List{}
}

func (s *List) Add(value interface{}) {
	el := &Node{
		Value: value,
	}
	if s.Head == nil {
		s.Head = el
	} else {
		current := s.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = el
	}
	s.Length++
}

func (s *List) Delete(value interface{}) {
	if s.Head == nil {
		return
	}
	var prev *Node
	current := s.Head
	if current == s.Head && s.Length == 1 && current.Value == value {
		s.Head = nil
		s.Length--
		return
	}
	for current.Next != nil {
		if current.Value == value {
			if current == s.Head {
				s.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			s.Length--
			return
		}
		prev = current
		current = current.Next
	}
}

func (s *List) Find(value interface{}) *Node {
	if s.Head == nil {
		return nil
	}
	var foundNode *Node
	current := s.Head
	if current == s.Head && s.Length == 1 && current.Value == value {
		return current
	}
	for current.Next != nil {
		if current.Value == value {
			foundNode = current
			break
		}
		current = current.Next
	}
	return foundNode
}

func (s *List) IsCycleInList() bool {
	if s.Head == nil {
		return false
	}
	haveCycle := false
	current := s.Head
	setOfNodes := make(map[*Node]bool)
	setOfNodes[s.Head] = true
	for current.Next != nil {
		if setOfNodes[current.Next] {
			return true
		}
		current = current.Next
	}
	return haveCycle
}

func (s *List) RevertList() {
	isCycleInList := s.IsCycleInList()
	if isCycleInList {
		return
	}
	if s.Head == nil {
		return
	}
	var prev *Node
	current := s.Head
	for current.Next != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	current.Next = prev
	s.Head = current
}

func ConcatSortedLists(list1 *List, list2 *List) *List {
	element1 := list1.Head
	element2 := list2.Head
	list := NewLinkedList()
	for element1 != nil || element2 != nil {
		if element1 != nil &&
			(element2 == nil ||
				compare(element2.Value, element1.Value)) {
			list.Add(element1.Value)
			element1 = element1.Next
		} else if element2 != nil &&
			(element1 == nil ||
				element1.Value == element2.Value ||
				compare(element1.Value, element2.Value)) {
			list.Add(element2.Value)
			element2 = element2.Next
		}
	}
	return list
}

func compare(value1 interface{}, value2 interface{}) bool {
	switch reflect.TypeOf(value1).Kind() {
	case reflect.Int:
		v1 := reflect.ValueOf(value1).Interface().(int)
		v2 := reflect.ValueOf(value2).Interface().(int)
		return v1 > v2
	case reflect.Int8:
		v1 := reflect.ValueOf(value1).Interface().(int8)
		v2 := reflect.ValueOf(value2).Interface().(int8)
		return v1 > v2
	case reflect.Int16:
		v1 := reflect.ValueOf(value1).Interface().(int16)
		v2 := reflect.ValueOf(value2).Interface().(int16)
		return v1 > v2
	case reflect.Int32:
		v1 := reflect.ValueOf(value1).Interface().(int32)
		v2 := reflect.ValueOf(value2).Interface().(int32)
		return v1 > v2
	case reflect.Int64:
		v1 := reflect.ValueOf(value1).Interface().(int64)
		v2 := reflect.ValueOf(value2).Interface().(int64)
		return v1 > v2
	case reflect.Uint:
		v1 := reflect.ValueOf(value1).Interface().(uint)
		v2 := reflect.ValueOf(value2).Interface().(uint)
		return v1 > v2
	case reflect.Uint8:
		v1 := reflect.ValueOf(value1).Interface().(uint8)
		v2 := reflect.ValueOf(value2).Interface().(uint8)
		return v1 > v2
	case reflect.Uint16:
		v1 := reflect.ValueOf(value1).Interface().(uint16)
		v2 := reflect.ValueOf(value2).Interface().(uint16)
		return v1 > v2
	case reflect.Uint32:
		v1 := reflect.ValueOf(value1).Interface().(uint32)
		v2 := reflect.ValueOf(value2).Interface().(uint32)
		return v1 > v2
	case reflect.Uint64:
		v1 := reflect.ValueOf(value1).Interface().(uint64)
		v2 := reflect.ValueOf(value2).Interface().(uint64)
		return v1 > v2
	case reflect.Float32:
		v1 := reflect.ValueOf(value1).Interface().(float32)
		v2 := reflect.ValueOf(value2).Interface().(float32)
		return v1 > v2
	case reflect.Float64:
		v1 := reflect.ValueOf(value1).Interface().(float64)
		v2 := reflect.ValueOf(value2).Interface().(float64)
		return v1 > v2
	case reflect.String:
		v1 := reflect.ValueOf(value1).Interface().(string)
		v2 := reflect.ValueOf(value2).Interface().(string)
		return v1 > v2
	case reflect.Bool:
		v1 := reflect.ValueOf(value1).Interface().(bool)
		return v1
	default:
		return false
	}
}
