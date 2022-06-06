package main

import (
	"image/color"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	showMine   bool
	showIndex  bool
	showData   bool
	source     string
	data       string
	target     string
	raw        string
	x          string
	mainWindow *g.WindowWidget
)

func loop(w float32, h float32) {
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	mainWindow = g.SingleWindowWithMenuBar().Pos(0, 0).Size(w, h/2)
	mainWindow.Layout(
		g.MenuBar().Layout(
			g.MenuItem("File"),
			g.Separator(),
			g.MenuItem("Settings"),
		),
		g.TabBar().TabItems(
			g.TabItem("Default"),
			g.TabItem("Raw"),
		),
		g.Table().Flags(g.TableFlagsNoBordersInBody).
			Columns(
				g.TableColumn("Source"),
				g.TableColumn("Datahash"),
				g.TableColumn("Target"),
				g.TableColumn("Rotation"),
				g.TableColumn("Nonce"),
			).
			Rows(
				g.TableRow(
					g.Label("ss"),
					g.Label("ssss"),
					g.Label("sssss"),
					g.Label(x),
					g.Label("sssss"),
				),
				g.TableRow(
					g.Label("ssss"),
				),
			),
	)
	g.PopStyleColorV(2)

	// if showMine {
	showMinePanel(w, h)
	// }

	showIndexPanel(w, h)

}

func showMinePanel(w float32, h float32) {
	res := &MineResult{}
	flags := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsAlwaysAutoResize | g.WindowFlagsNoCollapse
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	mineWindow := g.Window("Mine").Flags(flags).Pos(w-500, h/2).Size(w-(w-500), h-(h/2))
	mineWindow.Layout(
		leftLabel("Source", &source),
		// g.InputText(&source).Label("Source").Size(50),
		leftLabel("Data    ", &data),
		leftLabel("Target ", &target),
		g.Align(g.AlignCenter).To(
			g.Row(
				g.Button("Upload"),
				g.Button("Mine").OnClick(func() {
					res, raw = Mine(source, data, target, nil)
					x = res.Rotation
				}),
			)),
		g.Separator(),
		// g.InputTextMultiline(&raw).Flags(g.InputTextFlagsReadOnly).Size(g.Auto, g.Auto),
	)
	logs := g.Window("None").Flags(g.WindowFlagsNoTitleBar|g.WindowFlagsNoMove|g.WindowFlagsNoResize).Pos(w-490, (h/2)+(135)).Size(w-(w-480), h-((h/2)+145))
	logs.Layout(
		g.Align(g.AlignCenter).To(
			g.Label("Mining Logs"),
		),
		g.Table().Rows(g.TableRow(
			g.Label(x).Wrapped(true),
		)),
	)
	g.PopStyleColorV(2)
}

func leftLabel(label string, text *string) *g.CustomWidget {
	return g.Custom(func() {
		g.Align(g.AlignCenter).To(g.Custom(func() {
			g.Label(label).Build()
			imgui.SameLine()
			g.InputText(text).Size(200).Build()
		}),
		).Build()
	})
}

func showIndexPanel(w float32, h float32) {
	flags := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsAlwaysAutoResize | g.WindowFlagsNoCollapse
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	g.Window("Index").Flags(flags).Pos(0, h/2).Size(w-500, h/2).Layout(
		g.Table(),
	)
	g.PopStyleColorV(2)
}
