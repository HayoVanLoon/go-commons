package treemux

import (
	"testing"
)

func TestWildcardTrie_Equals(t *testing.T) {
	cases := []struct {
		left  wildcardTrie
		right wildcardTrie
	}{
		{
			left:  wildcardTrie{k: "foo"},
			right: wildcardTrie{k: "foo"},
		},
		{
			left:  wildcardTrie{k: "foo", v: 1},
			right: wildcardTrie{k: "foo", v: 1},
		},
		{
			left: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 1}}},
			right: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 1}}},
		},
	}

	for _, c := range cases {
		if !c.left.Equals(c.right) {
			t.Errorf("expected left == right")
		}
		if !c.right.Equals(c.left) {
			t.Errorf("expected right == left")
		}
	}
}

func TestWildcardTrie_Equals_Not(t *testing.T) {
	cases := []struct {
		left  wildcardTrie
		right wildcardTrie
	}{
		{
			left:  wildcardTrie{k: "foo"},
			right: wildcardTrie{k: "moo"},
		},
		{
			left:  wildcardTrie{k: "foo", v: 1},
			right: wildcardTrie{k: "foo", v: 2},
		},
		{
			left:  wildcardTrie{k: "foo"},
			right: wildcardTrie{k: "foo", v: 1},
		},
		{
			left: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 1}}},
			right: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 2}}},
		},
		{
			left: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 1}}},
			right: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar"}}},
		},
		{
			left: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 1}}},
			right: wildcardTrie{k: "foo", v: 1,
				children: []wildcardTrie{{k: "bar", v: 1}, {k: "bla", v: 1}}},
		},
	}

	for _, c := range cases {
		if c.left.Equals(c.right) {
			t.Errorf("expected left != right")
		}
		if c.right.Equals(c.left) {
			t.Errorf("expected right != left")
		}
	}
}

func TestWildcardTrie_Equals_Empty(t *testing.T) {
	left := wildcardTrie{}
	right := wildcardTrie{}

	if !left.Equals(right) {
		t.Errorf("expected left == right")
	}
	if !right.Equals(left) {
		t.Errorf("expected right == left")
	}
}

func TestWildcardTrie_Get(t *testing.T) {
	cases := []struct {
		input string
		exp1  interface{}
		exp2  bool
	}{
		{"", -1, true},
		{"/", nil, false},
		{"foo", 2, true},
		{"foo/", nil, false},
		{"/foo", 2, true},
		{"/foo/bar", 3, true},
		{"foo/bar/", nil, false},
		{"foo/slash", 5, true},
		{"foo/slash/", 6, true},
		{"foo/slash/whatever", 6, true},
		{"meow/woof", nil, false},
	}
	tr := wildcardTrie{
		v: -1,
		children: []wildcardTrie{
			{k: "meow", v: 1},
			{k: "foo", v: 2, children: []wildcardTrie{
				{k: "bar", v: 3},
				{k: "bla", v: 4},
				{k: "slash", v: 5, children: []wildcardTrie{
					{k: "", v: 6}}}}}},
	}

	for i, c := range cases {
		act1, act2 := tr.Get(c.input, "/")
		if act1 != c.exp1 || act2 != c.exp2 {
			t.Errorf("%v: [%s] expected (%v, %v), got (%v, %v)", i, c.input, c.exp1, c.exp2, act1, act2)
		}
	}
}

func TestWildcardTrie_Get_Wildcards(t *testing.T) {
	cases := []struct {
		input string
		exp1  interface{}
		exp2  bool
	}{
		{"", -1, true},
		{"/", nil, false},
		{"foo", 2, true},
		{"foo/", 99, true},
		{"/foo", 2, true},
		{"/foo/bar", 3, true},
		{"foo/bar/", nil, false},
		{"foo/slash", 99, true},
		{"foo/slash/", 6, true},
		{"meow/woof", nil, false},
	}
	tr := wildcardTrie{
		v: -1,
		children: []wildcardTrie{
			{k: "meow", v: 1},
			{k: "foo", v: 2, children: []wildcardTrie{
				{k: "bar", v: 3},
				{k: "*", v: 99},
				{k: "slash", v: 5, children: []wildcardTrie{
					{k: "", v: 6}}}}}},
	}

	for i, c := range cases {
		act1, act2 := tr.Get(c.input, "/")
		if act1 != c.exp1 || act2 != c.exp2 {
			t.Errorf("%v: [%s] expected (%v, %v), got (%v, %v)", i, c.input, c.exp1, c.exp2, act1, act2)
		}
	}
}

func TestWildcardTrie_Add_Happy(t *testing.T) {
	steps := []struct {
		input  string
		input2 interface{}
		exp2   wildcardTrie
	}{
		{"foo", 1,
			wildcardTrie{
				k: "", v: nil, children: []wildcardTrie{
					{k: "foo", v: 1}}}},
		{"foo/bar", 2,
			wildcardTrie{
				k: "", v: nil, children: []wildcardTrie{
					{k: "foo", v: 1, children: []wildcardTrie{
						{k: "bar", v: 2}}}}}},
		{"foo/*", 99,
			wildcardTrie{
				k: "", v: nil, children: []wildcardTrie{
					{k: "foo", v: 1, children: []wildcardTrie{
						{k: "bar", v: 2},
						{k: "*", v: 99}}}}}},
		{"foo/slash/", 6,
			wildcardTrie{
				k: "", v: nil, children: []wildcardTrie{
					{k: "foo", v: 1, children: []wildcardTrie{
						{k: "bar", v: 2},
						{k: "*", v: 99},
						{k: "slash", children: []wildcardTrie{
							{k: "", v: 6}}}}}}}},
		{"foo/slash", 5,
			wildcardTrie{
				k: "", v: nil, children: []wildcardTrie{
					{k: "foo", v: 1, children: []wildcardTrie{
						{k: "bar", v: 2},
						{k: "*", v: 99},
						{k: "slash", v: 5, children: []wildcardTrie{
							{k: "", v: 6}}}}}}}},
		{"/foo/bar", 666,
			wildcardTrie{
				k: "", v: nil, children: []wildcardTrie{
					{k: "foo", v: 1, children: []wildcardTrie{
						{k: "bar", v: 666},
						{k: "*", v: 99},
						{k: "slash", v: 5, children: []wildcardTrie{
							{k: "", v: 6}}}}}}}},
	}
	tr := wildcardTrie{k: ""}
	for i, step := range steps {
		tr.Add(step.input, "/", step.input2)

		if !tr.Equals(step.exp2) {
			t.Errorf("%v: [%s]\nexpected: %s,\ngot:      %s", i, step.input, step.exp2, tr)
			break
		}
	}
}
