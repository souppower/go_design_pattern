package bridge

type DisplayImpl interface {
	rawOpen() string
	rawPrint() string
	rawClose() string
}

type StringDisplayImpl struct {
	str string
}

func (sdi *StringDisplayImpl) rawOpen() string {
	return sdi.printLine()
}

func (sdi *StringDisplayImpl) rawPrint() string {
	return "|" + sdi.str + "|\n"
}

func (sdi *StringDisplayImpl) rawClose() string {
	return sdi.printLine()
}

func (sdi *StringDisplayImpl) printLine() string {
	str := "+"
	for _, _ = range sdi.str {
		str += string("-")
	}
	str += "+\n"
	return str
}

type DefaultDisplay struct {
	impl DisplayImpl
}

func (dd *DefaultDisplay) open() string {
	return dd.impl.rawOpen()
}

func (dd *DefaultDisplay) print() string {
	return dd.impl.rawPrint()
}

func (dd *DefaultDisplay) close() string {
	return dd.impl.rawClose()
}

func (dd *DefaultDisplay) Display() string {
	return dd.open() +
		dd.print() +
		dd.close()
}

type CountDisplay struct {
	*DefaultDisplay
}

func (cd *CountDisplay) MultiDisplay(num int) string {
	str := cd.open()
	for i := 0; i < num; i++ {
		str += cd.print()
	}
	str += cd.close()
	return str
}
