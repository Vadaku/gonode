package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	source           string
	data             string
	target           string
	mineRes          *MineResult
	correctRes       *MineResult
	rowsData         []*g.TableRowWidget
	mainWindow       *g.WindowWidget
	codeEditorWindow *g.WindowWidget
	mineLogs         []*g.TableRowWidget
	editor           g.CodeEditorWidget
)

var onload = true

func loop(w float32, h float32) {
	if onload {
		populateTable()
		editor.ShowWhitespaces(false).Text(`/**
	Source and data parameters are sha256 hashes.
	Target is a valid target.
**/
func mineFunction(source string, data string, target string, nonce int) string {
	rotationHash := source + data + strconv.Itoa(nonce)
	_ = rotationHash
	//Complete code below...

	return ""
}`)
		onload = false
	}
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	mainWindow = g.SingleWindowWithMenuBar().Pos(0, 0).Size(w, h/2).Flags(g.WindowFlagsHorizontalScrollbar | g.WindowFlagsNoResize |
		g.WindowFlagsNoMove | g.WindowFlagsMenuBar | g.WindowFlagsNoTitleBar)
	mainWindow.Layout(
		g.MenuBar().Layout(
			g.MenuItem("File"),
			g.Separator(),
			g.MenuItem("Settings"),
			g.Align(g.AlignRight).To(
				g.MenuItem("Refresh").OnClick(func() {
					rowsData = nil
					populateTable()
				}),
			),
		),
		g.TabBar().TabItems(
			g.TabItem("Default").Layout(
				g.Table().Flags(g.TableFlagsNoBordersInBody).
					Columns(
						g.TableColumn("Source"),
						g.TableColumn("Datahash"),
						g.TableColumn("Target"),
						g.TableColumn("Rotation"),
						g.TableColumn("Nonce"),
					).
					Rows(
						rowsData...,
					).Size(g.Auto, g.Auto),
			),
			g.TabItem("Raw"),
			g.TabItem("Tree").Layout(
				g.TreeTable().
					Columns(g.TableColumn("Name"), g.TableColumn("Size")).
					Rows(
						[]*g.TreeTableRowWidget{
							g.TreeTableRow("Folder1", g.Label("")).Children(
								g.TreeTableRow("File1", g.Label("1MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
							),
						}...,
					).
					Size(g.Auto, g.Auto),
			),
		),
	)
	g.PopStyleColorV(2)

	// if showMine {
	showMinePanel(w, h)
	// }

	showDataPanel(w, h)

}

func showMinePanel(w float32, h float32) {
	// res := &MineResult{}
	flags := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsAlwaysAutoResize | g.WindowFlagsNoCollapse
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	mineWindow := g.Window("Mine").Flags(flags).Pos(w-500, h/2).Size(w-(w-500), h-(h/2))
	mineWindow.Layout(
		leftLabel("Source", &source),
		leftLabel("Data    ", &data),
		leftLabel("Target ", &target),
		g.Align(g.AlignCenter).To(
			g.Row(
				g.Button("Upload"),
				g.Button("Mine").OnClick(func() {
					mineRes, _ = Mine(source, data, target, 0, true)
					// logString := res.Rotation
					mineLogs = append(mineLogs, g.TableRow(g.Label(mineRes.Rotation)))
					rowsData = nil
					populateTable()
				}),
			)),
		g.Separator(),
		// g.InputTextMultiline(&raw).Flags(g.InputTextFlagsReadOnly).Size(g.Auto, g.Auto),
	)
	logs := g.Window("None").Flags(g.WindowFlagsNoTitleBar|g.WindowFlagsNoMove|g.WindowFlagsNoResize|g.WindowFlagsAlwaysHorizontalScrollbar).Pos(w-490, (h/2)+(135)).Size(w-(w-480), h-((h/2)+145))
	logs.Layout(
		g.Align(g.AlignCenter).To(
			g.Label("Mining Logs"),
		),
		g.Table().Rows(
			mineLogs...,
		),
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

func populateTable() {
	f, err := os.Open("../.history/index/")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		contents, _ := ioutil.ReadFile("../.history/index/" + v.Name())
		if !strings.Contains(string(contents), "\n") {
			target := strings.Trim(string(contents)[168:200], "0")
			rowsData = append(rowsData, g.TableRow(
				g.Label(string(contents)[40:104]),
				g.Custom(func() {
					g.Selectable(string(contents)[104:168]).OnClick(func() {
						unlockData(string(contents)[40:104], string(contents)[104:168], target, string(contents)[265:])
					}).Build()
				}),
				g.Label(target),
				g.Label(v.Name()),
				g.Label(string(contents)[265:])).MinHeight(22),
			)
		}
	}
}

func showDataPanel(w float32, h float32) {
	flags := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsAlwaysAutoResize | g.WindowFlagsNoCollapse
	g.PushStyleColor(g.StyleColorWindowBg, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	g.PushStyleColor(g.StyleColorBorder, color.RGBA{R: 0, G: 255, B: 255, A: 255})
	codeEditorWindow = g.Window("Hashwall").Flags(flags).Pos(0, h/2).Size(w-500, h/2)
	codeEditorWindow.Layout(
		g.Custom(func() {
			imgui.SetScrollHereY(0.5)
		}),
		g.Label("Finish this mine function and successfully mine the correct rotation to unlock hashed data\n\n"),
		editor.HandleKeyboardInputs(true).Border(true).Size(w-600, (h-500)),
	)
	codeEditorWindow.Layout(g.Button("Test Code").OnClick(
		func() {
			executeCode()
		},
	))
	g.PopStyleColorV(2)
}

//Execute editor window code.
func executeCode() {
	codeEditorText := []byte(`package main
	import (
		_"crypto/sha256"
		_"encoding/hex"
		"strconv"
	)` +
		editor.GetText())

	if err := ioutil.WriteFile("mineFunc.go", codeEditorText, 0644); err != nil {
		log.Fatal("Couldn't write to temp file", err)
	}

	cmd := exec.Command("bash", "-c", "go test -run MineFunction")

	output, _ := cmd.Output()
	fmt.Println(string(output))

	cmd.Wait()
}

func unlockData(source string, datahash string, target string, nonce string) {
	correctRes = &MineResult{}
	intNonce, _ := strconv.Atoi(nonce)
	correctRes, _ = Mine(source, data, target, intNonce, false)
	fmt.Println(source, datahash, target, nonce)
}
