package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	keyword          string
	data             string
	target           string
	mineRes          *MineResult
	correctRes       *MineResult
	rowsData         []*g.TableRowWidget
	mainWindow       *g.WindowWidget
	codeEditorWindow *g.WindowWidget
	mineLogs         []*g.TableRowWidget
	editor           g.CodeEditorWidget
	slotty           []MineResult
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
		g.InputText(&keyword),
		g.Event().OnKeyPressed(g.KeyEnter, func() {
			popSlotty()
		}),
		g.TabBar().TabItems(
			g.TabItem("Full Index").Layout(
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
			g.TabItem("Slots").Layout(
				g.Custom(func() {
					g.PushStyleColor(g.StyleColorText, color.RGBA{R: 0, G: 255, B: 255, A: 255})
					for _, slots := range slotty {
						x := slots.Datahash
						g.Align(g.AlignCenter).To(
							g.Style().
								SetColor(g.StyleColorButtonHovered, color.RGBA{0x36, 0x74, 0xD5, 255}).
								SetColor(g.StyleColorButtonActive, color.RGBA{0x36, 0x74, 0xD5, 255}).
								SetStyle(g.StyleVarFramePadding, 20, 20).
								To(
									g.Button(x).Size(550, 80).OnClick(func() {
										fmt.Println(unlockData(slots.Source, slots.Datahash, slots.Target, slots.Nonce))
									}),
								),
						).Build()
					}
					g.PopStyleColor()
				}),
				g.PopupModal("Unlock").Layout(
					g.Label("Rotation not correct!"),
					g.ImageWithFile("bomb2.png").Size(300, 200),
					g.Button("Close").OnClick(g.CloseCurrentPopup),
				),
			),
			// g.TabItem("Tree").Layout(
			// 	g.TreeTable().
			// 		Columns(g.TableColumn("Name"), g.TableColumn("Size")).
			// 		Rows(
			// 			[]*g.TreeTableRowWidget{
			// 				g.TreeTableRow("Folder1", g.Label("")).Children(
			// 					g.TreeTableRow("File1", g.Label("1MB")),
			// 					g.TreeTableRow("File2", g.Label("2MB")),
			// 				),
			// 			}...,
			// 		).
			// 		Size(g.Auto, g.Auto),
			// ),
		),
	)
	g.PopStyleColorV(2)

	// if showMine {
	showMinePanel(w, h)
	// }
	showDataPanel(w, h)

}

func btnPopup() {
	g.OpenPopup("Unlock")
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
					logString := fmt.Sprintf("S: %s D: %s R: %s", mineRes.Source[:10], mineRes.Datahash[:10], mineRes.Rotation)
					mineLogs = append(mineLogs, g.TableRow(g.Label(logString)))
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

	res := MineResult{}
	for _, v := range files {
		contents, _ := ioutil.ReadFile("../.history/index/" + v.Name())
		if !strings.Contains(string(contents), "\n") {
			err := json.Unmarshal(contents, &res)
			if err != nil {
				log.Fatal(err)
			}
			target := res.Target
			rowsData = append(rowsData, g.TableRow(
				g.Label(res.Source),
				g.Label(res.Datahash),
				g.Label(target),
				g.Label(v.Name()),
				g.Label(res.Nonce)).MinHeight(22),
			)
		}
	}
}

func showDataPanel(w float32, h float32) {

	flags := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsAlwaysAutoResize | g.WindowFlagsNoCollapse | g.WindowFlagsAlwaysVerticalScrollbar
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
	codeEditorWindow.Layout(g.Button("Save Code").OnClick(
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

}

func unlockData(source string, datahash string, target string, nonce string) string {
	correctRes = &MineResult{}
	intNonce, _ := strconv.Atoi(nonce)
	editorRotation := mineFunction(source, datahash, target, intNonce)
	correctRes, _ = Mine(source, datahash, target, intNonce, false)
	cmdString := fmt.Sprintf("go test -run MineFunction -s=%s -d=%s -t=%s -nonce=%d", source, datahash, target, intNonce)

	cmd := exec.Command("bash", "-c", cmdString)

	err := cmd.Run()
	//Err is nil if rotation test is succeeds
	if err == nil {
		unlockedData := DBGetData(datahash)
		correctRes.Datahash = unlockedData
		AddToIndex(source, editorRotation, "WWW", *correctRes)
		popSlotty()
		return unlockedData
	}
	btnPopup()
	return "Rotation not correct!"
}

func popSlotty() {
	kwCheck := sha256.Sum256([]byte(keyword))
	kwHash := hex.EncodeToString(kwCheck[:])
	slotty = nil
	contents, err := ioutil.ReadFile("../.history/index/" + kwHash + ".txt")
	if err == nil {
		jsonRes := MineResult{}
		kwArr := strings.Split(string(contents), "\n")
		// fmt.Println("Last Element", reflect.TypeOf(kwArr[2]))
		for _, v := range kwArr {
			if v != "" {
				rotationFile, _ := ioutil.ReadFile("../.history/index/" + v + ".json")
				err = json.Unmarshal(rotationFile, &jsonRes)
				if err != nil {
					fmt.Errorf("%w", err)
				}
				slotty = append(slotty, jsonRes)
			}
		}
	}
}
