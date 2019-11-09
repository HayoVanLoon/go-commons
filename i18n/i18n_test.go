package i18n

import (
	"testing"
)

func TestIsQuestionMark(t *testing.T) {
	cases := []struct {
		input    rune
		expected bool
	}{
		{'?', true},
		{'n', false},
		{'؟', true},
		{'⁇', true},
		{'？', true},
	}
	for _, c := range cases {
		if IsQuestionMark(c.input) != c.expected {
			t.Errorf("failed for %s (expected: %v)", string(c.input), c.expected)
		}
	}
}
