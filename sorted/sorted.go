package sorted

import (
	"fmt"
	"math/rand"
	"strings"
)

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
	Insert(string) StringList
	Has(string) bool
	Delete(string) StringList
	Size() int
	Get(int) string

	// Returns the set as a sorted slice.
	//
	// This slice is a copy and can thus safely be modified.
	Slice() []string
}

// A simple implementation of a strings list with low overhead.
//
// Also implements the set through an extra flag.
// Performs all operations in O(log n).
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

func (ss *stringList) add(s string) interface{} {
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

func (ss *stringList) remove(s string) interface{} {
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

func (ss *stringList) String() string {
	return fmt.Sprintf("{%v}", strings.Join(ss.data, ","))
}

func (ss *stringList) Insert(s string) StringList {
	return ss.add(s).(StringList)
}

func (ss *stringList) Delete(s string) StringList {
	return ss.remove(s).(StringList)
}

func (ss *stringList) Add(s string) StringSet {
	return ss.add(s).(StringSet)
}

func (ss *stringList) Remove(s string) StringSet {
	return ss.remove(s).(StringSet)
}
