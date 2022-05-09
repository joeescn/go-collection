package collection

type mapset[Element comparable] struct {
	m Map[Element, struct{}]
}

// Set the value for a key.
func (s *mapset[Element]) Add(elem Element) bool {
	_, found := s.m.Get(elem)
	s.m.Set(elem, struct{}{})
	return !found
}

func (s *mapset[Element]) Adds(elems ...Element) {
	for _, e := range elems {
		s.m.Set(e, struct{}{})
	}
}

// Size return the number of keys.
func (s *mapset[Element]) Size() int {
	return s.m.Size()
}

// Clone returns a clone of the set.
func (s *mapset[Element]) Clone() Set[Element] {
	return &mapset[Element]{m: s.m.Clone()}
}

// Contains returns whether the given items
// are all in the set.
func (s *mapset[Element]) Contains(elements ...Element) bool {
	for _, elem := range elements {
		if !s.m.Has(elem) {
			return false
		}
	}
	return true
}

// Pop removes and returns an arbitrary item from the set.
func (s *mapset[Element]) Pop() (elem Element, ok bool) {
	s.m.Range(func(e Element, _ struct{}) bool {
		elem = e
		ok = true
		return true
	})
	return
}

// Elements returns a slice of the elements.
func (s *mapset[Element]) Elements() []Element {
	return s.m.Keys()
}

// Difference returns the difference between this set
// and other. The returned set will contain
// all elements of this set that are not also
// elements of other.
//
// 差集 返回一个集合，元素包含在集合 s ，但不在集合 other
func (s *mapset[Element]) Difference(other Set[Element]) Set[Element] {
	diff := &mapset[Element]{m: s.m.Clone()}
	other.Range(func(elem Element) bool {
		diff.m.Delete(elem)
		return false
	})
	return diff
}

// Intersection returns a new set containing only the elements
// that exist only in both sets.
//
// 交集 返回一个新集合，该集合的元素既包含在集合 s 又包含在集合 other 中
func (s *mapset[Element]) Intersection(other Set[Element]) Set[Element] {
	inter := &mapset[Element]{m: s.m.Clone()}
	inter.m.Range(func(elem Element, _ struct{}) bool {
		if !other.Contains(elem) {
			inter.m.Delete(elem)
		}
		return false
	})
	return inter
}

// Determines if every element in this set is in
// the other set but the two sets are not equal.
//
// 确定 set 中的所有元素是否都在 other set 中
func (s *mapset[Element]) IsProperSubset(other Set[Element]) bool {
	return s.Size() <= other.Size() && s.Difference(other).Size() == 0
}

// Determines if every element in the other set
// is in this set but the two sets are not
// equal.
//
// 确定 other 集合中的元素都在这个集合中
func (s *mapset[Element]) IsProperSuperset(other Set[Element]) bool {
	return other.IsProperSubset(s)
}

// Determines if every element in this set is in
// the other set.
//
// 确定 this set 是否是 other set 的子集
func (s *mapset[Element]) IsSubset(other Set[Element]) bool {
	return s.Size() <= other.Size() && s.Difference(other).Size() == 0
}

// Determines if every element in the other set
// is in this set.
//
// 判断 other 是否是 this 的超集
func (s *mapset[Element]) IsSuperset(other Set[Element]) bool {
	return other.IsSubset(s)
}

// Union returns a new set with all elements from
// this set and other.
//
// 并集 返回一个新集合，该集合的元素包含在集合 s 和 other 中
func (s *mapset[Element]) Union(other Set[Element]) Set[Element] {
	union := &mapset[Element]{m: s.m.Clone()}
	other.Range(func(elem Element) bool {
		union.m.Set(elem, struct{}{})
		return false
	})
	return union
}

// Range iterates over the set.
func (s *mapset[Element]) Range(fn func(Element) bool) {
	s.m.Range(func(e Element, _ struct{}) bool {
		return fn(e)
	})
}
