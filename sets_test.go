package sets_test

import (
	"encoding/json"
	"fmt"
	"slices"
	"testing"

	"github.com/byExist/sets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 2)
	sets.Add(s, 1)
	output := s.String()

	assert.Contains(t, output, "Set{")
	assert.Contains(t, output, "1")
	assert.Contains(t, output, "2")
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)

	data, err := json.Marshal(s)
	require.NoError(t, err)

	var decoded sets.Set[int]
	err = json.Unmarshal(data, &decoded)
	require.NoError(t, err)

	assert.True(t, sets.Equal(s, &decoded))
}

func TestNew(t *testing.T) {
	s := sets.New[int]()
	require.Equal(t, 0, sets.Len(s))
}

func TestCollect(t *testing.T) {
	input := []int{1, 2, 3, 4}
	s := sets.Collect(slices.Values(input))
	for _, v := range input {
		assert.True(t, sets.Contains(s, v))
	}
	assert.Equal(t, len(input), sets.Len(s))
}

func TestAdd(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	require.True(t, sets.Contains(s, 1))
	assert.Equal(t, 1, sets.Len(s))
}

func TestRemove(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Remove(s, 1)
	assert.False(t, sets.Contains(s, 1))
	assert.Equal(t, 0, sets.Len(s))
}

func TestContains(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	assert.True(t, sets.Contains(s, 1))
	assert.False(t, sets.Contains(s, 2))
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

func TestClear(t *testing.T) {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	sets.Clear(s)
	assert.Equal(t, 0, sets.Len(s))
}

func TestLen(t *testing.T) {
	s := sets.New[int]()
	assert.Equal(t, 0, sets.Len(s))
	sets.Add(s, 1)
	assert.Equal(t, 1, sets.Len(s))
	sets.Add(s, 2)
	assert.Equal(t, 2, sets.Len(s))
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

func TestEqual(t *testing.T) {
	a := sets.New[int]()
	b := sets.New[int]()
	assert.True(t, sets.Equal(a, b))

	sets.Add(a, 1)
	assert.False(t, sets.Equal(a, b))

	sets.Add(b, 1)
	assert.True(t, sets.Equal(a, b))
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

func ExampleCollect() {
	input := []int{1, 2, 3, 4}
	s := sets.Collect(slices.Values(input))
	elems := slices.Collect(sets.Values(s))
	slices.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2 3 4]
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

func ExampleContains() {
	s := sets.New[int]()
	sets.Add(s, 1)
	fmt.Println(sets.Contains(s, 1))
	fmt.Println(sets.Contains(s, 2))

	// Output:
	// true
	// false
}

func ExamplePop() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	v, ok := sets.Pop(s)
	fmt.Println("Popped:", v, ok)
	fmt.Println("Len after pop:", sets.Len(s))
	// Output:
	// Popped: 1 true
	// Len after pop: 1
}

func ExampleClear() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	fmt.Println("Before Clear:", sets.Len(s))
	sets.Clear(s)
	fmt.Println("After Clear:", sets.Len(s))
	// Output:
	// Before Clear: 2
	// After Clear: 0
}

func ExampleLen() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	fmt.Println(sets.Len(s))

	// Output:
	// 2
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

func ExampleIsDisjoint() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s2, 2)
	fmt.Println("s1 and s2 disjoint?", sets.IsDisjoint(s1, s2))

	sets.Add(s2, 1)
	fmt.Println("s1 and s2 disjoint after sharing element?", sets.IsDisjoint(s1, s2))
	// Output:
	// s1 and s2 disjoint? true
	// s1 and s2 disjoint after sharing element? false
}

func ExampleIsSubset() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s2, 1)
	sets.Add(s2, 2)
	fmt.Println("s1 ⊆ s2?", sets.IsSubset(s1, s2))

	sets.Add(s1, 1)
	fmt.Println("s1 ⊆ s2 after adding 1?", sets.IsSubset(s1, s2))

	sets.Add(s1, 3)
	fmt.Println("s1 ⊆ s2 after adding 3?", sets.IsSubset(s1, s2))
	// Output:
	// s1 ⊆ s2? true
	// s1 ⊆ s2 after adding 1? true
	// s1 ⊆ s2 after adding 3? false
}

func ExampleIsSuperset() {
	s1 := sets.New[int]()
	s2 := sets.New[int]()
	sets.Add(s1, 1)
	sets.Add(s1, 2)
	fmt.Println("s1 ⊇ s2?", sets.IsSuperset(s1, s2))

	sets.Add(s2, 1)
	fmt.Println("s1 ⊇ s2 after s2 has 1?", sets.IsSuperset(s1, s2))

	sets.Add(s2, 3)
	fmt.Println("s1 ⊇ s2 after s2 has 3?", sets.IsSuperset(s1, s2))
	// Output:
	// s1 ⊇ s2? true
	// s1 ⊇ s2 after s2 has 1? true
	// s1 ⊇ s2 after s2 has 3? false
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

func ExampleSet_String() {
	s := sets.New[int]()
	sets.Add(s, 3)
	sets.Add(s, 1)
	sets.Add(s, 2)
	fmt.Println(s.String())
	// Output:
	// Set{1, 2, 3}
}

func ExampleSet_MarshalJSON() {
	s := sets.New[int]()
	sets.Add(s, 1)
	sets.Add(s, 2)
	data, _ := json.Marshal(s)
	fmt.Println(string(data))
	// Output:
	// [1,2]
}

func ExampleSet_UnmarshalJSON() {
	var s sets.Set[int]
	err := json.Unmarshal([]byte(`[1,2,3]`), &s)
	if err != nil {
		panic(err)
	}
	elems := slices.Collect(sets.Values(&s))
	slices.Sort(elems)
	fmt.Println(elems)
	// Output:
	// [1 2 3]
}
