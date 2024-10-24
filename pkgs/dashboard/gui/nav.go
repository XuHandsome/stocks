package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

const (
	KeyEnter = "enter"
	KeyQ     = "quit"
	KeyF     = "filter"
	KeyR     = "refresh"
)

var defaultKeyBinding = make(map[string]string)

type Navigate struct {
	*tview.TextView
	keybindings map[string]string
}

func newNavigate() *Navigate {
	nav := &Navigate{
		TextView:    tview.NewTextView().SetTextColor(tcell.ColorYellow).SetDynamicColors(true),
		keybindings: make(map[string]string),
	}
	return nav
}

func (n *Navigate) AddKeyBindingsNavWithKey(panel string, extra string, keys ...string) {
	var info strings.Builder
	for _, k := range keys {
		val, ok := defaultKeyBinding[k]
		if !ok {
			continue
		}

		info.WriteString(val)
		info.WriteString("  ")
	}

	info.WriteString("\n")
	info.WriteString(extra)
	n.keybindings[panel] = info.String()
}

func (n *Navigate) update(panel string) {
	n.SetText(n.keybindings[panel])
}
