package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
)

var (
	showMine  bool
	showIndex bool
	showData  bool
)

func loop() {
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.MenuItem("File"),
		),
		g.Label("Go Node"),
		g.Column(
			g.Style().SetColor(g.StyleColorBorder, color.RGBA{0x36, 0x74, 0xD5, 255}).To(
				g.Checkbox("Mine", &showMine),
			),
			g.Checkbox("Index", &showIndex),
			g.Checkbox("Data", &showData),
		),
	)
	g.PopStyleColorV(2)

	if showMine {
		showMinePanel()
	}
}

func showMinePanel() {
	g.Window("Hey").Pos(10, 30).Size(200, 100).IsOpen(&showMine).Layout(
		g.Label("I'm a label in window 2"),
		g.Button("Hide me").OnClick(nil),
	)
}
