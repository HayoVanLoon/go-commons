package http

import "net/http"

type TreeMux interface {
	http.Handler
	Handle(s string, handler http.Handler)
	HandleFunc(s string, handler http.HandlerFunc)
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

func (t *treeMux) Handle(s string, handler http.Handler) {
	t.trie.Add(s, "/", handler)
}

func (t *treeMux) HandleFunc(s string, handler http.HandlerFunc) {
	t.trie.Add(s, "/", handler)
}

func NewTreeMux(notFound http.HandlerFunc) TreeMux {
	if notFound == nil {
		notFound = http.NotFound
	}
	return &treeMux{
		trie:     newWildcardTrie(),
		notFound: notFound,
	}
}
