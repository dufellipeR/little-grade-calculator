package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func HomeScreen(window *fyne.Window) {

	text1 := canvas.NewText("Instructions: ", color.White)
	text1.TextStyle = fyne.TextStyle{Bold: true}
	text2 := canvas.NewText("Enter \"t\" for right answer and \"f\" for wrong answer, then press \"Calculate\" ", color.White)

	input := widget.NewEntry()
	input.SetPlaceHolder("tftfftt...")
	button := widget.NewButton("Calculate", func() {
		if input.Text == "" {
			LoadResultScreen("f", window)
		} else {
			LoadResultScreen(input.Text, window)
		}

	})

	horizontal := container.NewHBox(layout.NewSpacer(), text1, text2, layout.NewSpacer())
	content := container.NewVBox(horizontal, layout.NewSpacer(), input, button, layout.NewSpacer())

	(*window).SetContent(content)
}
