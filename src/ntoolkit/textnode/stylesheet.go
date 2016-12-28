package textnode

const (
	Normal = 0
	Underline = 1 << 0
	Bold = 1 << 1
	Italic = 1 << 2
	StrikeThrough = 1 << 3
)

// Standard colors
var White Color = Color{1.0, 1.0, 1.0, 1.0}
var Black Color = Color{0.0, 0.0, 0.0, 1.0}

type StyleSheet struct {
	Default Style
	styles  map[string]*Style
}

type Style struct {
	Foreground Color
	Background Color
	Decoration int
}

type Color struct {
	red   float32
	green float32
	blue  float32
	alpha float32
}

func NewStyleSheet() *StyleSheet {
	return &StyleSheet{
		Default: Style{
			Foreground: White,
			Background: Black,
			Decoration: Normal},
		styles: make(map[string]*Style)}
}

// Clone a style
func (s *Style) Clone() *Style {
	return &Style{
		Foreground: s.Foreground,
		Background: s.Background,
		Decoration: s.Decoration}
}

// Add a new style to the style sheet
func (s *StyleSheet) New(id string) *Style {
	style := s.Default.Clone()
	s.styles[id] = style
	return style
}

func (s *StyleSheet) Set(id string, style *Style) {
	s.styles[id] = style
}

func (s *StyleSheet) Get(id string) *Style {
	if value, ok := s.styles[id]; ok {
		return value
	}
	return nil
}