package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {

	rootDir := NewDirectory("root")
	usrDir := NewDirectory("usr")
	fileA := NewFile("A", 1)

	rootDir.Add(usrDir)
	rootDir.Add(fileA)

	fileB := NewFile("B", 2)
	usrDir.Add(fileB)

	result := rootDir.PrintList("")

	expect := "/root (3)\n/root/usr (2)\n/root/usr/B (2)\n/root/A (1)\n"
	if result != expect {
		t.Errorf("Expected result to equal %s, but got %s.\n", expect, result)
	}

}
