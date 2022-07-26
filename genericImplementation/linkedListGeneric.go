package appGeneric

type Node[T comparable] struct {
	next  *Node[T]
	Value T
}
type List[T comparable] struct {
	head   *Node[T]
	Length int
}

func NewLinkedList[T comparable]() *List[T] {
	return &List[T]{}
}

func (s *List[T]) Add(value T) {
	el := &Node[T]{
		Value: value,
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

func (s *List[T]) Delete(value T) {
	if s.head == nil {
		return
	}
	var prev *Node[T]
	current := s.head
	if current == s.head && s.Length == 1 && current.Value == value {
		s.head = nil
		s.Length--
		return
	}
	for current.next != nil {
		if current.Value == value {
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

func (s *List[T]) Find(value T) *Node[T] {
	if s.head == nil {
		return nil
	}
	var foundNode *Node[T]
	current := s.head
	if current == s.head && s.Length == 1 && current.Value == value {
		return current
	}
	for current.next != nil {
		if current.Value == value {
			foundNode = current
			break
		}
		current = current.next
	}
	return foundNode
}

func (s *List[T]) IsCycleInList() bool {
	if s.head == nil {
		return false
	}
	haveCycle := false
	current := s.head
	setOfNodes := make(map[*Node[T]]bool)
	setOfNodes[s.head] = true
	for current.next != nil {
		if setOfNodes[current.next] {
			return true
		}
		current = current.next
	}
	return haveCycle
}
