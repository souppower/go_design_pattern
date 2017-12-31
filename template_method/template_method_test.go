package templateMethod

import (
	"testing"
)

func TestCharDisplay(t *testing.T) {
	display := &CharDisplay{Char: 'A'}
	result := display.Display(display)
	if result != "<<AAAAA>>" {
		t.Errorf("Expected result to be %s, but got %s.\n", "<<AAAAA>>", result)
	}
}

func TestStringDisplay(t *testing.T) {
	expect := "+-------+\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n+-------+\n"
	display := &StringDisplay{Str: "ABCDE"}
	result := display.Display(display)
	if result != expect {
		t.Errorf("Expected result to be \n%s, but got \n%s.\n", expect, result)
	}
}
