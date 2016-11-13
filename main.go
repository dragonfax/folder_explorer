package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	//create a button and connect the clicked signal
	/*var button = widgets.NewQPushButton2("Click me!", nil)
	button.ConnectClicked(func(flag bool) {
		widgets.QMessageBox_Information(nil, "OK", "You clicked me!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	}) */

	mdiArea2 := widgets.NewQMdiArea(nil)
	mdiWindow2 := widgets.NewQMdiSubWindow(nil, 0)
	mdiWindow2.SetMinimumSize2(100, 100)
	mdiArea2.AddSubWindow(mdiWindow2, 0)

	mdiArea := widgets.NewQMdiArea(nil)
	mdiWindow := widgets.NewQMdiSubWindow(nil, 0)
	mdiWindow.SetMinimumSize2(300, 300)
	mdiWindow.SetWidget(mdiArea2)
	mdiArea.AddSubWindow(mdiWindow, 0)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(mdiArea, 0, core.Qt__AlignCenter)

	//create a window, add the layout and show the window
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(600, 600)
	window.Layout().DestroyQObject()
	window.SetLayout(layout)
	window.Show()

	widgets.QApplication_Exec()
}
