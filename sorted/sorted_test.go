package sorted

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestStringSet_Add(t *testing.T) {
	cases := []struct {
		input    StringSet
		items    []string
		expected StringSet
	}{
		{&stringList{noDuplicates: true}, []string{"1"}, &stringList{[]string{"1"}, true}},
		{&stringList{noDuplicates: true}, []string{"1", "2", "3"}, &stringList{[]string{"1", "2", "3"}, true}},
		{&stringList{noDuplicates: true}, []string{"1", "12", "3"}, &stringList{[]string{"1", "12", "3"}, true}},
		{&stringList{noDuplicates: true}, []string{"1", "3", "2"}, &stringList{[]string{"1", "2", "3"}, true}},
		{&stringList{noDuplicates: true}, []string{"2", "3", "1"}, &stringList{[]string{"1", "2", "3"}, true}},
		{&stringList{noDuplicates: true}, []string{"3", "2", "1"}, &stringList{[]string{"1", "2", "3"}, true}},
	}
	for _, c := range cases {

		for _, a := range c.items {
			c.input.Add(a)
		}

		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("expected %v, got %v", c.expected, c.input)
		}
	}
}

func TestStringList_Add(t *testing.T) {
	cases := []struct {
		input    StringSet
		items    []string
		expected StringSet
	}{
		{&stringList{noDuplicates: false}, []string{"1"}, &stringList{[]string{"1"}, false}},
		{&stringList{noDuplicates: false}, []string{"1", "2", "3"}, &stringList{[]string{"1", "2", "3"}, false}},
		{&stringList{noDuplicates: false}, []string{"1", "12", "3"}, &stringList{[]string{"1", "12", "3"}, false}},
		{&stringList{noDuplicates: false}, []string{"1", "12", "1", "3"}, &stringList{[]string{"1", "1", "12", "3"}, false}},
		{&stringList{noDuplicates: false}, []string{"1", "3", "2"}, &stringList{[]string{"1", "2", "3"}, false}},
		{&stringList{noDuplicates: false}, []string{"2", "3", "1"}, &stringList{[]string{"1", "2", "3"}, false}},
		{&stringList{noDuplicates: false}, []string{"3", "2", "1"}, &stringList{[]string{"1", "2", "3"}, false}},
	}
	for _, c := range cases {

		for _, a := range c.items {
			c.input.Add(a)
		}

		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("expected %v, got %v", c.expected, c.input)
		}
	}
}

func TestStringSet_AddRandom(t *testing.T) {
	isOk := func(xs []string) bool {
		last := ""
		for _, x := range (xs) {
			if x < last {
				return false
			}
			last = x
		}
		return true
	}

	for i := 0; i < 5; i += 1 {
		xs := NewStringList()
		for i := 0; i < 20; i += 1 {
			xs.Insert(strconv.Itoa(rand.Int()))
		}
		if !isOk(xs.Slice()) {
			t.Errorf("incorrectly sorted")
		}
	}
}

func TestStringSet_Remove(t *testing.T) {
	cases := []struct {
		input    StringSet
		items    []string
		expected StringSet
	}{
		{&stringList{[]string{"1", "2", "3"}, true}, []string{"1"}, &stringList{[]string{"2", "3"}, true}},
		{&stringList{[]string{"1", "2", "3"}, true}, []string{"2"}, &stringList{[]string{"1", "3"}, true}},
		{&stringList{[]string{"1", "2", "3"}, true}, []string{"3"}, &stringList{[]string{"1", "2"}, true}},
		{&stringList{[]string{"1", "2", "2", "3"}, true}, []string{"2"}, &stringList{[]string{"1", "2", "3"}, true}},
		{&stringList{[]string{"1", "2", "3"}, true}, []string{"4"}, &stringList{[]string{"1", "2", "3"}, true}},
		{&stringList{[]string{"1"}, true}, []string{"1"}, &stringList{[]string{}, true}},
		{&stringList{[]string{}, true}, []string{"1"}, &stringList{[]string{}, true}},
		{&stringList{[]string{"1", "2", "3"}, true}, []string{"1", "2", "3"}, &stringList{[]string{}, true}},
	}
	for i, c := range cases {

		for _, a := range c.items {
			c.input.Remove(a)
		}

		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("case %v: expected %v, got %v", i, c.expected, c.input)
		}
	}
}

func TestStringList_Remove(t *testing.T) {
	cases := []struct {
		input    StringSet
		items    []string
		expected StringSet
	}{
		{&stringList{[]string{"1", "2", "3"}, false}, []string{"1"}, &stringList{[]string{"2", "3"}, false}},
		{&stringList{[]string{"1", "2", "3"}, false}, []string{"2"}, &stringList{[]string{"1", "3"}, false}},
		{&stringList{[]string{"1", "2", "3"}, false}, []string{"3"}, &stringList{[]string{"1", "2"}, false}},
		{&stringList{[]string{"1", "2", "3"}, false}, []string{"4"}, &stringList{[]string{"1", "2", "3"}, false}},
		{&stringList{[]string{"1"}, false}, []string{"1"}, &stringList{[]string{}, false}},
		{&stringList{[]string{}, false}, []string{"1"}, &stringList{[]string{}, false}},
		{&stringList{[]string{"1", "2", "3"}, false}, []string{"1", "2", "3"}, &stringList{[]string{}, false}},
	}
	for i, c := range cases {

		for _, a := range c.items {
			c.input.Remove(a)
		}

		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("case %v: expected %v, got %v", i, c.expected, c.input)
		}
	}
}

func TestStringSet_Has(t *testing.T) {
	cases := []struct {
		input    StringSet
		item     string
		expected bool
	}{
		{&stringList{[]string{"1", "2", "3"}, true}, "1", true},
		{&stringList{[]string{"1", "2", "3"}, true}, "2", true},
		{&stringList{[]string{"1", "2", "3"}, true}, "3", true},
		{&stringList{[]string{"1", "2", "3"}, true}, "4", false},
		{&stringList{[]string{}, true}, "4", false},
	}
	for i, c := range cases {

		if output := c.input.Has(c.item); output != c.expected {
			t.Errorf("case %v: expected %v, got %v", i, c.expected, output)
		}
	}
}

func TestStringSet_Size(t *testing.T) {
	cases := []struct {
		input    StringSet
		expected int
	}{
		{&stringList{[]string{"1", "2", "3"}, true}, 3},
		{&stringList{[]string{"1", "2"}, true}, 2},
		{&stringList{[]string{"1"}, true}, 1},
		{&stringList{[]string{}, true}, 0},
	}
	for i, c := range cases {

		if output := c.input.Size(); output != c.expected {
			t.Errorf("case %v: expected %v, got %v", i, c.expected, output)
		}
	}
}

func TestStringList_Get(t *testing.T) {
	cases := []struct {
		input    StringList
		i        int
		expected string
		pan      bool
	}{
		{&stringList{[]string{"1", "2", "3"}, true}, 0, "1", false},
		{&stringList{[]string{"1", "2", "3"}, true}, 1, "2", false},
		{&stringList{[]string{"1", "2", "3"}, true}, 2, "3", false},
		{&stringList{[]string{"1", "2", "3"}, true}, 3, "", true},
		{&stringList{[]string{}, true}, 0, "", true},
		{&stringList{[]string{"1"}, true}, 1, "", true},
	}
	for i, c := range cases {

		wrapped := func(idx int) string {
			defer func() {
				if err := recover(); (err != nil) != c.pan {
					t.Errorf("case %v: (un)expected panic", i)
				}
			}()
			return c.input.Get(c.i);
		}

		if output := wrapped(c.i); output != c.expected {
			t.Errorf("case %v: expected %v, got %v", i, c.expected, output)
		}

	}
}

func TestStringSet_Slice(t *testing.T) {
	cases := []struct {
		input    StringSet
		expected []string
	}{
		{&stringList{[]string{"1", "2", "3"}, true}, []string{"1", "2", "3"}},
		{&stringList{[]string{"1", "2"}, true}, []string{"1", "2"}},
		{&stringList{[]string{"1"}, true}, []string{"1"}},
		{&stringList{[]string{}, true}, []string{}},
	}
	for i, c := range cases {

		output := c.input.Slice()
		if !reflect.DeepEqual(output, c.expected) {
			t.Errorf("case %v: expected %v, got %v", i, c.expected, output)
		}

		if len(output) == 0 {
			output = append(output, "foo")
		} else {
			output[0] = "foo"
		}
		if !reflect.DeepEqual(c.input, &stringList{c.expected, true}) {
			t.Errorf("case %v: input changed: expected %v, got %v", i, c.expected, c.input)
		}
	}
}
