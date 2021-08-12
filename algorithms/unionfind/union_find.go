package unionfind

type Set struct {
	items []int
	size  []int
}

func (s *Set) Union1(u, v int) {
	uRoot := s.Find(u)
	vRoot := s.Find(v)

	s.items[vRoot] = uRoot
}

func (s *Set) Union(u, v int) {
	uRoot := s.Find(u)
	vRoot := s.Find(v)

	if uRoot == vRoot {
		return
	}

	if s.size[vRoot] < s.size[uRoot] {
		s.items[vRoot] = s.items[uRoot]
		s.size[uRoot] += s.size[vRoot]
	} else {
		s.items[uRoot] = s.items[vRoot]
		s.size[vRoot] += s.size[uRoot]
	}
}

func (s *Set) Find(i int) int {
	for s.items[i] != i {
		i = s.items[i]
	}

	return i
}

func (s *Set) Items() []int {
	return s.items
}

func New(i int) *Set {
	items := make([]int, i)
	for j := 0; j < i; j++ {
		items[j] = j
	}
	return &Set{
		items: items,
		size:  make([]int, i),
	}
}
