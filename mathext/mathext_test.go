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
package mathext

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	cases := []struct {
		msg      string
		input    float64
		expected int
	}{
		{"failed", 0, 0},
		{"failed", .1, 0},
		{"failed", .49, 0},
		{"failed", .5, 1},
		{"failed", .51, 1},
		{"failed", .9, 1},
		{"failed", 100.9, 101},
		{"failed", -.1, 0},
		{"failed", -.49, 0},
		{"failed", -.5, -1},
		{"failed", -.51, -1},
		{"failed", -.9, -1},
		{"failed", -100.9, -101},
	}
	for i, c := range cases {
		if actual := Round(c.input); actual != c.expected {
			t.Errorf("case %v: %s (expected: %v, got: %v)", i, c.msg, c.expected, actual)
		}
	}
}

func TestToCartesian(t *testing.T) {
	cases := []struct {
		msg string
		r   float64
		arc float64
		x   float64
		y   float64
	}{
		{"Origin", 0, 0, 0, 0},
		{"0", 10, 0, 10, 0},
		{"45", Pyth(10, 10), math.Pi * .25, 10, 10},
		{"90", 10, math.Pi * .5, 0, 10},
		{"135", Pyth(10, 10), math.Pi * .75, -10, 10},
		{"180", 10, math.Pi, -10, 0},
		{"225", Pyth(10, 10), math.Pi * 1.25, -10, -10},
		{"270", 10, math.Pi * 1.5, 0, -10},
		{"315", Pyth(10, 10), math.Pi * 1.75, 10, -10},
		{"360", 10, math.Pi * 2, 10, 0},
	}

	for i, c := range cases {
		// rather hefty rounding errors, chose big margin: 10 vs 9.899
		if x, y := ToCartesian(c.r, c.arc); Diff(x, c.x, .2) || Diff(y, c.y, .2) {
			t.Errorf("case %v: %s (expected: (%v,%v), got: (%v,%v))", i, c.msg, c.x, c.y, x, y)
		}
	}
}

func TestToPolar(t *testing.T) {
	cases := []struct {
		msg string
		x   float64
		y   float64
		r   float64
		arc float64
	}{
		{"Origin", 0, 0, 0, 0},
		{"0", 10, 0, 10, 0},
		{"45", 10, 10, Pyth(10, 10), math.Pi * .25},
		{"90", 0, 10, 10, math.Pi * .5},
		{"135", -10, 10, Pyth(10, 10), math.Pi * .75},
		{"180", -10, 0, 10, math.Pi},
		{"225", -10, -10, Pyth(10, 10), -math.Pi * .75},
		{"270", 0, -10, 10, -math.Pi * .5},
		{"315", 10, -10, Pyth(10, 10), -math.Pi * .25},
	}
	for i, c := range cases {
		// rather hefty rounding errors, chose big margin
		if r, arc := ToPolar(c.x, c.y); Diff(r, c.r, .2) || Diff(arc, c.arc, .2) {
			t.Errorf("case %v: %s (expected: (%v,%v), got: (%v,%v))", i, c.msg, c.r, c.arc, r, arc)
		}
	}
}

func TestPyth(t *testing.T) {
	cases := []struct {
		msg string
		x   float64
		y   float64
		z   float64
	}{
		{"", 1, 1, math.Sqrt(2)},
		{"", 1, 2, math.Sqrt(5)},
		{"", -1, 2, math.Sqrt(5)},
		{"", 1, -2, math.Sqrt(5)},
		{"edge case", 0, 2, 2},
	}
	for i, c := range cases {
		if z := Pyth(c.x, c.y); Diff(z, c.z, .01) {
			t.Errorf("case %v: %s (expected: %v, got: %v)", i, c.msg, c.z, z)
		}
	}
}

func TestDiff(t *testing.T) {
	cases := []struct {
		msg string
		a   float64
		b   float64
		d   float64
		o   bool
	}{
		{"", 1.1, 1, .2, true},
		{"", 1.1, 1, .05, false},
		{"", 1.1, 1, .1, false},
		{"", 1.1, 1, -.2, true},
		{"", 1.1, 1, -.05, false},
		{"", 0, 1, 2, true},
		{"", -1.1, 1, .1, false},
	}
	for i, c := range cases {
		if o := Diff(c.a, c.b, c.d); o != c.o {
			t.Errorf("case %v: %s (expected: %v, got: %v)", i, c.msg, c.o, o)
		}
	}
}
