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
package i18n

func IsQuestionMark(c rune) bool {
	for _, q := range [4]rune{'?', '؟', '⁇', '？'} {
		if q == c {
			return true
		}
	}
	return false
}

func IsExclamationMark(c rune) bool {
	for _, q := range [4]rune{'!', 'ǃ', '՜', '！'} {
		if q == c {
			return true
		}
	}
	return false
}
