package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type gameListContainer struct {
	headerLabel string
	buttonLabel string
	data        binding.ExternalStringList
	onAdd       func()
	// onSelect    func()
}

func (glc *gameListContainer) render() *fyne.Container {
	return container.NewBorder(
		widget.NewLabel(glc.headerLabel),
		widget.NewButton(glc.buttonLabel, glc.onAdd),
		nil,
		nil,
		widget.NewListWithData(
			glc.data,
			func() fyne.CanvasObject {
				return widget.NewLabel("bajsmacka")
			},
			func(di binding.DataItem, co fyne.CanvasObject) {
				lbl, ok := co.(*widget.Label)
				if !ok {
					return
				}
				item, ok := di.(binding.String)
				if !ok {
					return
				}
				lbl.Bind(item)
			},
		),
	)
}
