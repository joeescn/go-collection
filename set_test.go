package collection

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[string]()
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")

	other := NewSet[string]()
	other.Adds("google", "microsoft", "apple")

	diff := s.Difference(other)

	ret := NewSet("cherry", "banana")

	isProperSubset := diff.IsProperSuperset(ret)
	if !isProperSubset {
		t.Fatal("不同集合", diff.Elements())
	}

	diff.Add("mango")
	if diff.IsProperSubset(ret) {
		t.Fatal("不同集合", diff.Elements())
	}

}
