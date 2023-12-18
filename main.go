// package main
//
// import "fyne.io/fine/v2/app"
//
// // "fmt"
// // "net/http"
// // "os"
//
// func main() {
// 	a := app.New
// 	w := a.NewWindow("my first app")
//
// 	w.ShowAndRun()
//
// }

package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"

	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello helllllllll World")

	label := widget.NewLabel("hello for fine")
	w.SetContent(container.NewVBox(
		label,
	))
	w.ShowAndRun()
}
