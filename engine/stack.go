package engine

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() interface{} {
	if s.Size() == 0 {
		return nil
	}
	element := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]
	return element
}

func (s *Stack) Size() int {
	return len(s.elements)
}
