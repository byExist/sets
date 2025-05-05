package sets_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/byExist/sets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	s := sets.New[int]()
	require.Equal(t, 0, sets.Len(s))
}

func ExampleNew() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Add(s, 3)
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2 3]
}

func TestCollect(t *testing.T) {
	input := []int{1, 2, 3, 4}
	s := sets.Collect(slices.Values(input))
	for _, v := range input {
		assert.True(t, sets.Contains(s, v))
	}
	assert.Equal(t, len(input), sets.Len(s))
}

func ExampleCollect() {
	input := []int{1, 2, 3, 4}
	s := sets.Collect(slices.Values(input))
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2 3 4]
}

func TestAdd(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	require.True(t, sets.Contains(s, 1))
	assert.Equal(t, 1, sets.Len(s))
}

func ExampleAdd() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2]
}

func TestRemove(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Remove(s, 1)
	assert.False(t, sets.Contains(s, 1))
	assert.Equal(t, 0, sets.Len(s))
}

func ExampleRemove() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Remove(s, 1)
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [2]
}

func TestContains(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	assert.True(t, sets.Contains(s, 1))
	assert.False(t, sets.Contains(s, 2))
}

func ExampleContains() {
	s := sets.New[int]()
	sets.Add(s, 1)
	fmt.Println(sets.Contains(s, 1))
	fmt.Println(sets.Contains(s, 2))

	// Output:
	// true
	// false
}

func TestPop(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 42)
	v, ok := sets.Pop(s)
	require.True(t, ok)
	assert.Equal(t, 42, v)
	_, ok = sets.Pop(s)
	assert.False(t, ok)
}

func ExamplePop() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	fmt.Println(sets.Len(s))
	sets.Pop(s)
	fmt.Println(sets.Len(s))

	// Output:
	// 2
	// 1
}

func TestClear(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Clear(s)
	assert.Equal(t, 0, sets.Len(s))
}

func ExampleClear() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Clear(s)
	fmt.Println(sets.Len(s))

	// Output:
	// 0
}

func TestLen(t *testing.T) {
	s := sets.New[int]()
	assert.Equal(t, 0, sets.Len(s))
	sets.Add(s, 1)
	assert.Equal(t, 1, sets.Len(s))
	sets.Add(s, 2)
	assert.Equal(t, 2, sets.Len(s))
}

func ExampleLen() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	fmt.Println(sets.Len(s))

	// Output:
	// 2
}

func TestValues(t *testing.T) {
	input := []int{1, 2, 3}
	s := sets.New[int]()
	for _, v := range input {
		sets.Add(s, v)
	}
	var output []int
	for e := range sets.Values(s) {
		output = append(output, e)
	}
	// 원소 순서는 다를 수 있으니 정렬 후 비교
	assert.ElementsMatch(t, input, output)
}

func ExampleValues() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2]
}

func TestClone(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)

	clone := sets.Clone(s)

	// clone은 원본과 동일한 값을 가져야 함
	assert.True(t, sets.Equal(s, clone))

	// clone과 원본은 메모리상 다른 객체여야 함
	sets.Add(clone, 3)
	assert.False(t, sets.Equal(s, clone))
	assert.Equal(t, 2, sets.Len(s))
	assert.Equal(t, 3, sets.Len(clone))
}

func ExampleClone() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)

	clone := sets.Clone(s)
	elems := slices.Collect(sets.Values(clone))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2]
}

func TestEqual(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	assert.True(t, sets.Equal(a, b))

	sets.Add(a, 1)
	assert.False(t, sets.Equal(a, b))

	sets.Add(b, 1)
	assert.True(t, sets.Equal(a, b))
}

func ExampleEqual() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s2, 1)
	fmt.Println(sets.Equal(s1, s2))

	sets.Add(s2, 2)
	fmt.Println(sets.Equal(s1, s2))

	// Output:
	// true
	// false
}

func TestIsDisjoint(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(a, 1)
	sets.Add(b, 2)
	assert.True(t, sets.IsDisjoint(a, b))

	sets.Add(b, 1)
	assert.False(t, sets.IsDisjoint(a, b))
}

func ExampleIsDisjoint() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s2, 2)
	fmt.Println(sets.IsDisjoint(s1, s2))

	sets.Add(s2, 1)
	fmt.Println(sets.IsDisjoint(s1, s2))

	// Output:
	// true
	// false
}

func TestIsSubset(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(b, 1)
	sets.Add(b, 2)

	assert.True(t, sets.IsSubset(a, b))

	sets.Add(a, 1)
	assert.True(t, sets.IsSubset(a, b))

	sets.Add(a, 3)
	assert.False(t, sets.IsSubset(a, b))
}

func ExampleIsSubset() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s2, 1)
	sets.Add(s2, 2)
	fmt.Println(sets.IsSubset(s1, s2))

	sets.Add(s1, 1)
	fmt.Println(sets.IsSubset(s1, s2))

	sets.Add(s1, 3)
	fmt.Println(sets.IsSubset(s1, s2))

	// Output:
	// true
	// true
	// false
}

func TestIsSuperset(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(a, 1)
	sets.Add(a, 2)

	assert.True(t, sets.IsSuperset(a, b))

	sets.Add(b, 1)
	assert.True(t, sets.IsSuperset(a, b))

	sets.Add(b, 3)
	assert.False(t, sets.IsSuperset(a, b))
}

func ExampleIsSuperset() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s1, 2)
	fmt.Println(sets.IsSuperset(s1, s2))

	sets.Add(s2, 1)
	fmt.Println(sets.IsSuperset(s1, s2))

	sets.Add(s2, 3)
	fmt.Println(sets.IsSuperset(s1, s2))

	// Output:
	// true
	// true
	// false
}

func TestUnion(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(a, 1)
	sets.Add(b, 2)
	result := sets.Union(a, b)
	assert.True(t, sets.Contains(result, 1))
	assert.True(t, sets.Contains(result, 2))
	assert.Equal(t, 2, sets.Len(result))
}

func ExampleUnion() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s2, 2)
	union := sets.Union(s1, s2)
	elems := slices.Collect(sets.Values(union))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2]
}

func TestIntersection(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(a, 1)
	sets.Add(a, 2)
	sets.Add(b, 2)
	sets.Add(b, 3)
	result := sets.Intersection(a, b)
	assert.True(t, sets.Contains(result, 2))
	assert.Equal(t, 1, sets.Len(result))
}

func ExampleIntersection() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s1, 2)
	sets.Add(s2, 2)
	sets.Add(s2, 3)
	intersection := sets.Intersection(s1, s2)
	elems := slices.Collect(sets.Values(intersection))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [2]
}

func TestDifference(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(a, 1)
	sets.Add(a, 2)
	sets.Add(b, 2)
	sets.Add(b, 3)
	result := sets.Difference(a, b)
	assert.True(t, sets.Contains(result, 1))
	assert.False(t, sets.Contains(result, 2))
	assert.Equal(t, 1, sets.Len(result))
}

func ExampleDifference() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s1, 2)
	sets.Add(s2, 2)
	sets.Add(s2, 3)
	difference := sets.Difference(s1, s2)
	elems := slices.Collect(sets.Values(difference))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1]
}

func TestSymmetricDifference(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	sets.Add(a, 1)
	sets.Add(a, 2)
	sets.Add(b, 2)
	sets.Add(b, 3)
	result := sets.SymmetricDifference(a, b)
	assert.True(t, sets.Contains(result, 1))
	assert.True(t, sets.Contains(result, 3))
	assert.False(t, sets.Contains(result, 2))
	assert.Equal(t, 2, sets.Len(result))
}

func ExampleSymmetricDifference() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s1, 2)
	sets.Add(s2, 2)
	sets.Add(s2, 3)
	symDiff := sets.SymmetricDifference(s1, s2)
	elems := slices.Collect(sets.Values(symDiff))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 3]
}
