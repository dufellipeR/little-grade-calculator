package screens

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"little_grade_calculator/calculator"
)

func LoadResultScreen(input string, window *fyne.Window) {
	var correct int16
	var wrong int16
	var grade float32
	var total int16
	var msg string

	correct, wrong, total, grade = calculator.CalculateGrade(input)

	if grade >= 80 {
		msg = "YOU'RE APPROVED ! NICE ONE!"
	} else {
		msg = "TRY AGAIN !"
	}

	gradeText := canvas.NewText(fmt.Sprintf("%.2f%%", grade), color.White)
	gradeBox := container.NewHBox(layout.NewSpacer(), gradeText, layout.NewSpacer())

	text := canvas.NewText(msg, color.White)
	msgBox := container.NewHBox(layout.NewSpacer(), text, layout.NewSpacer())

	statsBox := container.NewHBox(
		layout.NewSpacer(),
		container.NewVBox(
			canvas.NewText(fmt.Sprintf("YOU TACKLED %d EXERCISES", total), color.White),
			canvas.NewText(fmt.Sprintf("YOU GOT %d ANSWERS RIGHT", correct), color.White),
			canvas.NewText(fmt.Sprintf("YOU MISSED %d ANSWERS", wrong), color.White),
		),
		layout.NewSpacer())

	button := widget.NewButton("BACK", func() {
		HomeScreen(window)
	})

	resultBox := container.NewVBox(layout.NewSpacer(), msgBox, gradeBox, statsBox, button, layout.NewSpacer())

	resultBox.Add(button)
	(*window).SetContent(resultBox)
}
