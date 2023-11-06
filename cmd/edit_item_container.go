package main

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type editMode uint8

const (
	modeAdd editMode = iota
	modeEdit
)

type editItemContainer struct {
	entry       *widget.Entry
	keepOpen    *widget.Check
	nameLabel   string
	checkLabel  string
	saveLabel   string
	cancelLabel string
	onSave      func()
	onCancel    func()
	mode        editMode
}

func (eic *editItemContainer) render() *fyne.Container {
	eic.entry = widget.NewEntry()
	eic.entry.Validator = func(s string) error {
		if s == "" {
			return errors.New("empty input")
		}
		return nil
	}
	eic.entry.OnSubmitted = func(_ string) {
		eic.onSave()
	}
	getCheck := func() fyne.CanvasObject {
		if eic.mode == modeAdd {
			eic.keepOpen = widget.NewCheck(eic.checkLabel, nil)
			return eic.keepOpen
		}
		return nil
	}
	btnSave := widget.NewButtonWithIcon(eic.saveLabel, theme.ConfirmIcon(), eic.onSave)
	btnSave.Importance = widget.SuccessImportance
	return container.NewVBox(
		container.NewBorder(
			nil,
			nil,
			widget.NewLabel(eic.nameLabel),
			nil,
			eic.entry,
		),
		container.NewBorder(
			nil,
			nil,
			getCheck(), // will be nil if mode == modeEdit
			container.NewHBox(
				widget.NewButtonWithIcon(eic.cancelLabel, theme.CancelIcon(), eic.onCancel),
				btnSave,
			),
		),
	)
}
