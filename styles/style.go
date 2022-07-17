package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles defines styles for the TUI.
type Styles struct {
	ActiveBorderColor   lipgloss.Color
	InactiveBorderColor lipgloss.Color

	App    lipgloss.Style
	Header lipgloss.Style

	Menu             lipgloss.Style
	MenuCursor       lipgloss.Style
	MenuItem         lipgloss.Style
	SelectedMenuItem lipgloss.Style

	DocumentTitleBorder lipgloss.Border
	DocumentNoteBorder  lipgloss.Border
	DocumentBodyBorder  lipgloss.Border

	DocumentTitle    lipgloss.Style
	DocumentTitleBox lipgloss.Style
	DocumentNote     lipgloss.Style
	DocumentNoteBox  lipgloss.Style
	DocumentBody     lipgloss.Style

	Footer      lipgloss.Style
	Branch      lipgloss.Style
	HelpKey     lipgloss.Style
	HelpValue   lipgloss.Style
	HelpDivider lipgloss.Style
}

// DefaultStyles returns default styles for the TUI.
func DefaultStyles() Styles {
	var s Styles

	s.ActiveBorderColor = lipgloss.Color("62")
	s.InactiveBorderColor = lipgloss.Color("236")

	s.App = lipgloss.NewStyle().
		Margin(1, 2)

	s.Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("62")).
		Align(lipgloss.Right).
		Bold(true)

	s.Menu = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(s.InactiveBorderColor).
		Padding(1, 2).
		MarginRight(1).
		Width(24)

	s.MenuCursor = lipgloss.NewStyle().
		Foreground(lipgloss.Color("213")).
		SetString(">")

	s.MenuItem = lipgloss.NewStyle().
		PaddingLeft(2)

	s.SelectedMenuItem = lipgloss.NewStyle().
		Foreground(lipgloss.Color("207")).
		PaddingLeft(1)

	s.DocumentTitleBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "┬",
		BottomLeft:  "├",
		BottomRight: "┴",
	}

	s.DocumentNoteBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┬",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┤",
	}

	s.DocumentBodyBorder = lipgloss.Border{
		Top:         "",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "",
		TopRight:    "",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	s.DocumentTitle = lipgloss.NewStyle().
		Padding(0, 2)

	s.DocumentTitleBox = lipgloss.NewStyle().
		BorderStyle(s.DocumentTitleBorder).
		BorderForeground(s.InactiveBorderColor)

	s.DocumentNote = lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(lipgloss.Color("168"))

	s.DocumentNoteBox = lipgloss.NewStyle().
		BorderStyle(s.DocumentNoteBorder).
		BorderForeground(s.InactiveBorderColor).
		BorderTop(true).
		BorderRight(true).
		BorderBottom(true).
		BorderLeft(false)

	s.DocumentBody = lipgloss.NewStyle().
		BorderStyle(s.DocumentBodyBorder).
		BorderForeground(s.InactiveBorderColor).
		PaddingRight(1)

	s.Footer = lipgloss.NewStyle().
		MarginTop(1)

	s.Branch = lipgloss.NewStyle().
		Foreground(lipgloss.Color("203")).
		Background(lipgloss.Color("236")).
		Padding(0, 1)

	s.HelpKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	s.HelpValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color("239"))

	s.HelpDivider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("237")).
		SetString(" • ")

	return s
}
