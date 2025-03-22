package auth

import (
	"testing"
)

func TestWildcardMatch(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		text    string
		want    bool
	}{
		{"Empty pattern and text", "", "", true},
		{"Empty pattern non-empty text", "", "abc", false},
		{"Star matches anything", "*", "anything", true},
		{"Simple exact match", "abc", "abc", true},
		{"Simple non-match", "abc", "def", false},
		{"Trailing wildcard", "abc*", "abcdef", true},
		{"Leading wildcard", "*abc", "xyzabc", true},
		{"Middle wildcard", "a*c", "abc", true},
		{"Middle wildcard multiple chars", "a*c", "abbc", true},
		{"Question mark single char", "a?c", "abc", true},
		{"Question mark wrong char count", "a?c", "abbc", false},
		{"Complex pattern", "a*b?c*", "axxbyc123", true},
		{"Complex pattern non-match", "a*b?c*", "axxbyyd123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WildcardMatch(tt.pattern, tt.text); got != tt.want {
				t.Errorf("WildcardMatch(%q, %q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
			}
		})
	}
}
