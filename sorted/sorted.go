package sorted

import "math/rand"

type StringSet interface {
	Add(string) StringSet
	Has(string) bool
	Remove(string) StringSet
	Size() int

	// Returns the set as a sorted slice.
	//
	// This slice is a copy and can thus safely be modified.
	Slice() []string
}

type StringList interface {
	Add(string) StringSet
	Has(string) bool
	Remove(string) StringSet
	Size() int
	Get(int) string

	// Returns the set as a sorted slice.
	//
	// This slice is a copy and can thus safely be modified.
	Slice() []string
}

// Implementation of a sorted set
type stringList struct {
	data []string
	noDuplicates bool
}

func NewStringSet() StringSet {
	return &stringList{make([]string, 0, 10), true}
}

func NewStringList() StringList {
	return &stringList{make([]string, 0, 10), false}
}

func (ss *stringList) Add(s string) StringSet {
	if len(ss.data) == 0 {
		ss.data = append(ss.data, s)
		return ss
	}

	lower, upper := 0, len(ss.data)
	for ; lower != upper; {
		i := lower + rand.Intn(upper - lower)
		if s < ss.data[i] {
			upper = i
		} else if s > ss.data[i] {
			lower = i + 1
		} else {
			if ss.noDuplicates {
				return ss
			} else {
				break
			}
		}
	}

	ss.data = append(ss.data, ss.data[0])
	copy(ss.data[lower+1:], ss.data[lower:])
	ss.data[lower] = s

	return ss
}

func (ss *stringList) Has(s string) bool{
	lower, upper := 0, len(ss.data)
	for ; lower != upper; {
		i := lower + rand.Intn(upper - lower)
		if s < ss.data[i] {
			upper = i
		} else if s > ss.data[i] {
			lower = i + 1
		} else {
			return true
		}
	}
	return false
}

func (ss *stringList) Remove(s string) StringSet {
	if len(ss.data) == 0 {
		return ss
	}

	lower, upper := 0, len(ss.data)
	for ; lower != upper; {
		i := lower + rand.Intn(upper - lower)
		if s < ss.data[i] {
			upper = i
		} else if s > ss.data[i] {
			lower = i + 1
		} else {
			copy(ss.data[i:], ss.data[i+1:])
			ss.data[len(ss.data)-1] = ""
			ss.data = ss.data[:len(ss.data)-1]
			return ss
		}
	}

	return ss
}

func (ss *stringList) Size() int {
	return len(ss.data)
}

func (ss *stringList) Slice() []string {
	dst := make([]string, len(ss.data))
	copy(dst, ss.data)
	return dst
}

func (ss *stringList) Get(i int) string {
	if i < 0 {
		panic("index < 0")
	}
	if i >= ss.Size() {
		panic("index past end of list")
	}

	return ss.data[i]
}