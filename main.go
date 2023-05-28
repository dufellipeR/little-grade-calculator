package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"little_grade_calculator/screens"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grade Calculator")

	screens.HomeScreen(&myWindow)

	myWindow.Resize(fyne.Size{Width: 1280, Height: 720})
	myWindow.ShowAndRun()
}
