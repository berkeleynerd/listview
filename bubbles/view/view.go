package view

import (
	"github.com/berkeleynerd/listview/bubbles/markdown"
	"github.com/berkeleynerd/listview/common"
	style "github.com/berkeleynerd/listview/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/muesli/reflow/wrap"
)

const (
	documentNameMaxWidth = 32
)

type Model struct {
	name         string
	document     common.Document
	styles       style.Styles
	width        int
	widthMargin  int
	height       int
	heightMargin int
	Active       bool
	box          markdown.Model
}

func New(document common.Document, styles style.Styles, width, widthMargin, height, heightMargin int) Model {
	m := Model{
		name:         document.Name,
		document:     document,
		styles:       styles,
		width:        width,
		widthMargin:  widthMargin,
		height:       height,
		heightMargin: heightMargin,
	}
	m.box = markdown.New(document.Content, styles, width, widthMargin+styles.DocumentBody.GetHorizontalBorderSize(), height, heightMargin+lipgloss.Height(m.headerView())-styles.DocumentBody.GetVerticalBorderSize())
	return m
}

func (m Model) Init() tea.Cmd {
	return m.box.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width == m.width && msg.Height == m.height {
			return m, nil
		}
		m.width = msg.Width
		m.height = msg.Height
	}
	box, cmd := m.box.Update(msg)
	m.box = box.(markdown.Model)
	return m, cmd
}

func (m Model) headerView() string {
	// Render document title
	title := m.name
	title = truncate.StringWithTail(title, documentNameMaxWidth, "…")
	title = m.styles.DocumentTitle.Render(title)

	var note string
	note = m.document.Note
	noteWidth := m.width -
		m.widthMargin -
		lipgloss.Width(title) -
		m.styles.DocumentTitleBox.GetHorizontalFrameSize()
	note = wrap.String(note, noteWidth-m.styles.DocumentNote.GetHorizontalFrameSize())
	note = m.styles.DocumentNote.Copy().Width(noteWidth).Render(note)

	height := max(lipgloss.Height(title), lipgloss.Height(note))
	titleBoxStyle := m.styles.DocumentTitleBox.Copy().Height(height)
	noteBoxStyle := m.styles.DocumentNoteBox.Copy().Height(height)
	if m.Active {
		titleBoxStyle = titleBoxStyle.BorderForeground(m.styles.ActiveBorderColor)
		noteBoxStyle = noteBoxStyle.BorderForeground(m.styles.ActiveBorderColor)
	}
	title = titleBoxStyle.Render(title)
	note = noteBoxStyle.Render(note)

	return lipgloss.JoinHorizontal(lipgloss.Top, title, note)
}

func (m Model) View() string {
	header := m.headerView()
	bs := m.styles.DocumentBody.Copy()
	if m.Active {
		bs = bs.BorderForeground(m.styles.ActiveBorderColor)
	}
	body := bs.Width(m.width - m.widthMargin - m.styles.DocumentBody.GetVerticalFrameSize()).
		Height(m.height - m.heightMargin - lipgloss.Height(header)).
		Render(m.box.View())
	return header + body
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
