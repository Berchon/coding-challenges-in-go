package main

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want bool
	}{
		{"tamanhos diferentes", "abc", "abcd", false},
		{"mesmo tamanho não anagrama", "abc", "def", false},
		{"com espaços válidos", "listen ", "sil ent", true},
		{"espaço extra", "listen", "silent ", false},
		{"ambas vazias", "", "", true},
		{"uma vazia e um espaço", "", " ", false},
		{"case sensitive", "Abc", "abc", false},
		{"unicode acentuação", "olá", "ola", false},
		{"unicode válido", "olá", "láo", true},
		{"caracteres especiais", "ol@", "#ol", false},
		{"símbolos válidos", "ol@", "@ol", true},
		{"anagrama clássico", "listen", "silent", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.s, tt.t)

			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdering(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r []rune
		// Expected output for this test case.
		want []rune
	}{
		{"vazio", []rune{}, []rune{}},
		{"um elemento", []rune{'a'}, []rune{'a'}},
		{"dois elementos", []rune{'b', 'a'}, []rune{'a', 'b'}},
		{"três elementos", []rune{'c', 'a', 'b'}, []rune{'a', 'b', 'c'}},
		{"quatro elementos", []rune{'d', 'b', 'a', 'c'}, []rune{'a', 'b', 'c', 'd'}},
		{"com espaços", []rune{' ', 'a', 'b'}, []rune{' ', 'a', 'b'}},
		{"com caracteres especiais", []rune{'@', '#', '$'}, []rune{'#', '$', '@'}},
		{"com unicode", []rune{'á', 'b', 'c'}, []rune{'b', 'c', 'á'}},
		{"com unicode e espaços", []rune{'á', ' ', 'b', 'c'}, []rune{' ', 'b', 'c', 'á'}},
		{"com unicode, espaços e caracteres especiais", []rune{'á', ' ', 'b', '@', 'c'}, []rune{' ', '@', 'b', 'c', 'á'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ordering(tt.r)
			if !reflect.DeepEqual(tt.r, tt.want) {
				t.Errorf("Ordering() = %v, want %v", tt.r, tt.want)
			}
		})
	}
}
