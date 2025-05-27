# sets [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/sets.svg)](https://pkg.go.dev/github.com/byExist/sets) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/sets)](https://goreportcard.com/report/github.com/byExist/sets)

A minimal, Python-like set utility for Go.

## ✨ Features

- ✅ Type-safe and generic set based on Go's `map`
- ✅ Constant-time: `Add`, `Remove`, `Contains`, `Len`, `Clear`
- ✅ Supports union, intersection, difference, symmetric difference
- ✅ Provides iterators and JSON (un)marshaling
- ❌ Not concurrency-safe
- ❌ Only for `comparable` types

## 🧱 Example

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

## 🔍 API

| Function/Method | Description |
|------------------|-------------|
| `New()` | Create a new empty set |
| `Add(set, elem)` | Add an element |
| `Remove(set, elem)` | Remove an element |
| `Contains(set, elem)` | Check membership |
| `Len(set)` | Number of elements |
| `Clear(set)` | Remove all elements |
| `Values(set)` | Get iterator |
| `Union(a, b)` | Union of two sets |
| `Intersection(a, b)` | Intersection of two sets |
| `Difference(a, b)` | Difference of two sets |
| `SymmetricDifference(a, b)` | Symmetric difference |
| `Equal(a, b)` | Check equality |
| `IsSubset(a, b)` | Check subset |
| `IsSuperset(a, b)` | Check superset |
| `IsDisjoint(a, b)` | Check disjointness |
| `Clone(set)` | Copy set |

## ⚠️ Limitations

- Not safe for concurrent access  
  Use a sync.Mutex if multiple goroutines will access the set.
- Only works with types that are `comparable` in Go  
  (e.g., slices and maps cannot be used as set elements)

## 🪪 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
