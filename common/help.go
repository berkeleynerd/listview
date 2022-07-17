package common

import (
	"fmt"
	style "github.com/berkeleynerd/listview/styles"
)

type BubbleHelper interface {
	Help() []HelpEntry
}

type HelpEntry struct {
	Key   string
	Value string
}

func HelpEntryRender(h HelpEntry, s style.Styles) string {
	return fmt.Sprintf("%s %s", s.HelpKey.Render(h.Key), s.HelpValue.Render(h.Value))
}
