package main

import rl "github.com/gen2brain/raylib-go/raylib"

type TaskPanel struct {
	panelWidth    int
	panelHeight   int
	taskBoxWidth  int
	taskBoxHeight int
	posX          int
	posY          int
	taskBoxY      int
	task          Task
}

func NewTaskPanel(pw, ph, ibw, ibh, x, y int, t Task) TaskPanel {
	return TaskPanel{
		panelWidth:    pw,
		panelHeight:   ph,
		taskBoxHeight: ibh,
		taskBoxWidth:  ibw,
		posX:          x,
		posY:          y,
		taskBoxY:      y + 30,
		task:          t,
	}
}

func (p *TaskPanel) DrawTaskPanel() {

	rl.DrawRectangle(int32(p.posX), int32(p.posY), int32(p.panelWidth), int32(p.panelHeight), goboard.Overlay)
	rl.DrawRectangleLines(int32(p.posX), int32(p.posY), int32(p.panelWidth), int32(p.panelHeight), goboard.Muted)

	panelTitleX := p.posX + (p.panelWidth-int(rl.MeasureText(p.task.Title, 14)))/2
	rl.DrawTextEx(goboard.Font, p.task.Title, rl.Vector2{X: float32(panelTitleX), Y: float32(p.posY + 6)}, 16, 1, goboard.Pine)

	// Draw the text box rectangle (centered)
	rl.DrawRectangle(int32(p.posX), int32(p.taskBoxY), int32(p.taskBoxWidth), int32(p.taskBoxHeight), goboard.Surface)
	rl.DrawRectangleLines(int32(p.posX), int32(p.taskBoxY), int32(p.taskBoxWidth), int32(p.taskBoxHeight), rl.White)

	rl.DrawTextEx(goboard.Font, p.task.Description, rl.Vector2{X: float32((p.posX + 10)), Y: float32((p.taskBoxY + 10))}, float32(20), 1, goboard.Text)
	rl.DrawTextEx(goboard.Font, p.task.RepeatDays, rl.Vector2{X: float32((p.posX + 10)), Y: float32((p.taskBoxY + 10) + (18 * 1))}, float32(20), 1, goboard.Text)
	rl.DrawTextEx(goboard.Font, p.task.StartTime.String(), rl.Vector2{X: float32((p.posX + 10)), Y: float32((p.taskBoxY + 10) + (18 * 2))}, float32(20), 1, goboard.Text)
	rl.DrawTextEx(goboard.Font, p.task.EndTime.String(), rl.Vector2{X: float32((p.posX + 10)), Y: float32((p.taskBoxY + 10) + (18 * 3))}, float32(20), 1, goboard.Text)

}
