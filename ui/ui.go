/**
 * @author XieKong
 * @date   2019/7/31 10:29
 */
package ui

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

const (
	title = "测试"
	width = 550
	height = 600
)

var (
	mainWin *gtk.Window
	tabPanel *TabPanel
)

func Show() {
	// 环境初始化
	gtk.Init(&os.Args)
	// 创建主窗口
	mainUI()
	// 主事件循环，等待用户操作
	gtk.Main()
}

// 主界面
func mainUI() {
	//runtime.GOMAXPROCS(10)
	//glib.ThreadInit(nil)
	//gdk.ThreadsInit()
	//gdk.ThreadsEnter()
	// 创建窗口
	mainWin = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	// 设置窗口居中显示
	mainWin.SetPosition(gtk.WIN_POS_CENTER)
	// 设置标题
	mainWin.SetTitle(title)
	// 设置窗口的宽度和高度
	mainWin.SetSizeRequest(width, height)

	vbox := gtk.NewVBox(false, 1)
	menu := menu()
	// 创建菜单
	vbox.PackStart(menu, false, false, 0)
	// 创建tab
	tabPanel = NewTabPanel()
	vbox.Add(tabPanel.Notebook())
	mainWin.Add(vbox)

	mainWin.Connect("destroy", func(ctx *glib.CallbackContext) {
		//mainWin.Destroy()
		gtk.MainQuit()
	})
	// 显示
	mainWin.ShowAll()
}

// 菜单
func menu() *gtk.MenuBar {
	menubar := gtk.NewMenuBar()

	menu := gtk.NewMenu()
	//myCenter := gtk.NewMenuItemWithMnemonic("我的")
	//menubar.Append(myCenter)
	//myCenter.SetSubmenu(menu);
	//
	//bookMenuItem := gtk.NewMenuItemWithMnemonic("书架")
	//bookMenuItem.Connect("activate", func() {
	//	page := gtk.NewFrame("书架")
	//	tabPanel.AddPage(page, "书架")
	//})
	//menu.Append(bookMenuItem)
	//
	//menu = gtk.NewMenu()
	settings := gtk.NewMenuItemWithMnemonic("设置")
	menubar.Append(settings)
	settings.SetSubmenu(menu);

	sourceMenuItem := gtk.NewMenuItemWithMnemonic("源设置")
	sourceMenuItem.Connect("activate", func() {
		sourcePage := SourceUI()
		tabPanel.AddPage(sourcePage, "源设置")
	})
	menu.Append(sourceMenuItem)

	return menubar
}