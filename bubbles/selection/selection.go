package selection

import (
	"strings"

	"github.com/berkeleynerd/listview/common"
	"github.com/berkeleynerd/listview/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

type ActiveMsg struct {
	Name  string
	Index int
}

type Model struct {
	Items        []string
	SelectedItem int
	styles       style.Styles
}

func New(items []string, styles style.Styles) Model {
	return Model{
		Items:  items,
		styles: styles,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	s := strings.Builder{}
	documentNameMaxWidth := m.styles.Menu.GetWidth() - // menu width
		m.styles.Menu.GetHorizontalPadding() - // menu padding
		lipgloss.Width(m.styles.MenuCursor.String()) - // cursor
		m.styles.MenuItem.GetHorizontalFrameSize() // menu item gaps
	for i, item := range m.Items {
		item := truncate.StringWithTail(item, uint(documentNameMaxWidth), "…")
		if i == m.SelectedItem {
			s.WriteString(m.styles.MenuCursor.String())
			s.WriteString(m.styles.SelectedMenuItem.Render(item))
		} else {
			s.WriteString(m.styles.MenuItem.Render(item))
		}
		if i < len(m.Items)-1 {
			s.WriteRune('\n')
		}
	}
	return s.String()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up":
			if m.SelectedItem > 0 {
				m.SelectedItem--
				cmds = append(cmds, m.sendActiveMessage)
			}
		case "j", "down":
			if m.SelectedItem < len(m.Items)-1 {
				m.SelectedItem++
				cmds = append(cmds, m.sendActiveMessage)
			}
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) Help() []common.HelpEntry {
	return []common.HelpEntry{
		{Key: "↑/↓", Value: "navigate"},
	}
}

func (m Model) sendActiveMessage() tea.Msg {
	if m.SelectedItem >= 0 && m.SelectedItem < len(m.Items) {
		return ActiveMsg{
			Name:  m.Items[m.SelectedItem],
			Index: m.SelectedItem,
		}
	}
	return nil
}
