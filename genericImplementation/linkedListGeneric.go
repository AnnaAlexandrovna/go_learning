package appGeneric

type Node[T comparable] struct {
	Next  *Node[T]
	Value T
}
type List[T comparable] struct {
	Head   *Node[T]
	Length int
}

func NewLinkedList[T comparable]() *List[T] {
	return &List[T]{}
}

func (s *List[T]) Add(value T) {
	el := &Node[T]{
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

func (s *List[T]) Delete(value T) {
	if s.Head == nil {
		return
	}
	var prev *Node[T]
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

func (s *List[T]) Find(value T) *Node[T] {
	if s.Head == nil {
		return nil
	}
	var foundNode *Node[T]
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

func (s *List[T]) IsCycleInList() bool {
	if s.Head == nil {
		return false
	}
	haveCycle := false
	current := s.Head
	setOfNodes := make(map[*Node[T]]bool)
	setOfNodes[s.Head] = true
	for current.Next != nil {
		if setOfNodes[current.Next] {
			return true
		}
		current = current.Next
	}
	return haveCycle
}

func (s *List[T]) RevertList() {
	isCycleInList := s.IsCycleInList()
	if isCycleInList {
		return
	}
	if s.Head == nil {
		return
	}
	var prev *Node[T]
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
