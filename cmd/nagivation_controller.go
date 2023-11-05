package main

import (
	"log"

	"fyne.io/fyne/v2"
)

type renderer interface {
	render() *fyne.Container
}

type pageID uint8

const (
	pIDGameList pageID = iota
	pIDEditItem
)

type navigationController struct {
	win      fyne.Window
	pages    map[pageID]renderer
	previous pageID
	current  pageID
}

func (nc *navigationController) set(id pageID, c renderer) {
	if nc.pages == nil {
		nc.pages = make(map[pageID]renderer)
	}
	nc.pages[id] = c
}

func (nc *navigationController) load(id pageID) {
	if nc.win == nil {
		log.Println("nil win for navigationController, aborting load")
		return
	}
	if page, ok := nc.pages[id]; ok {
		nc.previous = nc.current
		nc.current = id
		nc.win.SetContent(page.render())
		return
	}
	log.Printf("navigationController.load(): no value for id: %d", id)
}

func (nc *navigationController) back() {
	nc.load(nc.previous)
}
