package listview

import (
	"fmt"
	"github.com/berkeleynerd/listview/bubbles/selection"
	"github.com/berkeleynerd/listview/bubbles/view"
	"github.com/berkeleynerd/listview/common"
	style "github.com/berkeleynerd/listview/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type ActiveBox uint8

const (
	SELECTION ActiveBox = 0
	DOCUMENT  ActiveBox = 1
)

type MenuEntry struct {
	Name     string
	Document string
	bubble   view.Model
}

type Model struct {
	title     string
	width     int
	height    int
	boxes     []tea.Model
	activeBox ActiveBox
	styles    style.Styles

	documentMenu   []MenuEntry
	documentSelect selection.Model

	// save the last resize message to re-send when a different document is selected
	lastResize tea.WindowSizeMsg
}

func New(title string, documents []common.Document) Model {
	boxes := make([]tea.Model, 2)
	m := Model{title: title, boxes: boxes, activeBox: SELECTION, styles: style.DefaultStyles()}

	mes := m.menuEntriesFromList(documents)
	m.documentMenu = mes
	rs := make([]string, 0)
	for _, m := range mes {
		rs = append(rs, m.Name)
	}
	s := selection.New(rs, m.styles)
	m.documentSelect = s
	m.boxes[SELECTION] = m.documentSelect

	m.boxes[DOCUMENT] = m.documentMenu[0].bubble
	m.documentSelect.SelectedItem = 0
	m.activeBox = 0

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", "tab", "shift+tab":
			m.activeBox = (m.activeBox + 1) % 2
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.lastResize = msg
		m.width = msg.Width
		m.height = msg.Height
		for i, bx := range m.boxes {
			mx, cmd := bx.Update(msg)
			m.boxes[i] = mx
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	case selection.ActiveMsg:
		m.boxes[DOCUMENT] = m.documentMenu[msg.Index].bubble
		cmds = append(cmds, func() tea.Msg {
			return m.lastResize
		})
	}

	ab, cmd := m.boxes[m.activeBox].Update(msg)
	m.boxes[m.activeBox] = ab
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) viewForBox(i int) string {
	isActive := i == int(m.activeBox)
	switch box := m.boxes[i].(type) {
	case selection.Model:
		// List
		var s lipgloss.Style
		s = m.styles.Menu
		if isActive {
			s = s.Copy().BorderForeground(m.styles.ActiveBorderColor)
		}
		return s.Render(box.View())
	case view.Model:
		// Document
		box.Active = isActive
		return box.View()
	default:
		panic(fmt.Sprintf("unknown box type %T", box))
	}
}

func (m Model) headerView() string {
	w := m.width - m.styles.App.GetHorizontalFrameSize()
	return m.styles.Header.Copy().Width(w).Render(m.title)
}

func (m Model) footerView() string {
	w := &strings.Builder{}
	var h []common.HelpEntry

	h = []common.HelpEntry{
		{Key: "tab", Value: "section"},
	}
	if box, ok := m.boxes[m.activeBox].(common.BubbleHelper); ok {
		help := box.Help()
		for _, he := range help {
			h = append(h, he)
		}
	}

	h = append(h, common.HelpEntry{Key: "q", Value: "quit"})
	for i, v := range h {
		fmt.Fprint(w, common.HelpEntryRender(v, m.styles))
		if i != len(h)-1 {
			fmt.Fprint(w, m.styles.HelpDivider)
		}
	}
	help := w.String()
	gap := lipgloss.NewStyle().
		Width(m.width -
			lipgloss.Width(help) -
			m.styles.App.GetHorizontalFrameSize()).
		Render("")
	footer := lipgloss.JoinHorizontal(lipgloss.Top, help, gap)
	return m.styles.Footer.Render(footer)
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.headerView())
	s.WriteRune('\n')
	sb := m.viewForBox(0)
	db := m.viewForBox(1)
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, sb, db))
	s.WriteRune('\n')
	s.WriteString(m.footerView())
	return m.styles.App.Render(s.String())
}

func (m Model) menuEntriesFromList(documents []common.Document) []MenuEntry {
	mes := make([]MenuEntry, 0)
	for i, cr := range documents {
		me := MenuEntry{Name: cr.Name, Document: cr.Content}
		r := documents[i]
		boxLeftWidth := m.styles.Menu.GetWidth() + m.styles.Menu.GetHorizontalFrameSize()
		// TODO: also send this along with a tea.WindowSizeMsg
		var heightMargin = lipgloss.Height(m.headerView()) +
			lipgloss.Height(m.footerView()) +
			m.styles.DocumentBody.GetVerticalFrameSize() +
			m.styles.App.GetVerticalMargins()
		rb := view.New(r, m.styles, m.width, boxLeftWidth, m.height, heightMargin)
		rb.Init()
		me.bubble = rb
		mes = append(mes, me)
	}
	return mes
}
