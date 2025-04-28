# sets [![GoDoc](https://pkg.go.dev/badge/github.com/yourname/sets.svg)](https://pkg.go.dev/github.com/yourname/sets) [![Go Report Card](https://goreportcard.com/badge/github.com/yourname/sets)](https://goreportcard.com/report/github.com/yourname/sets)

## What is "sets"?

**sets** is a lightweight and generic Go package that provides an easy-to-use interface for managing sets of comparable elements.  
It supports common set operations like union, intersection, difference, and symmetric difference, along with basic utilities like adding, removing, cloning, and iterating elements.

## Features

- Create new sets easily
- Add, remove, and check for elements
- Clone sets
- Iterate over set elements
- Perform standard set operations (Union, Intersection, Difference, SymmetricDifference)
- Check set relationships (Equal, Subset, Superset, Disjoint)

## Installation

To install sets, use the following command:

```bash
go get github.com/yourname/sets
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/yourname/sets"
)

func main() {
	// Create a new set
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Add(s, 3)

	// Iterate over set elements
	for e := range sets.Values(s) {
		fmt.Println(e)
	}

	// Check if an element exists
	fmt.Println("Contains 2?", sets.Contains(s, 2))

	// Remove an element
	sets.Remove(s, 2)

	// Set operations
	t := sets.New[int]()
	sets.Add(t, 3)
	sets.Add(t, 4)

	union := sets.Union(s, t)
	for e := range sets.Values(union) {
		fmt.Println("Union element:", e)
	}
}
```

## Usage

The sets package provides:

- Set creation and basic operations (add, remove, check elements)
- Iterating elements
- Standard set operations (union, intersection, difference, symmetric difference)
- Checking set relationships (subset, superset, disjoint, equal)

Refer to the API Overview below for detailed usage.

## API Overview

### Constructors

- `New[E comparable]() *set[E]`
- `Collect[E comparable](i iter.Seq[E]) *set[E]`

### Basic Operations

- `Add(s *set[E], e E)`
- `Remove(s *set[E], e E)`
- `Contains(s *set[E], e E) bool`
- `Pop(s *set[E]) (E, bool)`
- `Clear(s *set[E])`
- `Len(s *set[E]) int`
- `Values(s *set[E]) iter.Seq[E]`
- `Clone(s *set[E]) *set[E]`

### Set Relations

- `Equal(a, b *set[E]) bool`
- `IsDisjoint(a, b *set[E]) bool`
- `IsSubset(a, b *set[E]) bool`
- `IsSuperset(a, b *set[E]) bool`

### Set Operations

- `Union(a, b *set[E]) *set[E]`
- `Intersection(a, b *set[E]) *set[E]`
- `Difference(a, b *set[E]) *set[E]`
- `SymmetricDifference(a, b *set[E]) *set[E]`

## Notes

- The sets package **is not concurrency-safe**.  
  If you need to access a set from multiple goroutines, you must synchronize access externally.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
