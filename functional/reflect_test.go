package functional_test

import (
	"reflect"
	"testing"

	"github.com/mizumoto-cn/fpkit/functional"

	"github.com/stretchr/testify/assert"
)

// . Tests for PtrOf
func TestPtrOf(t *testing.T) {
	vint := 1
	vstring := "hello"
	vstruct := struct{ Name string }{Name: "mizumoto"}
	cases := []struct {
		title string
		v     any
		want  any
	}{
		{
			title: "int",
			v:     vint,
			want:  vint,
		},
		{
			title: "string",
			v:     vstring,
			want:  vstring,
		},
		{
			title: "struct",
			v:     vstruct,
			want:  vstruct,
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got := functional.PtrOf(c.v)
			// To properly test the PtrOf function, you should compare the values pointed to by the pointers for equality, not the pointers themselves.
			assert.True(t, reflect.DeepEqual(reflect.ValueOf(got).Elem().Interface(), c.want))
		})
	}
}
