package mediator

type mediator interface {
	createColleagues()
	colleagueChanged()
}

type colleague interface {
	setMediator(mediator mediator)
	setColleagueEnabled(enabled bool)
}

type button struct {
	Enabled  bool
	mediator mediator
}

func (b *button) setMediator(mediator mediator) {
	b.mediator = mediator
}

func (b *button) setColleagueEnabled(enabled bool) {
	b.Enabled = enabled
}

type radioButton struct {
	enabled  bool
	checked  bool
	mediator mediator
}

func (rb *radioButton) setMediator(mediator mediator) {
	rb.mediator = mediator
}

func (rb *radioButton) setColleagueEnabled(enabled bool) {
	rb.enabled = enabled
}

func (rb *radioButton) Check(checked bool) {
	rb.checked = checked
	rb.mediator.colleagueChanged()
}

type loginForm struct {
	RadioButton radioButton
	Button      button
}

// NewLoginForm returns a loginForm instance
func NewLoginForm() *loginForm {
	lf := &loginForm{}
	lf.createColleagues()
	return lf
}

func (lf *loginForm) createColleagues() {
	lf.RadioButton = radioButton{}
	lf.Button = button{}
	lf.RadioButton.setMediator(lf)
	lf.Button.setMediator(lf)
}

func (lf *loginForm) colleagueChanged() {
	if !lf.RadioButton.checked {
		lf.Button.setColleagueEnabled(false)
	} else {
		lf.Button.setColleagueEnabled(true)
	}
}
