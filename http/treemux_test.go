package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testHandler struct {
}

func (testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(r.URL.Path + "!"))
}

func TestTreeMux_Handle(t *testing.T) {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("foo!bar!"))
	}

	tr := NewTreeMux(nil)
	tr.HandleFunc("foo/bar", handleFunc)
	tr.HandleFunc("/moo", handleFunc)
	tr.Handle("/moo/", testHandler{})

	s := httptest.NewServer(tr)
	defer s.Close()

	cases := []struct {
		path string
		code int
		body string
	}{
		{"/foo/bar", 200, "foo!bar!"},
		{"/foo/bla", 404, ""},
		{"/moo", 200, "foo!bar!"},
		{"/moo/", 200, "/moo/!"},
		{"/moo/meh", 200, "/moo/meh!"},
	}
	for i, c := range cases {
		resp, err := s.Client().Get(s.URL + c.path)
		if err != nil {
			t.Errorf("%v %s: unexpected error %v", i, c.path, err)
			continue
		}
		if resp.StatusCode != c.code {
			t.Errorf("%v %s: expected %v, got %v", i, c.path, c.code, resp.StatusCode)
			continue
		}
		if c.code != 200 {
			continue
		}
		bs, _ := ioutil.ReadAll(resp.Body)
		if string(bs) != c.body {
			t.Errorf("%v %s: expected %s, got %v", i, c.path, c.body, string(bs))
		}
	}
}
