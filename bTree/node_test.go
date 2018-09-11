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

func TestNode_Search(t *testing.T) {
	tree := &Node{Key: 6, Left: &Node{Key: 3}}

	for i, test := range tests {
		if res := tree.Search(test.input); res != test.output {
			t.Errorf("%d: got %v, expected %v", i, res, test.output)
		}
	}
}

func TestNode_Delete(t *testing.T) {
	tree := &Node{Key: 6}

	tree.insert(15)
	tree.insert(13)
	tree.insert(14)
	tree.delete(6, nil)
	if tree.String() != "13 14 15" {
		t.Errorf("got [%s], expected [%v]", tree.String(), "13 14 15")
	}
}
