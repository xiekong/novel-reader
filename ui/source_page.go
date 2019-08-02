/**
 * @author XieKong
 * @date   2019/8/2 16:26
 */
package ui

import (
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"novel-reader/modules"
	"novel-reader/utils"
	"unsafe"
)

func SourceUI() gtk.IWidget {
	ui := gtk.NewVBox(false, 1)

	group := gtk.NewFrame("源")
	ui.PackStart(group, false, false, 0)

	hbox := gtk.NewHBox(false, 1)
	nameLbl := gtk.NewLabel("名称")
	hbox.PackStart(nameLbl, false, false, 0)
	nameTxt := gtk.NewEntry()
	hbox.Add(nameTxt)
	addresLbl := gtk.NewLabel("网址")
	hbox.Add(addresLbl)
	addresTxt := gtk.NewEntry()
	hbox.Add(addresTxt)
	addBtn := gtk.NewButtonWithLabel("新增")
	addBtn.Clicked(func() {
		if len(nameTxt.GetText()) > 0 && len(addresTxt.GetText()) > 0 {
			source := &modules.Source{
				Name: nameTxt.GetText(),
				Address: addresTxt.GetText(),
			}
			if err := modules.InsertSource(source); err != nil {
				utils.Logger.Error(err)
			}
			nameTxt.SetText("")
			addresTxt.SetText("")
		}
	})
	hbox.Add(addBtn)
	group.Add(hbox)

	ui.Add(sourceListView())
	return ui
}

func sourceListView() gtk.IWidget {
	swin := gtk.NewScrolledWindow(nil, nil)
	listView := gtk.NewTreeView()
	swin.Add(listView)
	store := gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_BOOL, gdkpixbuf.GetType())

	listView.SetModel(store)
	listView.AppendColumn(gtk.NewTreeViewColumnWithAttributes("name", gtk.NewCellRendererText(), "text", 0))
	listView.AppendColumn(gtk.NewTreeViewColumnWithAttributes("check", gtk.NewCellRendererToggle(), "active", 1))
	listView.AppendColumn(gtk.NewTreeViewColumnWithAttributes("icon", gtk.NewCellRendererPixbuf(), "pixbuf", 2))
	n := 0
	gtk.StockListIDs().ForEach(func(d unsafe.Pointer, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk.TreeIter
		store.Append(&iter)
		store.Set(&iter,
			0, id,
			1, (n == 1),
			2, gtk.NewImage().RenderIcon(id, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf,
		)
		n = 1 - n
	})
	return swin
}