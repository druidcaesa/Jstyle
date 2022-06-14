package set

// Empty Use an empty structure to save space
type Empty struct{}

// Set Implement a set because set cannot have duplicate values, so use the key of the map to solve this problem
type Set[K comparable] struct {
	m map[K]struct{}
}

// New returns an empty hashset.
func New[K comparable]() Set[K] {
	return Set[K]{
		m: make(map[K]struct{}),
	}
}

// Put Add value to set
func (s Set[K]) Put(val K) {
	s.m[val] = struct{}{}
}

// Exists Judge whether the value exists
func (s Set[K]) Exists(val K) bool {
	_, ok := s.m[val]
	return ok
}

// Remove Delete the value of set
func (s Set[K]) Remove(val K) {
	delete(s.m, val)
}

// Size Get the length of the set.
func (s Set[K]) Size() int {
	return len(s.m)
}

// ForEach Iterate over set.
func (s Set[K]) ForEach(fn func(key K)) {
	for k, _ := range s.m {
		fn(k)
	}
}
