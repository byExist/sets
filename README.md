# sets [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/sets.svg)](https://pkg.go.dev/github.com/byExist/sets) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/sets)](https://goreportcard.com/report/github.com/byExist/sets)

## What is "sets"?

**sets** is a generic, lightweight wrapper around Go's `map` type.
It provides a convenient and type-safe way to perform set operations, such as element management (add/remove), containment checks, and standard set operations like union and intersection.  
It is particularly useful for deduplication, fast membership testing, and working with logically grouped values like tags or permissions.


## Installation

To install sets, use the following command:

```bash
go get github.com/byExist/sets
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/byExist/sets"
)

func main() {
	// Create a new set
	s := sets.New[int]()
	
	// Add elements to the set
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Add(s, 3)

	// Iterate over set elements in sorted order
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	for _, e := range elems {
		fmt.Println(e)
	}

	// Check for membership
	fmt.Println("Contains 2?", sets.Contains(s, 2))

	// Remove an element
	sets.Remove(s, 2)

	// Set operations
	t := sets.New[int]()
	sets.Add(t, 3)
	sets.Add(t, 4)

	union := sets.Union(s, t)
	unionElems := slices.Collect(sets.Values(union))
	slices.Sort(unionElems)
	for _, e := range unionElems {
		fmt.Println("Union element:", e)
	}
}
```

**Output:**
```
1
2
3
Contains 2? true
Union element: 1
Union element: 3
Union element: 4
```

## API Overview

### Constructors

| Function                      | Description                   | Time Complexity |
|-------------------------------|------------------------------|-----------------|
| `New[E comparable]()`          | Create a new empty set        | O(1)            |
| `Collect[E comparable](i iter.Seq[E])` | Create a set from an iterator | O(n)            |

### Basic Operations

| Function                       | Description                       | Time Complexity |
|--------------------------------|---------------------------------|-----------------|
| `Add(s *Set[E], e E)`           | Add an element to the set        | O(1)            |
| `Remove(s *Set[E], e E)`        | Remove an element from the set   | O(1)            |
| `Contains(s *Set[E], e E)`      | Check if an element exists in the set | O(1)        |
| `Pop(s *Set[E]) (E, bool)`      | Remove and return an arbitrary element | O(1)       |
| `Clear(s *Set[E])`              | Remove all elements from the set | O(1)            |
| `Len(s *Set[E]) int`            | Get the number of elements in the set | O(1)        |
| `Values(s *Set[E]) iter.Seq[E]` | Get an iterator over the set elements | O(n)        |
| `Clone(s *Set[E]) *Set[E]`      | Create a copy of the set          | O(n)            |

### Set Relations

| Function                         | Description                          | Time Complexity |
|----------------------------------|------------------------------------|-----------------|
| `Equal(a, b *Set[E])`             | Check if two sets are equal         | O(n)            |
| `IsDisjoint(a, b *Set[E])`        | Check if two sets have no elements in common | O(n)      |
| `IsSubset(a, b *Set[E])`          | Check if set a is a subset of set b | O(n)            |
| `IsSuperset(a, b *Set[E])`        | Check if set a is a superset of set b | O(n)          |

### Set Operations

| Function                             | Description                         | Time Complexity |
|-------------------------------------|-----------------------------------|-----------------|
| `Union(a, b *Set[E]) *Set[E]`       | Return the union of two sets        | O(n)            |
| `Intersection(a, b *Set[E]) *Set[E]`| Return the intersection of two sets | O(n)            |
| `Difference(a, b *Set[E]) *Set[E]`  | Return the difference of two sets  | O(n)            |
| `SymmetricDifference(a, b *Set[E]) *Set[E]` | Return the symmetric difference of two sets | O(n) |

## Limitations

- Not safe for concurrent access  
  Use a sync.Mutex if multiple goroutines will access the set.
- Only works with types that are `comparable` in Go  
  (e.g., slices and maps cannot be used as set elements)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
