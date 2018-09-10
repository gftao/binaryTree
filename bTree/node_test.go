package bTree

import "testing"

var tests = []struct {
	input  int
	output bool
}{
	{6, true},
	{16, false},
	{3, true},
}

func TestSearch(t *testing.T) {
	tree := &Node{Key: 6, Left: &Node{Key: 3}}

	for i, test := range tests {
		if res := tree.Search(test.input); res != test.output {
			t.Errorf("%d: got %v, expected %v", i, res, test.output)
		}
	}
}

func TestDelete(t *testing.T) {
	tree := &Node{Key: 6, Left: &Node{Key: 3}}
	tree.Delete(3, nil)
}
