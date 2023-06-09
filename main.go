package main

import (
	"log"
	"os"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// create Gtk app
	const gtkID = "xyz.pinegrove.juke"
	gtkJuke, err :=
			gtk.ApplicationNew(gtkID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	// activate
	gtkJuke.Connect("activate", func() {
		// create winMain
		winMain, err := gtk.ApplicationWindowNew(gtkJuke)
		if err != nil {
			log.Fatal("Could not spawn `winMain`.", err)
		}

		winMain.SetTitle("juke")
		winMain.ShowAll()
	})
	// run Gtk app
	gtkJuke.Run(os.Args)
}
