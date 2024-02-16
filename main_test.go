package main

import (
	"testing"
)

func TestInsertAndInorderTraversal(t *testing.T) {
	root := insert(nil, 2)
	insert(root, 1)
	insert(root, 3)

	result := inorderTraversal(root)

	expectedOutput := "1 2 3 "
	if result != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, result)
	}
}
