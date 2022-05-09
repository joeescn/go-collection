package collection

type Set[Element comparable] interface {
	// Set the value for a key.
	Add(Element) bool

	// Adds the given elements to the set.
	Adds(...Element)

	// Size return the number of keys.
	Size() int

	// Clone returns a clone of the set.
	Clone() Set[Element]

	// Contains returns whether the given items
	// are all in the set.
	Contains(elements ...Element) bool

	// Pop removes and returns an arbitrary item from the set.
	Pop() (Element, bool)

	// Elements returns a slice of the elements.
	Elements() []Element

	// Difference returns the difference between this set
	// and other. The returned set will contain
	// all elements of this set that are not also
	// elements of other.
	//
	// 差集 返回一个集合，元素包含在集合 s ，但不在集合 other
	Difference(other Set[Element]) Set[Element]

	// Intersection returns a new set containing only the elements
	// that exist only in both sets.
	//
	// 交集 返回一个集合，元素包含在集合 s 和 other
	Intersection(other Set[Element]) Set[Element]

	// Determines if every element in this set is in
	// the other set but the two sets are not equal.
	//
	// 确定 set 中的所有元素是否都在 other set 中
	IsProperSubset(other Set[Element]) bool

	// Determines if every element in the other set
	// is in this set but the two sets are not
	// equal.
	//
	// 确定 other 集合中的元素都在这个集合中
	IsProperSuperset(other Set[Element]) bool

	// Determines if every element in this set is in
	// the other set.
	//
	// 确定 this set 是否是 other set 的子集
	IsSubset(other Set[Element]) bool

	// Determines if every element in the other set
	// is in this set.
	//
	// 判断 other 是否是 this 的超集
	IsSuperset(other Set[Element]) bool

	// Union returns a new set with all elements from
	// this set and other.
	//
	// 并集 返回一个集合，元素包含在集合 s 和 other
	Union(other Set[Element]) Set[Element]

	// Range iterates over the set.
	Range(fn func(Element) bool)
}

// NewSet returns a new empty set.
// 	map is thread unsafe
func NewSet[Element comparable](elems ...Element) Set[Element] {
	s := &mapset[Element]{m: NewMap[Element, struct{}]()}
	s.Adds(elems...)
	return s
}

// NewSyncSet returns a new empty set.
// 	map is thread safe
func NewSyncSet[Element comparable](elems ...Element) Set[Element] {
	s := &mapset[Element]{m: NewSyncMap[Element, struct{}]()}
	s.Adds(elems...)
	return s

}
