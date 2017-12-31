package builder

type builder interface {
	makeTitle(title string) string
	makeString(str string) string
	makeItems(items []string) string
	close() string
}

// Director contains a builder instance
type Director struct {
	builder builder
}

// Construct does all the heavy lifting
func (d *Director) Construct() string {
	result := d.builder.makeTitle("Title")
	result += d.builder.makeString("String")
	result += d.builder.makeItems([]string{
		"Item1",
		"Item2",
	})
	result += d.builder.close()
	return result
}

// TextBuilder implemets the builder interface
type TextBuilder struct {
}

func (tb *TextBuilder) makeTitle(title string) string {
	return "# " + title + "\n"
}

func (tb *TextBuilder) makeString(str string) string {
	return "## " + str + "\n"
}

func (tb *TextBuilder) makeItems(items []string) string {
	var result string
	for _, item := range items {
		result += "- " + item + "\n"
	}
	return result
}

func (tb *TextBuilder) close() string {
	return "\n"
}
