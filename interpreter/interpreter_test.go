package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {

	program := "program go right end"

	node := ProgramNode{}
	context := NewContext(program)
	node.Parse(context)

	expect := "program: go right "
	if node.ToString() != expect {
		t.Errorf("Expected result to equal %s, but got %s.\n", expect, node.ToString())
	}

}
