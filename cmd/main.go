package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
)

func main() {
	app := app.New()
	w := app.NewWindow("Game Pad")
	w.SetIcon(theme.DocumentIcon())
	w.SetMaster()
	w.Resize(fyne.NewSize(400.0, 500.0))

	nav := navigationController{
		win:     w,
		current: pIDGameList,
	}

	gameNames := []string{}

	gameList := gameListContainer{
		headerLabel: "Select game",
		buttonLabel: "Add",
		data:        binding.BindStringList(&gameNames),
		onAdd: func() {
			nav.load(pIDEditItem)
		},
	}

	editItem := editItemContainer{
		nameLabel:   "Name",
		checkLabel:  "Keep open",
		saveLabel:   "Save",
		cancelLabel: "Cancel",
	}
	editItem.onCancel = func() {
		editItem.entry.SetText("")
		nav.back()
	}
	editItem.onSave = func() {
		if err := editItem.entry.Validate(); err != nil {
			log.Println(err.Error())
			return
		}
		if err := gameList.data.Append(editItem.entry.Text); err != nil {
			log.Println(err.Error())
		}
		editItem.entry.SetText("")
		if !editItem.keepOpen.Checked {
			nav.back()
		}
	}

	nav.set(pIDGameList, &gameList)
	nav.set(pIDEditItem, &editItem)
	nav.load(nav.current)

	w.Show()
	app.Run()

	log.Printf("%#v", gameNames)
}
