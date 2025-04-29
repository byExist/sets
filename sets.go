package sets

import (
	"iter"
)

var empty = struct{}{}

// Set represents a generic Set of comparable elements.
type Set[E comparable] struct {
	data map[E]struct{}
}

// New creates and returns a new empty set.
func New[E comparable]() *Set[E] {
	return &Set[E]{data: make(map[E]struct{})}
}

// Collect creates a set from an iterator sequence.
func Collect[E comparable](i iter.Seq[E]) *Set[E] {
	s := New[E]()
	for e := range i {
		s.data[e] = empty
	}
	return s
}

// Clone returns a deep copy of the given set.
func Clone[E comparable](s *Set[E]) *Set[E] {
	result := New[E]()
	for e := range s.data {
		result.data[e] = empty
	}
	return result
}

// Len returns the number of elements in the set.
func Len[E comparable](s *Set[E]) int {
	return len(s.data)
}

// Contains checks if the set contains the given element.
func Contains[E comparable](s *Set[E], e E) bool {
	if _, ok := s.data[e]; ok {
		return true
	}
	return false
}

// Add inserts an element into the set.
func Add[E comparable](s *Set[E], e E) {
	s.data[e] = empty
}

// Remove deletes an element from the set.
func Remove[E comparable](s *Set[E], e E) {
	delete(s.data, e)
}

// Pop removes and returns an arbitrary element from the set.
func Pop[E comparable](s *Set[E]) (E, bool) {
	for e := range s.data {
		delete(s.data, e)
		return e, true
	}
	var zero E
	return zero, false
}

// Clear removes all elements from the set.
func Clear[E comparable](s *Set[E]) {
	s.data = make(map[E]struct{})
}

// Values returns an iterator over all elements in the set.
func Values[E comparable](s *Set[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range s.data {
			if !yield(e) {
				break
			}
		}
	}
}

// Equal checks if two sets contain exactly the same elements.
func Equal[E comparable](a, b *Set[E]) bool {
	if len(a.data) != len(b.data) {
		return false
	}
	for e := range a.data {
		if _, ok := b.data[e]; !ok {
			return false
		}
	}
	return true
}

// IsDisjoint checks if two sets have no elements in common.
func IsDisjoint[E comparable](a, b *Set[E]) bool {
	for e := range a.data {
		if _, ok := b.data[e]; ok {
			return false
		}
	}
	return true
}

// IsSubset checks if the first set is a subset of the second set.
func IsSubset[E comparable](a, b *Set[E]) bool {
	if len(a.data) > len(b.data) {
		return false
	}
	for e := range a.data {
		if _, ok := b.data[e]; !ok {
			return false
		}
	}
	return true
}

// IsSuperset checks if the first set is a superset of the second set.
func IsSuperset[E comparable](a, b *Set[E]) bool {
	if len(a.data) < len(b.data) {
		return false
	}
	for e := range b.data {
		if _, ok := a.data[e]; !ok {
			return false
		}
	}
	return true
}

// Union returns a new set containing all elements from both sets.
func Union[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		result.data[e] = empty
	}
	for e := range b.data {
		result.data[e] = empty
	}
	return result
}

// Intersection returns a new set containing only elements present in both sets.
func Intersection[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		if _, ok := b.data[e]; ok {
			result.data[e] = empty
		}
	}
	return result
}

// Difference returns a new set containing elements in the first set but not in the second.
func Difference[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		if _, ok := b.data[e]; !ok {
			result.data[e] = empty
		}
	}
	return result
}

// SymmetricDifference returns a new set containing elements present in either set but not in both.
func SymmetricDifference[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		if _, ok := b.data[e]; !ok {
			result.data[e] = empty
		}
	}
	for e := range b.data {
		if _, ok := a.data[e]; !ok {
			result.data[e] = empty
		}
	}
	return result
}
