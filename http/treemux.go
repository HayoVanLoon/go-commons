package http

import "net/http"

// A TreeMux is a request multiplexer that uses a tree structure to route
// requests.
//
// Wildcards ("*") can be used to map multiple URLs to the same function.
//
// Example:
// After the following mapping:
//   t.Handle("/foo/*/bla", fn)
// The following two requests would each be handled by `fn`:
//   "/foo/bar/bla"
//   "/foo/moo/bla"
//
type TreeMux interface {
	http.Handler

	// Add a new http.Handler for the given path. When a path already exists in
	// the tree, the old data is overwritten.
	//
	// The first element is always expected to be empty. Therefore following
	// statements are idempotent.
	//   t.Handle("/foo/bar", fn)
	//   t.Handle("foo/bar", fn)
	//
	// The wildcard is a flexible, retrieval-time parameter. It plays no role
	// whatsoever at construction-time.
	Handle(path string, handler http.Handler)

	// Add a new http.HandlerFunc for the given path. See Handle for more
	// details.
	HandleFunc(path string, handler http.HandlerFunc)
}

type treeMux struct {
	trie     *wildcardTrie
	notFound http.HandlerFunc
}

func (t treeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, found := t.trie.Get(r.URL.Path, "/")
	if !found {
		t.notFound(w, r)
		return
	}
	switch h := v.(type) {
	case http.Handler:
		h.ServeHTTP(w, r)
	case http.HandlerFunc:
		h(w, r)
	}
}

func (t *treeMux) Handle(path string, handler http.Handler) {
	t.trie.Add(path, "/", handler)
}

func (t *treeMux) HandleFunc(path string, handler http.HandlerFunc) {
	t.trie.Add(path, "/", handler)
}

// Creates a new tree-based request multiplexer. If a request cannot be matched,
// the standard http.NotFound will be used.
func NewTreeMux() TreeMux {
	return NewTreeMuxWithNotFound(nil)
}

// Creates a new tree-based request multiplexer. If a request cannot be matched,
// the specified HandlerFunc will be used. If set to `nil`, the default
// http.NotFound will be used.
func NewTreeMuxWithNotFound(notFound http.HandlerFunc) TreeMux {
	if notFound == nil {
		notFound = http.NotFound
	}
	return &treeMux{
		trie:     newWildcardTrie(),
		notFound: notFound,
	}
}
