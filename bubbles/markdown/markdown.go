package markdown

import (
	"github.com/berkeleynerd/listview/styles"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wrap"
)

type Model struct {
	viewport     viewport.Model
	styles       style.Styles
	height       int
	heightMargin int
	width        int
	widthMargin  int
	content      string
}

func New(content string, styles style.Styles, width, widthMargin, height, heightMargin int) Model {
	heightMargin = heightMargin + lipgloss.Height("") // Prevents screen shake
	b := Model{
		viewport:     viewport.Model{},
		styles:       styles,
		widthMargin:  widthMargin,
		heightMargin: heightMargin,
		content:      content,
	}
	b.SetSize(width, height)
	return b
}
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		md := m.glamourize()
		m.viewport.SetContent(md)
	}
	rv, cmd := m.viewport.Update(msg)
	m.viewport = rv
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
	m.viewport.Width = w - m.widthMargin
	m.viewport.Height = h - m.heightMargin
}

func (m *Model) GotoTop() {
	m.viewport.GotoTop()
}

func (m Model) View() string {
	return m.viewport.View()
}

func (m Model) glamourize() string {
	w := m.width - m.widthMargin - m.styles.DocumentBody.GetHorizontalFrameSize()
	rm := m.content
	f, _ := RenderMarkdown(rm, w)
	return wrap.String(f, w)
}
