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

// Small extensions on the math package
package mathext

import "math"

func Round(f float64) int {
	if math.Signbit(f) {
		return int(f - .5)
	} else {
		return int(f + .5)
	}
}

// Translates a cartesian coordinate into a polar coordinate.
func ToPolar(x, y float64) (r float64, arc float64) {
	r = math.Sqrt(x*x + y*y)
	arc = math.Atan2(y, x)
	return
}

// Translates a polar coordinate into a cartesian coordinate.
func ToCartesian(r, arc float64) (x float64, y float64) {
	x = r * math.Cos(arc)
	y = r * math.Sin(arc)
	return
}

// Calculates the hypothenuse from the two other sides of a right triangle.
func Pyth(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
