package main

import (
	"harishkrishnan24/pixl/apptype"
	"harishkrishnan24/pixl/swatch"
	"harishkrishnan24/pixl/ui"
	"image/color"

	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA64{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
