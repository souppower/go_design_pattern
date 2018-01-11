package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	maker := PageMaker{}
	result := maker.MakeWelcomePage("a@a.com")
	expect := "# Welcome to a's page!"

	if result != expect {
		t.Errorf("Expected result to equal %s, but got %s.", expect, result)
	}
}
