package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var version = "dev"

func main() {
	fyneApp := app.New()
	window := fyneApp.NewWindow("Hello GUI " + version)
	window.Resize(fyne.NewSize(460, 320))

	name := widget.NewEntry()
	name.SetPlaceHolder("Enter your name")
	name.SetText("GitHub")

	result := widget.NewLabel("Press the button to say hello")
	result.Alignment = fyne.TextAlignCenter

	greet := widget.NewButton("Say hello", func() {
		result.SetText(Greeting(name.Text))
	})

	window.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Hello from Go GUI! 🎨🐹", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Version: "+version),
		widget.NewSeparator(),
		widget.NewLabel("Your name"),
		name,
		greet,
		result,
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("Sum 1..10 = %d", SumRange(1, 10))),
		widget.NewButton("Exit", func() { fyneApp.Quit() }),
	))

	window.ShowAndRun()
}
