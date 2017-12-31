package adapter

// Decorator interface
type Decorator interface {
	Decorate() string
}

// Banner struct
type Banner struct {
	str string
}

func (b *Banner) getString() string {
	return "*" + b.str + "*"
}

// EmbeddedDecorateBanner struct
type EmbeddedDecorateBanner struct {
	*Banner
}

// NewEmbeddedDecorateBanner returns an EmbeddedDecorateBanner instance
func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	return &EmbeddedDecorateBanner{&Banner{str}}
}

// Decorate is an adapter method
func (edc *EmbeddedDecorateBanner) Decorate() string {
	return edc.getString()
}
