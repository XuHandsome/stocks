package gui

import (
	"github.com/gdamore/tcell/v2"
)

func (g *Gui) SetGlobalKeybinding(event *tcell.EventKey) {
	key := event.Rune()
	switch key {
	case 'q':
		g.Stop()
	case 'r':
		g.refresh()
	}
}

func (g *Gui) refresh() {
	page := g.panels.currentPage
	pagePanel, ok := g.panels.panel[page]
	if !ok {
		return
	}

	currentPanel := pagePanel[g.panels.currentPanel]
	currentPanel.UpdateData(g)
}
