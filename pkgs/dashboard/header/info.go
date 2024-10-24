package header

import (
	"fmt"
	"github.com/XuHandsome/stocks/pkgs/dashboard/gui"
	"github.com/rivo/tview"
	"strings"
	"time"
)

type Info struct {
	gui *gui.Gui

	*tview.TextView
	text           string
	lastUpdateTime string
}

func NewInfo(g *gui.Gui) *Info {
	h := &Info{
		gui:      g,
		TextView: tview.NewTextView(),
	}

	h.setData()

	ticker := time.NewTicker(g.GlobalStat.UpdateInterval)
	go func() {
		for range ticker.C {
			h.UpdateData(h.gui)
		}
	}()

	return h
}

func (h *Info) queryState() {
	var out strings.Builder

	h.lastUpdateTime = time.Now().Format("2006-01-02 15:04:05")

	stat := h.gui.GlobalStat

	out.WriteString("\n\n\n")
	space := "                                "
	out.WriteString(fmt.Sprintf(space+"[red]当前版本: [white]%s\n", stat.StocksVersion))
	out.WriteString(fmt.Sprintf(space+"[red]刷新间隔: [white]%s\n", stat.UpdateInterval))
	//out.WriteString(fmt.Sprintf(space+"[red]运行时间: [white]%s\n", h.lastUpdateTime))
	h.text = h.text + out.String()
}

func (h *Info) setData() {
	h.queryState()
	h.TextView.Clear().SetDynamicColors(true)
	h.TextView.SetText(h.text)
	//getText := "本次从Tview中获取到的内容：" + h.TextView.GetText(true)
	//ioutil.WriteFile("./log.text", []byte(getText), 0644)
}

func (h *Info) UpdateData(g *gui.Gui) {
	go g.App.QueueUpdateDraw(func() {
		h.setData()
	})
}
