package dashboard

import (
	"github.com/XuHandsome/stocks/pkgs/config"
	"github.com/XuHandsome/stocks/pkgs/dashboard/gui"
	"github.com/XuHandsome/stocks/pkgs/dashboard/header"
	"github.com/XuHandsome/stocks/pkgs/dashboard/overview"
	"github.com/rivo/tview"
)

type Dashboard struct {
	gui *gui.Gui
}

func New(mainConfig config.MainConfig) *Dashboard {
	return &Dashboard{
		gui: gui.New(mainConfig),
	}
}

func (d *Dashboard) Start() error {
	if err := d.gui.Start(); err != nil {
		return err
	}

	logo := header.NewHeader()
	info := header.NewInfo(d.gui)
	overView := overview.NewOverviewPanel(d.gui)

	d.gui.AddPanels(gui.MainPage, overView)

	grid := tview.NewGrid().SetRows(6, 0, 3).
		SetColumns(0, 2).
		AddItem(logo, 0, 0, 1, 1, 0, 0, false).
		AddItem(info, 0, 2, 1, 1, 0, 0, false).
		AddItem(overView.Table, 1, 0, 1, 3, 0, 0, true).
		AddItem(d.gui.Nav.TextView, 2, 0, 1, 3, 0, 0, false)

	d.gui.Pages = tview.NewPages().
		AddAndSwitchToPage(gui.MainPage, grid, true)

	d.gui.App.SetRoot(d.gui.Pages, true)
	d.gui.SetCurrentPage(gui.MainPage)

	if err := d.gui.App.Run(); err != nil {
		d.gui.App.Stop()
		return err
	}

	return nil
}

func (d *Dashboard) Stop() {}
