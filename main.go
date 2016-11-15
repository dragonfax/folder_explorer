package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var htmlDelegate *HTMLDelegate

type HTMLDelegate struct {
	widgets.QStyledItemDelegate
}

func InitDelegate() *HTMLDelegate {
	item := NewHTMLDelegate(nil)

	item.ConnectPaint(paint)
	item.ConnectSizeHint(sizeHint)

	return item
}

func paint(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) {

	fmt.Printf("started painting index %d\n", index.Row())

	options := widgets.NewQStyleOptionViewItem2(option)
	htmlDelegate.InitStyleOption(options, index)

	painter.Save()

	doc := gui.NewQTextDocument(nil)
	text := options.Text()
	doc.SetHtml(text)

	options.SetText("")
	options.Widget().Style().DrawControl(widgets.QStyle__CE_ItemViewItem, options, painter, nil)

	painter.Translate(core.NewQPointF2(options.Rect().TopLeft()))
	clip := core.NewQRectF4(0, 0, float64(options.Rect().Width()), float64(options.Rect().Height()))
	doc.DrawContents(painter, clip)

	painter.Restore()

	fmt.Printf("drew '%s'\n", text)
}

func sizeHint(option *widgets.QStyleOptionViewItem, index *core.QModelIndex) *core.QSize {
	fmt.Println("started sizing")

	options := widgets.NewQStyleOptionViewItem2(option)
	htmlDelegate.InitStyleOption(options, index)

	var doc gui.QTextDocument
	doc.SetHtml(options.Text())
	doc.SetTextWidth(float64(options.Rect().Width()))

	w := int(doc.IdealWidth())
	h := int(doc.Size().Height())

	fmt.Println("finished sizing")
	return core.NewQSize2(w, h)
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// background window
	bgWindow := widgets.NewQMdiSubWindow(nil, core.Qt__FramelessWindowHint|core.Qt__WindowStaysOnBottomHint)
	// bgWindow.ShowMaximized()

	list := widgets.NewQListView(nil)
	listModel := core.NewQStringListModel2([]string{"<p>test</p>3 on a new line", "<b>test 1</b>", "<br/><br/> another text", "<i>test 2</i>"}, nil)
	if listModel.RowCount(core.NewQModelIndex()) != 4 {
		panic("wrong number of items in list")
	}
	htmlDelegate = InitDelegate()
	list.SetModel(listModel)
	list.SetItemDelegate(htmlDelegate)
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
