package ui

import (
	"harishkrishnan24/pixl/apptype"
	"harishkrishnan24/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
