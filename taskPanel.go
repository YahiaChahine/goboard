package main

import rl "github.com/gen2brain/raylib-go/raylib"

type TaskPanel struct {
	PanelWidth    int
	PanelHeight   int
	taskBoxWidth  int
	taskBoxHeight int
	PosX          int
	PosY          int
	TaskBoxY      int
	task          Task
}

func NewTaskPanel(pw, ph, ibw, ibh, x, y int, t Task) TaskPanel {
	return TaskPanel{
		PanelWidth:    pw,
		PanelHeight:   ph,
		taskBoxHeight: ibh,
		taskBoxWidth:  ibw,
		PosX:          x,
		PosY:          y,
		TaskBoxY:      y + 30,
		task:          t,
	}
}

func (p *TaskPanel) DrawTaskPanel() {

	rl.DrawRectangle(int32(p.PosX), int32(p.PosY), int32(p.PanelWidth), int32(p.PanelHeight), goboard.Overlay)
	rl.DrawRectangleLines(int32(p.PosX), int32(p.PosY), int32(p.PanelWidth), int32(p.PanelHeight), goboard.Muted)

	panelTitleX := p.PosX + (p.PanelWidth-int(rl.MeasureText(p.task.Title, 14)))/2
	rl.DrawTextEx(goboard.Font, p.task.Title, rl.Vector2{X: float32(panelTitleX), Y: float32(p.PosY + 6)}, 16, 1, goboard.Pine)

	// Draw the text box rectangle (centered)
	rl.DrawRectangle(int32(p.PosX), int32(p.TaskBoxY), int32(p.taskBoxWidth), int32(p.taskBoxHeight), goboard.Surface)
	rl.DrawRectangleLines(int32(p.PosX), int32(p.TaskBoxY), int32(p.taskBoxWidth), int32(p.taskBoxHeight), rl.White)

	rl.DrawTextEx(goboard.Font, p.task.Description, rl.Vector2{X: float32((p.PosX + 10)), Y: float32((p.TaskBoxY + 10))}, float32(20), 1, goboard.Text)
	rl.DrawTextEx(goboard.Font, p.task.RepeatDays, rl.Vector2{X: float32((p.PosX + 10)), Y: float32((p.TaskBoxY + 10) + (18 * 1))}, float32(20), 1, goboard.Text)
	rl.DrawTextEx(goboard.Font, p.task.StartTime.String(), rl.Vector2{X: float32((p.PosX + 10)), Y: float32((p.TaskBoxY + 10) + (18 * 2))}, float32(20), 1, goboard.Text)
	rl.DrawTextEx(goboard.Font, p.task.EndTime.String(), rl.Vector2{X: float32((p.PosX + 10)), Y: float32((p.TaskBoxY + 10) + (18 * 3))}, float32(20), 1, goboard.Text)

}
