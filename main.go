package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/mattn/go-sqlite3"
)

var (
	inputBoxPanel  TextInputPanel
	goboard        *Goboard
	panelWidth          = 500
	panelHeight         = 400
	inputBoxWidth       = 500
	inputBoxHeight      = 380
	insertFlag     bool = false
)

func update() {}
func quit() {
	rl.CloseWindow()
}
func input() {
	if rl.IsKeyPressed(rl.KeyI) && !insertFlag {
		insertFlag = true
		rl.GetCharPressed()
	}
	if rl.IsKeyPressed(rl.KeyBackSlash) {
		inputBoxPanel.Reset()
		insertFlag = false
	}

}
func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(25, 23, 36, 255))
	if insertFlag {

		inputBoxPanel.Write()
		inputBoxPanel.DrawTextInputPanel()

	}
	rl.EndDrawing()
}

func init() {

	var err error
	rl.InitWindow(800, 600, "Go Board")

	height := rl.GetMonitorHeight(rl.GetCurrentMonitor())
	width := 600

	rl.SetWindowSize(width, height)
	rl.SetWindowPosition(0, 50)
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	goboard, err = NewGoboard("goboard.db", width, height)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	inputBoxPanel = NewInputTextPanel(panelWidth, panelHeight, inputBoxWidth, inputBoxHeight, (goboard.WindowWidth-panelWidth)/2, (goboard.WindowHeight-panelHeight)/2)
	rl.SetTextureFilter(goboard.Font.Texture, rl.FilterPoint)

}

func main() {

	defer goboard.Close()
	defer quit()
	goboard.ReadTasks()
	for !rl.WindowShouldClose() {

		input()
		update()
		render()
	}
}
