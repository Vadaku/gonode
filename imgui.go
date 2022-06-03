package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
)

var (
	showMine  bool
	showIndex bool
	showData  bool
	source    string
	data      string
	target    string
	raw       string
)

func loop() {
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Checkbox("Mine", &showMine),
			g.Checkbox("Index", &showIndex),
			g.Checkbox("Data", &showData),
		),
	)
	g.PopStyleColorV(2)

	if showMine {
		showMinePanel()
	}

	if showIndex {
		showIndexPanel()
	}

	if showData {
		showDataPanel()
	}

}

func showMinePanel() {
	g.Window("Hey").Flags(g.WindowFlagsNoTitleBar).Pos(800, 0).Size(400, 400).IsOpen(&showMine).Layout(
		g.InputText(&source).Label("Source"),
		g.InputText(&data).Label("Data"),
		g.InputText(&target).Label("Target"),
		g.Button("Mine").OnClick(func() {
			_, raw = Mine(source, data, target, nil)
		}),
		g.Label(raw),
	)
}

func showDataPanel() {}

func showIndexPanel() {}
