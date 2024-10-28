package ds

type Stack[t any] struct {
	elements []t
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// pop return 2 values first bool value to determine whether stack is empty or not second value represents the value poped if stack is empty the second value will be zero value of type t/
func (s *Stack[t]) Pop() (bool, t) {
	if len(s.elements) == 0 {
		var zero t
		return true, zero
	}
	popedEle := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return false, popedEle
}
