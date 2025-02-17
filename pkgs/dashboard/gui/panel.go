package gui

type Panels struct {
	currentPanel int
	currentPage  string
	panel        map[string][]Panel // key: page name, value: panel list
}

func newPanels() *Panels {
	return &Panels{
		panel: make(map[string][]Panel),
	}
}

type Panel interface {
	Name() string
	SetData()
	UpdateData(*Gui)
	SetKeybinding(*Gui)
	Focus()
	UnFocus()
}

func (g *Gui) AddPanels(page string, pa ...Panel) {
	panels, ok := g.panels.panel[page]
	if !ok {
		g.panels.panel[page] = pa
		return
	}

	panels = append(panels, pa...)
	g.panels.panel[page] = panels
	return
}

func (g *Gui) SetCurrentPage(page string) {
	g.panels.currentPage = page
	g.panels.currentPanel = 0
}

func (g *Gui) NextPanel() {
	page := g.panels.currentPage
	pagePanel, ok := g.panels.panel[page]
	if !ok {
		return
	}

	idx := (g.panels.currentPanel + 1) % len(pagePanel)
	g.SwitchPanel(pagePanel[idx].Name())
}

func (g *Gui) PrevPanel() {
	page := g.panels.currentPage
	pagePanel, ok := g.panels.panel[page]
	if !ok {
		return
	}

	g.panels.currentPanel--

	if g.panels.currentPanel < 0 {
		g.panels.currentPanel = len(pagePanel) - 1
	}

	idx := (g.panels.currentPanel) % len(pagePanel)
	g.SwitchPanel(pagePanel[idx].Name())
}

func (g *Gui) CurrentPanel() Panel {
	page := g.panels.currentPage
	pagePanel, ok := g.panels.panel[page]
	if !ok {
		return nil
	}

	return pagePanel[g.panels.currentPanel]
}

func (g *Gui) SwitchPanel(panelName string) {
	page := g.panels.currentPage
	pagePanel, ok := g.panels.panel[page]
	if !ok {
		return
	}

	for i, panel := range pagePanel {
		if panel.Name() == panelName {
			panel.Focus()
			g.Nav.update(panelName)
			g.panels.currentPanel = i
		} else {
			panel.UnFocus()
		}
	}
}

func (g *Gui) CloseAndSwitchPanel(removePage string, lastPage string, switchPanel string) {
	g.Pages.RemovePage(removePage).ShowPage(lastPage)
	g.SetCurrentPage(lastPage)
	g.SwitchPanel(switchPanel)
}
