package appInterface

type Node struct {
	next *Node
	Key  interface{}
}
type List struct {
	head   *Node
	Length int
}

func NewLinkedList() *List {
	return &List{}
}

func (s *List) Add(key interface{}) {
	el := &Node{
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

func (s *List) Delete(key interface{}) {
	if s.head == nil {
		return
	}
	var prev *Node
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

func (s *List) Find(key interface{}) *Node {
	if s.head == nil {
		return nil
	}
	var foundNode *Node
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
