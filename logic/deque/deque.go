package deque

type Deque[T any] struct {
	sl []T
}

func (d *Deque[T]) GetSlice() []T {
	// for tests
	return d.sl
}
func New[T any]() Deque[T] {
	return Deque[T]{make([]T, 0)}
}

func (s *Deque[T]) Push_back(values ...T) {
	s.sl = append(s.sl, values...)
}

func (s *Deque[T]) Push_front(values ...T) {
	s.sl = append(values, s.sl...)
}
func (s *Deque[T]) Pop_back() T {
	last := s.sl[len(s.sl)-1]
	s.sl = s.sl[:len(s.sl)-1]
	return last
}
func (s *Deque[T]) Pop_front() T {
	first := s.sl[0]
	if s.Len() > 1 {
		s.sl = s.sl[1:]
	} else {
		s.sl = make([]T, 0)
	}
	return first
}

func (s *Deque[T]) Len() int {
	return len(s.sl)
}

func (s *Deque[T]) Index(i int) T {
	return s.sl[i]
}

func (s *Deque[T]) Pop_back_slice(n int) []T {
	if n > s.Len() {
		res := s.sl
		s.sl = nil
		return res
	} else {
		res := s.sl[len(s.sl)-n:]
		s.sl = s.sl[:len(s.sl)-n]
		return res
	}
}

func (d *Deque[T]) Inspect(f func(T)) {
	for _, v := range d.sl {
		f(v)
	}

}
