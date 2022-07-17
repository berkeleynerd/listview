package markdown

import (
	"github.com/charmbracelet/glamour"
	gansi "github.com/charmbracelet/glamour/ansi"
)

const (
	GlamourMaxWidth = 120
)

func DefaultStyles() gansi.StyleConfig {
	noColor := ""
	s := glamour.DarkStyleConfig
	s.Document.StylePrimitive.Color = &noColor
	s.CodeBlock.Chroma.Text.Color = &noColor
	s.CodeBlock.Chroma.Name.Color = &noColor
	return s
}

func Glamourize(w int, md string) (string, error) {
	if w > GlamourMaxWidth {
		w = GlamourMaxWidth
	}
	tr, err := glamour.NewTermRenderer(
		glamour.WithStyles(DefaultStyles()),
		glamour.WithWordWrap(w),
	)

	if err != nil {
		return "", err
	}
	mdt, err := tr.Render(md)
	if err != nil {
		return "", err
	}
	return mdt, nil
}

func RenderMarkdown(content string, width int) (string, error) {
	md, err := Glamourize(width, content)
	if err != nil {
		return "", err
	}
	return md, nil
}
