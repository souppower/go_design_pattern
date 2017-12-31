package adapter

// ComposiotionDecorateBanner holds a banner struct as a field
type ComposiotionDecorateBanner struct {
	banner *Banner
}

// NewCompositionDecorateBanner returns a ComposiotionDecorateBanner instance
func NewCompositionDecorateBanner(str string) *ComposiotionDecorateBanner {
	return &ComposiotionDecorateBanner{&Banner{str}}
}

// Decorate is an adapter method
func (cdb *ComposiotionDecorateBanner) Decorate() string {
	return cdb.banner.getString()
}
