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
	tasksPanels    []TaskPanel
	fullscreenFlag bool = false
	move           bool = false
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
	if (rl.IsKeyDown(rl.KeyLeftAlt) || rl.IsKeyDown(rl.KeyRightAlt)) && rl.IsKeyDown(rl.KeyEnter) {
		rl.ToggleBorderlessWindowed()
		fullscreenFlag = !fullscreenFlag
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mouseP := rl.GetMousePosition()
		if mouseP.X > 4 && mouseP.X < 355 {
			if int(mouseP.Y)%200 >= 0 && int(mouseP.Y)%200 <= 20 {
				move = true
			}
		}
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) && move {
		mouseP := rl.GetMousePosition()
		if len(tasksPanels) > 0 {
			tasksPanels[0].posX = int(mouseP.X)
			tasksPanels[0].posY = int(mouseP.Y)
			tasksPanels[0].taskBoxY = int(mouseP.Y) + 30
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		move = false
	}
}
func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(25, 23, 36, 255))
	for _, t := range tasksPanels {
		t.DrawTaskPanel()
	}

	if insertFlag {

		inputBoxPanel.Write()
		inputBoxPanel.Animate()

	}
	if fullscreenFlag {
		//rl.DrawLine(int32(goboard.WindowWidth+40), 0, int32(goboard.WindowHeight), int32(goboard.WindowWidth+40), rl.White)
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
	tasks := goboard.ReadTasks()
	for i, t := range tasks {
		tasksPanels = append(tasksPanels, NewTaskPanel(350, 180, 350, 160, 4, 0+200*i, t))
	}
	for !rl.WindowShouldClose() {

		input()
		update()
		render()
	}
}
