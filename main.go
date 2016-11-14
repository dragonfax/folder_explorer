package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const RowsCount = 12

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// background window
	bgWindow := widgets.NewQMdiSubWindow(nil, core.Qt__FramelessWindowHint|core.Qt__WindowStaysOnBottomHint)
	// bgWindow.ShowMaximized()

	list := widgets.NewQListWidget(nil)
	for i := 1; i <= 3; i++ {
		content := "<p><b>test</b> and</p> something <i>else</i>"
		item := widgets.NewQListWidgetItem2("", nil, 0)
		list.AddItem2(item)
		itemWidget := widgets.NewQLabel2(content, nil, 0)
		itemWidget.SetTextFormat(core.Qt__RichText)
		itemWidget.SetWordWrap(true)
		list.SetItemWidget(item, itemWidget)
	}
	bgWindow.SetWidget(list)

	// child folder window
	cWindow := widgets.NewQMdiSubWindow(nil, 0)
	cWindow.SetMinimumSize2(100, 100)

	// mdi for parent folder
	mdiArea := widgets.NewQMdiArea(nil)
	mdiArea.AddSubWindow(bgWindow, 0)
	mdiArea.AddSubWindow(cWindow, 0)
	bgWindow.Lower()
	cWindow.Raise()

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(mdiArea, 0, core.Qt__AlignCenter)

	//create a window, add the layout and show the window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(600, 600)
	window.Layout().DestroyQObject()
	window.SetLayout(layout)
	window.Show()

	h := mdiArea.Viewport().Height()
	w := mdiArea.Viewport().Width()
	bgWindow.Resize2(w, h)

	widgets.QApplication_Exec()
}
