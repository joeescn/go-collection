# go-collection
The missing `generic` set collection for the Go language.  Until Go has sets built-in...use this.

## example
```golang
package main

import (
    "fmt"
    collection "github.com/deckarep/go-collection"
)

func main() {
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
        	fmt.Println(("不同集合", diff.Elements()))
	}

	diff.Add("mango")
	if diff.IsProperSubset(ret) {
		fmt.Println(("不同集合", diff.Elements()))
	}
}
```
