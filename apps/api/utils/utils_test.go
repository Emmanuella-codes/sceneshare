package utils

import "testing"

func TestFormatTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "zero seconds",
			input:    0,
			expected: "0:00",
		},
		{
			name:     "single digit seconds",
			input:    5,
			expected: "0:05",
		},
		{
			name:     "minutes and seconds",
			input:    83,
			expected: "1:23",
		},
		{
			name:     "exact minutes",
			input:    600,
			expected: "10:00",
		},
		{
			name:     "hours minutes and seconds",
			input:    3605,
			expected: "1:00:05",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatTimestamp(tc.input)
			if got != tc.expected {
				t.Fatalf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
