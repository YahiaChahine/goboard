package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	panelToMove         = -1
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
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		mouseP := rl.GetMousePosition()
		if panelToMove == -1 {
			for i, panel := range tasksPanels {
				if panel.PosX <= int(mouseP.X) && (panel.PosX+panel.PanelWidth) >= int(mouseP.X) {
					if panel.PosY <= int(mouseP.Y) && panel.TaskBoxY >= int(mouseP.Y) {
						panelToMove = i
					}
				}
			}
		} else {
			tasksPanels[panelToMove].PosX = int(mouseP.X)
			tasksPanels[panelToMove].PosY = int(mouseP.Y)
			tasksPanels[panelToMove].TaskBoxY = int(mouseP.Y) + 30
		}
	}

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		panelToMove = -1
	}
}
func render() {
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	// Define the starting x position
	startX := int32(500)

	// Calculate cell width and height based on screen size
	cellWidth := (screenWidth - startX) / 7
	cellHeight := screenHeight / 24

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
		days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
		for i, day := range days {
			rl.DrawText(day, startX+int32(i)*cellWidth+10, 10, 20, goboard.Rose)
		}

		// Draw the hourly marks and vertical lines
		for hour := 0; hour < 24; hour++ {
			// Draw the hour label on the left
			hourLabel := fmt.Sprintf("%02d:00", hour)
			rl.DrawText(hourLabel, startX-60, int32(hour)*cellHeight+50, 20, goboard.Rose)

			// Draw a short horizontal line at the first vertical line to mark the hour
			rl.DrawLine(startX, int32(hour)*cellHeight+50, startX+10, int32(hour)*cellHeight+50, goboard.Rose)

			// Draw vertical lines for each day
			for day := 0; day < 7; day++ {
				rl.DrawLine(
					startX+int32(day)*cellWidth, 50,
					startX+int32(day)*cellWidth, screenHeight,
					rl.White,
				)
			}
		}
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
