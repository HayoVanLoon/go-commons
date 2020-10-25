package http

import (
	"fmt"
	"strings"
)

type wildcardTrie struct {
	k        string
	v        interface{}
	children []wildcardTrie
}

func newWildcardTrie() *wildcardTrie {
	return &wildcardTrie{k: ""}
}

// Breaks up a string using the specified separator and adds the data to the
// trie. When a path already exists in the trie, the old data is overwritten.
//
// The first element is always expected to be empty. Therefore following
// statements are idempotent.
//   trie.Add("/foo/bar", "/", 1)
//   trie.Add("foo/bar", "/", 1)
//
// The wildcard is a flexible, retrieval-time parameter. It plays no role
// whatsoever at construction-time. One could even apply different wildcard
// schemes for different purposes on the same trie.
// See Get for more details on wildcard behaviour.
func (t *wildcardTrie) Add(s, sep string, v interface{}) {
	xs := strings.Split(s, sep)
	if xs[0] == "" {
		t.grow(1, xs, v)
		return
	}
	t.grow(0, xs, v)
}

func (t *wildcardTrie) grow(idx int, xs []string, v interface{}) {
	if len(xs) == idx {
		t.v = v
		return
	}
	for i := range t.children {
		if t.children[i].k == xs[idx] {
			t.children[i].grow(idx+1, xs, v)
			return
		}
	}
	if len(xs) > idx {
		c := wildcardTrie{k: xs[idx], children: nil}
		if len(xs) == idx+1 {
			c.v = v
		} else {
			c.grow(idx+1, xs, v)
		}
		t.children = append(t.children, c)
	}
}

// Attempts to retrieve the data from the specified path, split up by the
// specified separator using the default wildcard "*".
//
// Wildcard elements hold no special status over other elements. When, due to a
// wildcard, a path has two valid end points, the one inserted earliest wins.
func (t *wildcardTrie) Get(s, sep string) (interface{}, bool) {
	return t.GetWithWildcard(s, sep, "*")
}

// Attempts to retrieve the data from the specified path, split up by the
// specified separator using the specified wildcard.
func (t *wildcardTrie) GetWithWildcard(s, sep, wildcard string) (interface{}, bool) {
	// TODO(hvl): input validation
	xs := strings.Split(s, sep)
	if xs[0] == "" {
		return t.get(0, xs, wildcard)
	}
	for _, c := range t.children {
		if v, found := c.get(0, xs, wildcard); found {
			return v, found
		}
	}
	return nil, false
}

func (t *wildcardTrie) get(idx int, xs []string, wildcard string) (interface{}, bool) {
	if xs[idx] != t.k && t.k != wildcard {
		if t.k == "" && len(t.children) == 0 {
			return t.v, true
		}
		return nil, false
	}
	if len(xs)-idx == 1 {
		return t.v, true
	}
	for _, c := range t.children {
		if v, found := c.get(idx+1, xs, wildcard); found {
			return v, found
		}
	}
	return nil, false
}

func (t *wildcardTrie) Equals(other wildcardTrie) bool {
	if t.k != other.k || t.v != other.v || len(t.children) != len(other.children) {
		return false
	}
	for i, c := range t.children {
		if !c.Equals(other.children[i]) {
			return false
		}
	}
	return true
}

func (t wildcardTrie) String() string {
	b := &strings.Builder{}
	t.string(b)
	return b.String()
}

func (t *wildcardTrie) string(b *strings.Builder) {
	b.WriteString("{\"")
	b.WriteString(t.k)
	b.WriteString(fmt.Sprintf("\"=%v", t.v))
	if len(t.children) > 0 {
		b.WriteString(",[")
		t.children[0].string(b)
		for i := 1; i < len(t.children); i += 1 {
			b.WriteRune(',')
			t.children[i].string(b)
		}
		b.WriteRune(']')
	}
	b.WriteRune('}')
}
