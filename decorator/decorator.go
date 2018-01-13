package decorator

type display interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
	Show(display display) string
}

type defaultDisplay struct {
	display
}

func (dd *defaultDisplay) Show(display display) string {
	var str string
	for i := 0; i < display.getRows(); i++ {
		str += display.getRowText(i)
	}
	return str
}

type StringDisplay struct {
	*defaultDisplay
	str string
}

// NewStringDisplay returns a new StringDisplay instance
func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		&defaultDisplay{},
		str,
	}
}

func (sd *StringDisplay) getColumns() int {
	return len(sd.str)
}

func (sd *StringDisplay) getRows() int {
	return 1
}

func (sd *StringDisplay) getRowText(row int) string {
	if row == 0 {
		return sd.str
	}

	return ""
}

type border struct {
	*defaultDisplay
	display display
}

type SideBorder struct {
	*border
	borderChar string
}

func NewSideBorder(display display, borderChar string) *SideBorder {
	return &SideBorder{
		&border{display: display},
		borderChar,
	}
}

func (sb *SideBorder) getColumns() int {
	return len(sb.borderChar)*2 + sb.display.getColumns()
}

func (sb *SideBorder) getRows() int {
	return sb.display.getRows()
}

func (sb *SideBorder) getRowText(row int) string {
	return sb.borderChar + sb.display.getRowText(row) + sb.borderChar
}
