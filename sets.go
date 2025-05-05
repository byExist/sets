package sets

import (
	"encoding/json"
	"fmt"
	"iter"
	"slices"
	"strings"
)

var exists = struct{}{}

// Set represents a generic Set of comparable elements.
type Set[E comparable] struct {
	data map[E]struct{}
}

// String implements fmt.Stringer for Set[E].
func (s *Set[E]) String() string {
	elems := slices.Collect(Values(s))
	var b strings.Builder
	b.WriteString("Set{")
	for i, e := range elems {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprint(e))
	}
	b.WriteString("}")
	return b.String()
}

// MarshalJSON implements json.Marshaler for Set[E].
func (s *Set[E]) MarshalJSON() ([]byte, error) {
	elems := slices.Collect(Values(s))
	return json.Marshal(elems)
}

// UnmarshalJSON implements json.Unmarshaler for Set[E].
func (s *Set[E]) UnmarshalJSON(data []byte) error {
	var elems []E
	if err := json.Unmarshal(data, &elems); err != nil {
		return err
	}
	s.data = make(map[E]struct{}, len(elems))
	for _, e := range elems {
		s.data[e] = exists
	}
	return nil
}

// New creates and returns a new empty set.
func New[E comparable]() *Set[E] {
	return &Set[E]{data: make(map[E]struct{})}
}

// Collect creates a set from an iterator sequence.
func Collect[E comparable](i iter.Seq[E]) *Set[E] {
	s := New[E]()
	for e := range i {
		s.data[e] = exists
	}
	return s
}

// Add inserts an element into the set.
func Add[E comparable](s *Set[E], e E) {
	s.data[e] = exists
}

// Remove deletes an element from the set.
func Remove[E comparable](s *Set[E], e E) {
	delete(s.data, e)
}

// Contains checks if the set contains the given element.
func Contains[E comparable](s *Set[E], e E) bool {
	_, ok := s.data[e]
	return ok
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

// Len returns the number of elements in the set.
func Len[E comparable](s *Set[E]) int {
	return len(s.data)
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

// Clone returns a deep copy of the given set.
func Clone[E comparable](s *Set[E]) *Set[E] {
	result := New[E]()
	for e := range s.data {
		result.data[e] = exists
	}
	return result
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
	if len(a.data) > len(b.data) {
		a, b = b, a
	}
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

// Union (A ∪ B): returns a new set containing all elements from both sets.
func Union[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		result.data[e] = exists
	}
	for e := range b.data {
		result.data[e] = exists
	}
	return result
}

// Intersection (A ∩ B): returns a new set containing only elements present in both sets.
func Intersection[E comparable](a, b *Set[E]) *Set[E] {
	if len(a.data) > len(b.data) {
		a, b = b, a
	}
	result := New[E]()
	for e := range a.data {
		if _, ok := b.data[e]; ok {
			result.data[e] = exists
		}
	}
	return result
}

// Difference (A − B): returns a new set containing elements in the first set but not in the second.
func Difference[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		if _, ok := b.data[e]; !ok {
			result.data[e] = exists
		}
	}
	return result
}

// SymmetricDifference (A △ B): returns a new set containing elements present in either set but not in both.
func SymmetricDifference[E comparable](a, b *Set[E]) *Set[E] {
	result := New[E]()
	for e := range a.data {
		if _, ok := b.data[e]; !ok {
			result.data[e] = exists
		}
	}
	for e := range b.data {
		if _, ok := a.data[e]; !ok {
			result.data[e] = exists
		}
	}
	return result
}
