package gui

import (
	"github.com/XuHandsome/stocks/pkgs/config"
	"github.com/XuHandsome/stocks/pkgs/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

const (
	MainPage = "main"
)

// Considering that different terminals support limited colors, only the most common colors are used here
const (
	ColorSelectedForeground = tcell.ColorBlack
	ColorTeal               = tcell.ColorTeal
	ColorWhite              = tcell.ColorWhite
	ColorRed                = tcell.ColorRed
	ColorYellow             = tcell.ColorYellow

	ColorTextWhite  = "[white]"
	ColorTextPurple = "[purple]"
	ColorTextTeal   = "[teal]"
)

type Gui struct {
	App   *tview.Application
	Pages *tview.Pages

	panels *Panels
	Nav    *Navigate
	Config config.MainConfig

	GlobalStat Stat
}

type Stat struct {
	StocksVersion  string
	UpdateInterval time.Duration
}

func New(mainConfig config.MainConfig) *Gui {
	return &Gui{
		App:    tview.NewApplication(),
		panels: newPanels(),
		Nav:    newNavigate(),
		Config: mainConfig,
	}
}

func (g *Gui) Start() error {
	g.GlobalStat.StocksVersion = core.GetVersion()
	g.GlobalStat.UpdateInterval = g.Config.Global.UpdateInterval * time.Millisecond
	return nil
}

func (g *Gui) Stop() {}

func NewTableSelectedStyle(color tcell.Color) tcell.Style {
	return tcell.StyleDefault.Background(color).Foreground(ColorSelectedForeground)
}

func TableFocus(table *tview.Table, gui *Gui, color tcell.Color) {
	table.SetSelectable(true, false)
	table.SetBorderColor(color)
	table.SetTitleColor(color)
	gui.App.SetFocus(table)
}

func TableUnFocus(table *tview.Table) {
	table.SetSelectable(false, false)
	table.SetBorderColor(ColorWhite)
	table.SetTitleColor(ColorWhite)
}
