package main

import (
	"fmt"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const RowsCount = 12
const ColsCount = 3

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// background window
	bgWindow := widgets.NewQMdiSubWindow(nil, core.Qt__FramelessWindowHint|core.Qt__WindowStaysOnBottomHint)
	// bgWindow.ShowMaximized()

	// table list
	table := widgets.NewQTableWidget2(RowsCount, ColsCount, nil)
	table.SetShowGrid(true)
	table.SetHorizontalHeaderLabels([]string{"ID", "Name", "Date"})
	table.VerticalHeader().SetVisible(false)
	for row := 1; row <= RowsCount; row++ {
		for col := 1; col <= ColsCount; col++ {
			var item *widgets.QTableWidgetItem
			if col == 1 {
				item = widgets.NewQTableWidgetItem2(fmt.Sprintf("%d", row), 0)
			} else if col == 2 {
				item = widgets.NewQTableWidgetItem2("Test", 0)
			} else if col == 3 {
				item = widgets.NewQTableWidgetItem2(time.Now().Format(time.Stamp), 0)
			}

			table.SetItem(row, col, item)
		}
	}

	bgWindow.SetWidget(table)

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
