package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	loginForm := NewLoginForm()

	state := loginForm.Button.Enabled
	if state {
		t.Errorf("Expected state to be false, but got true.\n")
	}

	loginForm.RadioButton.Check(true)

	state = loginForm.Button.Enabled
	if !state {
		t.Errorf("Expect state to be true, but got false.\n")
	}

}
