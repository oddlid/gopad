package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
	app := app.NewWithID("net.oddware.gopad")
	w := app.NewWindow("Game Pad")
	w.SetIcon(resourceAppIcon)
	w.SetMaster()
	w.Resize(fyne.NewSize(400.0, 500.0))

	nav := navigationController{
		win:     w,
		current: pIDGameList,
	}

	gameNames := []string{"Poker", "10000", "Blaff"}

	gameList := gameListContainer{
		nav:         &nav,
		headerLabel: "Select game",
		buttonLabel: "Add",
		data:        binding.BindStringList(&gameNames),
		onAdd: func() {
			nav.load(pIDAddItem)
		},
	}

	addItem := editItemContainer{
		nameLabel:   "Name",
		checkLabel:  "Keep open",
		saveLabel:   "Save",
		cancelLabel: "Cancel",
		mode:        modeAdd,
	}
	addItem.onCancel = func() {
		addItem.entry.SetText("")
		nav.back()
	}
	addItem.onSave = func() {
		if err := addItem.entry.Validate(); err != nil {
			log.Println(err.Error())
			return
		}
		if err := gameList.data.Append(addItem.entry.Text); err != nil {
			log.Println(err.Error())
		}
		addItem.entry.SetText("")
		if !addItem.keepOpen.Checked {
			nav.back()
		}
	}

	editItem := editItemContainer{
		nameLabel:   "Name",
		saveLabel:   "Save",
		cancelLabel: "Cancel",
		mode:        modeEdit,
	}
	editItem.onCancel = func() {
		editItem.entry.SetText("")
		nav.back()
	}

	nav.set(pIDGameList, &gameList)
	nav.set(pIDAddItem, &addItem)
	nav.set(pIDEditItem, &editItem)
	nav.load(nav.current)

	w.Show()
	app.Run()

	log.Printf("%#v", gameNames)
}
