/**
 * @author XieKong
 * @date   2019/8/2 11:32
 */
package ui

import (
	"github.com/mattn/go-gtk/gtk"
)

type tabPage struct {
	name string
	widget gtk.IWidget
}

func newTabPage(name string, widget gtk.IWidget) *tabPage {
	return &tabPage{
		name: name,
		widget: widget,
	}
}

type TabPanel struct {
	notebook *gtk.Notebook
	activatePage *tabPage
	pages []*tabPage
}

func NewTabPanel() *TabPanel {
	return &TabPanel{
		notebook: gtk.NewNotebook(),
		pages: make([]*tabPage, 0),
	}
}

func (t *TabPanel) ExistsTabPage(name string) (*tabPage, int) {
	if len(name) >= 0 {
		for _, tg := range t.pages {
			if tg.name == name {
				return tg, t.notebook.PageNum(tg.widget)
			}
		}
	}
	return nil, -1
}

func (t *TabPanel) AddPage(widget gtk.IWidget, name string) {
	exists, index := t.ExistsTabPage(name)
	if exists == nil {
		t.pages = append(t.pages, newTabPage(name, widget))
		i := t.notebook.AppendPage(widget, gtk.NewLabel(name));
		t.notebook.ShowAll()
		t.notebook.SetCurrentPage(i)
	} else {
		t.notebook.SetCurrentPage(index)
	}
}

func (t *TabPanel) RemovePage(name string) {
	if len(name) > 0 {
		exists, num := t.ExistsTabPage(name)
		if exists != nil {
			index := -1;
			for i, tg := range t.pages {
				if tg.name == name {
					index = i;
					break
				}
			}
			if index > -1 {
				t.pages = append(t.pages[:index], t.pages[index+1:]...)
			}
			t.notebook.RemovePage(exists.widget, num)
		}
	}
}

func (t *TabPanel) Notebook() *gtk.Notebook {
	return t.notebook
}

func (t *TabPanel) Pages() []*tabPage {
	return t.pages
}