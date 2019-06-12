package pkg

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i <= 9 {
		s.data[s.i] = k
		s.i++
	}
}

func (s *stack) pop() int {
	if s.i > 0 {
		defer func() { s.i-- }()
		return s.data[s.i]
	}
	return 0
}
