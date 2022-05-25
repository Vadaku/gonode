package imgui

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v4"
)

type windowFlags struct {
	noTitlebar     bool
	noScrollbar    bool
	noMenu         bool
	noMove         bool
	noResize       bool
	noCollapse     bool
	noNav          bool
	noBackground   bool
	noBringToFront bool
}

type inputTextFlags struct {
	noOverwrite bool
}

func (f inputTextFlags) textCombined() imgui.InputTextFlags {
	flags := imgui.InputTextFlagsNone
	if f.noOverwrite {
		flags |= imgui.InputTextFlagsAlwaysOverwriteMode
	}
	return flags
}

func (f windowFlags) windowCombined() imgui.WindowFlags {
	flags := imgui.WindowFlagsNone
	if f.noTitlebar {
		flags |= imgui.WindowFlagsNoTitleBar
	}
	if f.noScrollbar {
		flags |= imgui.WindowFlagsNoScrollbar
	}
	if !f.noMenu {
		flags |= imgui.WindowFlagsMenuBar
	}
	if f.noMove {
		flags |= imgui.WindowFlagsNoMove
	}
	if f.noResize {
		flags |= imgui.WindowFlagsNoResize
	}
	if f.noCollapse {
		flags |= imgui.WindowFlagsNoCollapse
	}
	if f.noNav {
		flags |= imgui.WindowFlagsNoNav
	}
	if f.noBackground {
		flags |= imgui.WindowFlagsNoBackground
	}
	if f.noBringToFront {
		flags |= imgui.WindowFlagsNoBringToFrontOnFocus
	}
	return flags
}

var window = struct {
	flags   windowFlags
	noClose bool

	widgets widgets
}{}

var inputText = struct {
	flags inputTextFlags
}{}

func ShowMine(keepOpen *bool) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 150, Y: 0}, 0, imgui.Vec2{})
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 500, Y: 125}, 0)

	// fmt.Println(imgui.WindowSize())
	if window.noClose {
		keepOpen = nil
	}
	window.flags.noResize = true
	if !imgui.BeginV("Mine", keepOpen, window.flags.windowCombined()) {
		// Early out if the window is collapsed, as an optimization.
		imgui.End()
		return
	}

	var (
		sourceText string
		dataText   string
		targetText string
	)

	inputText.flags.noOverwrite = false
	imgui.InputText("Source", &sourceText)
	imgui.InputText("Data", &dataText)
	imgui.InputText("Target", &targetText)
	imgui.SetKeyboardFocusHereV(1)

	if imgui.Button("Mine") {
		fmt.Println(sourceText)
	}

	imgui.End()
}

func testText(x string) {

}

func ShowIndex(keepOpen *bool) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 150, Y: 125}, 0, imgui.Vec2{})
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 500, Y: 125}, 0)

	if window.noClose {
		keepOpen = nil
	}
	window.flags.noResize = true
	if !imgui.BeginV("Index", keepOpen, window.flags.windowCombined()) {
		// Early out if the window is collapsed, as an optimization.
		imgui.End()
		return
	}

	// End of ShowDemoWindow()
	imgui.End()
}

func ShowData(keepOpen *bool) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 150, Y: 250}, 0, imgui.Vec2{})
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 500, Y: 125}, 0)

	if window.noClose {
		keepOpen = nil
	}
	window.flags.noResize = true
	if !imgui.BeginV("Data", keepOpen, window.flags.windowCombined()) {
		// Early out if the window is collapsed, as an optimization.
		imgui.End()
		return
	}

	// All demo contents
	window.widgets.show()

	// End of ShowDemoWindow()
	imgui.End()
}

func ShowHashwall(keepOpen *bool) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 150, Y: 375}, 0, imgui.Vec2{})
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 500, Y: 125}, 0)

	if window.noClose {
		keepOpen = nil
	}
	window.flags.noResize = true
	if !imgui.BeginV("Hashwall", keepOpen, window.flags.windowCombined()) {
		// Early out if the window is collapsed, as an optimization.
		imgui.End()
		return
	}

	if imgui.Button("Mine To Unlock") {
		fmt.Println("Button Pressed")
	}

	// End of ShowDemoWindow()
	imgui.End()
}

type widgets struct {
	buttonClicked int
	check         bool
	radio         int
}

func (widgets *widgets) show() {
	if !imgui.CollapsingHeader("Widgets") {
		return
	}

	if imgui.TreeNode("Basic") {

		sourceText := "source"
		dataText := ""
		targetText := ""

		imgui.InputText("Source", &sourceText)
		imgui.InputText("Data", &dataText)
		imgui.InputText("Target", &targetText)

		if imgui.Button("Mine") {
			fmt.Println(sourceText)
		}

		imgui.TreePop()
	}
}
