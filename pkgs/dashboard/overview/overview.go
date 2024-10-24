package overview

import (
	"fmt"
	"github.com/XuHandsome/stocks/pkgs/config"
	"github.com/XuHandsome/stocks/pkgs/dashboard/gui"
	"github.com/XuHandsome/stocks/pkgs/stocksUntil"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"sync"
	"time"
)

const PanelName = "Stocks overview"

var (
	cellColor = gui.ColorRed
)

type Panel struct {
	gui *gui.Gui
	*tview.Table

	stocks []stocksUntil.Stock
}

func NewOverviewPanel(g *gui.Gui) *Panel {
	l := &Panel{
		gui:   g,
		Table: tview.NewTable().SetBorders(false).SetSelectable(true, false).Select(0, 0).SetFixed(1, 1),
	}

	l.SetTitleAlign(tview.AlignCenter)
	l.SetSelectedStyle(gui.NewTableSelectedStyle(cellColor))
	l.SetBorder(true)
	l.SetData()
	ticker := time.NewTicker(l.gui.Config.Global.UpdateInterval * time.Millisecond)
	go func() {
		for range ticker.C {
			l.UpdateData(l.gui)
		}
	}()
	//l.SetKeybinding(g)
	//g.Nav.AddKeyBindingsNavWithKey(PanelName, "", gui.KeyQ, gui.KeyR)

	return l
}

func (l *Panel) Name() string {
	return PanelName
}

func (l *Panel) queryState(mainConfig config.MainConfig) {
	provider := stocksUntil.SinaStockProvider{}
	var wg sync.WaitGroup
	wg.Add(len(mainConfig.Stocks))

	stocks := make([]stocksUntil.Stock, len(mainConfig.Stocks))

	for i, stock := range mainConfig.Stocks {
		go func(index int, stockCode config.StockInfo) {
			defer wg.Done()
			stockFetch, err := provider.Fetch(stock.Code)
			if err != nil {
				fmt.Println(err)
				return
			}
			stockFetch.HoldPrice = stock.HoldPrice
			stockFetch.HoldNumber = stock.HoldNumber
			stockFetch.Profit = stocksUntil.Calculator(stockFetch.Price, stock.HoldPrice, stock.HoldNumber)
			stocks[index] = stockFetch
		}(i, stock)
	}

	wg.Wait()

	l.stocks = stocks

	count := len(mainConfig.Stocks)
	l.SetTitle(fmt.Sprintf(" Overview | All: %s%d ", gui.ColorTextPurple, count))
}

func (l *Panel) SetData() {
	l.queryState(l.gui.Config)

	renderTable(l.Table, l.stocks)
}

func renderTable(t *tview.Table, stocks []stocksUntil.Stock) {
	table := t.Clear()

	headers := []string{
		"Name", "Code",
		"今开", "昨收",
		"最低", "最高",
		//"百分比", "涨跌",
		"当前",
		"成本", "持仓",
		"市值", "盈亏",
	}

	for i, header := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            header,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorWhite,
			BackgroundColor: tcell.ColorDefault,
			Attributes:      tcell.AttrBold,
		})
	}

	for i, stock := range stocks {
		table.SetCell(i+1, 0, tview.NewTableCell(stock.Name).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 1, tview.NewTableCell(stock.Code).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 2, tview.NewTableCell(strconv.FormatFloat(stock.Open, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 3, tview.NewTableCell(strconv.FormatFloat(stock.YestClose, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 4, tview.NewTableCell(strconv.FormatFloat(stock.Low, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 5, tview.NewTableCell(strconv.FormatFloat(stock.High, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		//table.SetCell(i+1, 2, tview.NewTableCell(strconv.FormatFloat(stock.Percent, 'f', -1, 64)).
		//	SetMaxWidth(0).
		//	SetExpansion(1))
		//
		//table.SetCell(i+1, 3, tview.NewTableCell(strconv.FormatFloat(stock.Updown, 'f', -1, 64)).
		//	SetMaxWidth(0).
		//	SetExpansion(1))

		table.SetCell(i+1, 6, tview.NewTableCell(strconv.FormatFloat(stock.Price, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 7, tview.NewTableCell(strconv.FormatFloat(stock.HoldPrice, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 8, tview.NewTableCell(strconv.Itoa(stock.HoldNumber)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 9, tview.NewTableCell(strconv.FormatFloat(float64(stock.HoldNumber)*stock.Price, 'f', -1, 64)).
			SetMaxWidth(0).
			SetExpansion(1))

		table.SetCell(i+1, 10, tview.NewTableCell(stock.Profit).
			SetMaxWidth(0).
			SetExpansion(1))

	}
}

func (l *Panel) UpdateData(g *gui.Gui) {
	go g.App.QueueUpdateDraw(func() {
		l.SetData()
	})
}

func (l *Panel) Focus() {
	gui.TableFocus(l.Table, l.gui, cellColor)
}

func (l *Panel) UnFocus() {
	gui.TableUnFocus(l.Table)
}

func (l *Panel) SetKeybinding(g *gui.Gui) {
	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		g.SetGlobalKeybinding(event)
		switch event.Key() {
		}
		return event
	})
}
