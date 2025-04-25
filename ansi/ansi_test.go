package ansi

import (
	"testing"
)

func TestActualIndex(t *testing.T) {
	tests := []struct {
		name         string
		ansiStr      string
		displayIndex int
		expected     int
		err          string
	}{
		{
			name:         "No ANSI codes, valid index",
			ansiStr:      "Hello, World!",
			displayIndex: 7,
			expected:     7,
		},
		{
			name:         "With ANSI codes, valid index",
			ansiStr:      "\x1b[31mHello\x1b[0m, \x1b[32mWorld!\x1b[0m",
			displayIndex: 7,
			expected:     21,
		},
		{
			name:         "With ANSI codes, index at start",
			ansiStr:      "\x1b[31mHello\x1b[0m, \x1b[32mWorld!\x1b[0m",
			displayIndex: 0,
			expected:     8,
		},
		{
			name:         "With ANSI codes, index out of range",
			ansiStr:      "\x1b[31mHello\x1b[0m, \x1b[32mWorld!\x1b[0m",
			displayIndex: 50,
			err:          "index out of range",
		},
		{
			name:         "Empty string",
			ansiStr:      "",
			displayIndex: 0,
			err:          "index out of range",
		},
		{
			name:         "Only ANSI codes",
			ansiStr:      "\x1b[31m\x1b[0m\x1b[32m\x1b[0m",
			displayIndex: 0,
			err:          "index out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ActualIndex(tt.ansiStr, tt.displayIndex)
			if actual != tt.expected {
				t.Errorf("ActualIndex(%q, %d) = %d; want %d", tt.ansiStr, tt.displayIndex, actual, tt.expected)
			}
			if err != nil && err.Error() != tt.err {
				t.Errorf("ActualIndex(%q, %d) error = %v; want %v", tt.ansiStr, tt.displayIndex, err, tt.err)
			}
			if err == nil && tt.err != "" {
				t.Errorf("Expected error but got nil for ActualIndex(%q, %d)", tt.ansiStr, tt.displayIndex)
			}
		})
	}
}
