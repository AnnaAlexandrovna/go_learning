package appInterface

type Node struct {
	next  *Node
	Value interface{}
}
type List struct {
	head   *Node
	Length int
}

func NewLinkedList() *List {
	return &List{}
}

func (s *List) Add(value interface{}) {
	el := &Node{
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

func (s *List) Delete(value interface{}) {
	if s.head == nil {
		return
	}
	var prev *Node
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

func (s *List) Find(value interface{}) *Node {
	if s.head == nil {
		return nil
	}
	var foundNode *Node
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
