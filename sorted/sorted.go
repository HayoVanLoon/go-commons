/*
 * Copyright 2019 Hayo van Loon
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package sorted

import (
	"fmt"
	"math/rand"
	"strings"
)

// Practice type, nothing too serious
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

// Practice type, nothing too serious
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
	data         []string
	noDuplicates bool
}

func NewStringSet() StringSet {
	return &stringList{[]string{}, true}
}

func NewStringList() StringList {
	return &stringList{[]string{}, false}
}

func (sl *stringList) add(s string) interface{} {
	if len(sl.data) == 0 {
		sl.data = append(sl.data, s)
		return sl
	}

	lower, upper := 0, len(sl.data)
	for lower != upper {
		i := lower + rand.Intn(upper-lower)
		if s < sl.data[i] {
			upper = i
		} else if s > sl.data[i] {
			lower = i + 1
		} else {
			if sl.noDuplicates {
				return sl
			} else {
				break
			}
		}
	}

	sl.data = append(sl.data, sl.data[0])
	copy(sl.data[lower+1:], sl.data[lower:])
	sl.data[lower] = s

	return sl
}

func (sl *stringList) Has(s string) bool {
	lower, upper := 0, len(sl.data)
	for lower != upper {
		i := lower + rand.Intn(upper-lower)
		if s < sl.data[i] {
			upper = i
		} else if s > sl.data[i] {
			lower = i + 1
		} else {
			return true
		}
	}
	return false
}

func (sl *stringList) remove(s string) interface{} {
	if len(sl.data) == 0 {
		return sl
	}

	lower, upper := 0, len(sl.data)
	for lower != upper {
		i := lower + rand.Intn(upper-lower)
		if s < sl.data[i] {
			upper = i
		} else if s > sl.data[i] {
			lower = i + 1
		} else {
			copy(sl.data[i:], sl.data[i+1:])
			sl.data[len(sl.data)-1] = ""
			sl.data = sl.data[:len(sl.data)-1]
			return sl
		}
	}

	return sl
}

func (sl *stringList) Size() int {
	return len(sl.data)
}

func (sl *stringList) Slice() []string {
	dst := make([]string, len(sl.data))
	copy(dst, sl.data)
	return dst
}

func (sl *stringList) Get(i int) string {
	if i < 0 {
		panic("index < 0")
	}
	if i >= sl.Size() {
		panic("index past end of list")
	}

	return sl.data[i]
}

func (sl *stringList) String() string {
	return fmt.Sprintf("{%v}", strings.Join(sl.data, ","))
}

func (sl *stringList) Insert(s string) StringList {
	return sl.add(s).(StringList)
}

func (sl *stringList) Delete(s string) StringList {
	return sl.remove(s).(StringList)
}

func (sl *stringList) Add(s string) StringSet {
	return sl.add(s).(StringSet)
}

func (sl *stringList) Remove(s string) StringSet {
	return sl.remove(s).(StringSet)
}
