package header

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

type Header struct {
	*tview.TextView
}

var logo = []string{
	`_________ __                 __                    /   ^__^                /`,
	`/   _____//  |_  ____   ____ |  | __  ______      /    (oo)\_________     /`,
	`\_____  \\   __\/  _ \_/ ___\|  |/ / /  ___/           (__)\        ) \  /`,
	`/        \|  | (  <_> )  \___|    <  \___ \                ||----w ||  \/`,
	`/_______  /|__|  \____/ \___  >__|_ \/____  >              ||      ||`,
}

func NewHeader() *Header {
	h := &Header{
		TextView: tview.NewTextView().SetDynamicColors(true),
	}

	h.display()

	return h
}

func (h *Header) display() {
	h.SetTextColor(tcell.ColorRed)

	var out strings.Builder
	for _, l := range logo {
		out.WriteString(l)
		out.WriteString("\n")
	}
	h.SetText(out.String())
}
