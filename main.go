package main

import (
	"LangAssistant/api"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Lang Assistant")
	w.Resize(fyne.NewSize(200, 150))

	input := widget.NewEntry()
	result := widget.NewLabel("")
	translate := widget.NewButton("translate", func() {
		text := input.Text
		res, err := api.BaiduTranslate(text)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		result.SetText(res.TransResult[0].Dst)
		fmt.Println(res)
	})

	content := container.NewVBox(
		input,
		result,
		translate,
	)

	w.SetContent(content)

	w.ShowAndRun()
}
