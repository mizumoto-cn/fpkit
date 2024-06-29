package functional_test

import (
	"testing"

	"github.com/mizumoto-cn/gogenerics/functional"

	assert "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 6, functional.Sum(1, 2, 3))
	assert.Equal(t, 6.5, functional.Sum(1.5, 2, 3))
	assert.Equal(t, 6+4i, functional.Sum(1, 2, 3+4i))
	assert.Equal(t, 6.5+4i, functional.Sum(1.5, 2, 3+4i))
	assert.Equal(t, 5.54+6i, functional.Sum(1.5+2i, -2.3, 3+4i, 1, 2.34))
}

func TestCompareTo(t *testing.T) {
	assert.Equal(t, 1, functional.CompareTo(3, 2))
	assert.Equal(t, -1, functional.CompareTo(2.3, 33))
	assert.Equal(t, 0, functional.CompareTo(-2, 2.4))
	assert.Equal(t, 1, functional.CompareTo("b", "a"))
	assert.Equal(t, -1, functional.CompareTo("a", "b"))
}

func cmp2[T functional.Orderable](a, b T) bool {
	return functional.CompareTo(a, b) > 0
}

func TestSort(t *testing.T) {
	var cmpInt functional.Comparator[int] = func(a, b int) bool {
		return functional.CompareTo(a, b) < 0
	}
	var cmpString functional.Comparator[string] = func(a, b string) bool {
		return functional.CompareTo(a, b) < 0
	}
	var cmpFloat functional.Comparator[float64] = func(a, b float64) bool {
		return functional.CompareTo(a, b) < 0
	}
	casesInt := []struct {
		in   []int
		want []int
		cmp  functional.Comparator[int]
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}, cmpInt},
		{[]int{3, 2, 1, 4, 5}, []int{1, 2, 3, 4, 5}, cmpInt},
		{[]int{3, 2, 1, 4, 5, 3, 2, 1, 4, 5}, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, cmpInt},
		{[]int{3, 2, 1, 4, 5, 3, 2, 1, 4, 5}, []int{5, 5, 4, 4, 3, 3, 2, 2, 1, 1}, cmp2[int]},
	}

	for _, c := range casesInt {
		functional.Sort(c.in, c.cmp)
		assert.Equal(t, c.want, c.in)
	}

	casesString := []struct {
		in   []string
		want []string
		cmp  functional.Comparator[string]
	}{
		{[]string{"c", "b", "a"}, []string{"a", "b", "c"}, cmpString},
		{[]string{"c", "b", "a", "d", "e"}, []string{"a", "b", "c", "d", "e"}, cmpString},
		{[]string{"c", "b", "a", "d", "e", "c", "b", "a", "d", "e"}, []string{"a", "a", "b", "b", "c", "c", "d", "d", "e", "e"}, cmpString},
		{[]string{"c", "b", "a", "d", "e", "c", "b", "a", "d", "e"}, []string{"e", "e", "d", "d", "c", "c", "b", "b", "a", "a"}, cmp2[string]},
	}
	for _, c := range casesString {
		functional.Sort(c.in, c.cmp)
		assert.Equal(t, c.want, c.in)
	}

	casesFloat := []struct {
		in   []float64
		want []float64
		cmp  functional.Comparator[float64]
	}{
		{[]float64{3.3, 2.2, 1.1}, []float64{1.1, 2.2, 3.3}, cmpFloat},
		{[]float64{3.3, 2.2, 1.1, 4.4, 5.5}, []float64{1.1, 2.2, 3.3, 4.4, 5.5}, cmpFloat},
		{[]float64{3.3, 2.2, 1.1, 4.4, 5.5, 3.3, 2.2, 1.1, 4.4, 5.5}, []float64{1.1, 1.1, 2.2, 2.2, 3.3, 3.3, 4.4, 4.4, 5.5, 5.5}, cmpFloat},
		{[]float64{3.3, 2.2, 1.1, 4.4, 5.5, 3.3, 2.2, 1.1, 4.4, 5.5}, []float64{5.5, 5.5, 4.4, 4.4, 3.3, 3.3, 2.2, 2.2, 1.1, 1.1}, cmp2[float64]},
	}
	for _, c := range casesFloat {
		functional.Sort(c.in, c.cmp)
		assert.Equal(t, c.want, c.in)
	}

}
