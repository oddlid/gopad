package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gameListContainer struct {
	nav         *navigationController
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
				return container.NewBorder(
					nil,
					nil,
					nil,
					container.NewHBox(
						widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil),
						widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
					),
					widget.NewLabel("bajsmacka"),
				)
			},
			func(di binding.DataItem, co fyne.CanvasObject) {
				c1, ok := co.(*fyne.Container)
				if !ok {
					return
				}
				// log.Printf("# obects: %d", len(c.Objects))
				lbl, ok := c1.Objects[0].(*widget.Label)
				if !ok {
					return
				}

				item, ok := di.(binding.String)
				if !ok {
					return
				}
				lbl.Bind(item)

				c2, ok := c1.Objects[1].(*fyne.Container)
				if !ok {
					return
				}

				btnEdit, ok := c2.Objects[0].(*widget.Button)
				if !ok {
					return
				}
				btnEdit.OnTapped = func() {
					glc.nav.load(pIDEditItem)
				}

				btnDelete, ok := c2.Objects[1].(*widget.Button)
				if !ok {
					return
				}
				btnDelete.OnTapped = func() {
					//
				}
			},
		),
	)
}
