package appGeneric

type Node[T comparable] struct {
	next *Node[T]
	Key  T
}
type List[T comparable] struct {
	head   *Node[T]
	Length int
}

func InitList[T comparable]() *List[T] {
	return &List[T]{}
}

func (s *List[T]) Add(key T) {
	el := &Node[T]{
		Key: key,
	}
	if s.head == nil {
		s.head = el
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = el
	}
	s.Length++
}

func (s *List[T]) Delete(key T) {
	if s.head == nil {
		return
	}
	var prev *Node[T]
	current := s.head
	if current == s.head && s.Length == 1 && current.Key == key {
		s.head = nil
		s.Length--
		return
	}
	for current.next != nil {
		if current.Key == key {
			if current == s.head {
				s.head = current.next
			} else {
				prev.next = current.next
			}
			s.Length--
			return
		}
		prev = current
		current = current.next
	}
}

func (s *List[T]) Find(key T) *Node[T] {
	if s.head == nil {
		return nil
	}
	var foundNode *Node[T]
	current := s.head
	if current == s.head && s.Length == 1 && current.Key == key {
		return current
	}
	for current.next != nil {
		if current.Key == key {
			foundNode = current
			break
		}
		current = current.next
	}
	return foundNode
}

// func compare[T comparable](key1 T, key2 T) bool {
// 	r := reflect.TypeOf(key1).Kind()
// 	if r == reflect.Int ||
// 		r == reflect.Int8 ||
// 		r == reflect.Int16 ||
// 		r == reflect.Int32 ||
// 		r == reflect.Int64 ||
// 		r == reflect.Float32 ||
// 		r == reflect.Float64 ||
// 		r == reflect.String ||
// 		r == reflect.Bool ||
// 		r == reflect.Uint ||
// 		r == reflect.Uint8 ||
// 		r == reflect.Uint16 ||
// 		r == reflect.Uint32 ||
// 		r == reflect.Uint64 {
// 		v1 := reflect.ValueOf(key1).Interface()
// 		v2 := reflect.ValueOf(key2).Interface()
// 		fmt.Println("v1", v1, "v2", v2, v1 == v2)
// 		return v1 == v2
// 	}
// 	return false
// }
