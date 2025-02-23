package main

import (
	"fmt"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextInputPanel struct {
	panelWidth     int
	panelHeight    int
	inputBoxWidth  int
	inputBoxHeight int
	posX           int
	posY           int
	inputBoxY      int
	buffer         []string
	cursorX        int
	cursorY        int
	cursorBlink    float32
	initFlag       bool
}

func NewInputTextPanel(pw, ph, ibw, ibh, x, y int) TextInputPanel {
	return TextInputPanel{
		panelWidth:     pw,
		panelHeight:    ph,
		inputBoxHeight: ibh,
		inputBoxWidth:  ibw,
		posX:           x,
		posY:           y,
		inputBoxY:      y + 30,
		cursorX:        0,
		cursorY:        0,
		cursorBlink:    0,
		buffer:         []string{},
		initFlag:       false,
	}
}

func (p *TextInputPanel) Animate() {
	if !p.initFlag {
		p.initFlag = true
		for i := 100; i > 0; i = i - 1 {
			rl.BeginDrawing()
			rl.DrawRectangle(int32(p.posX/i), int32(p.posY/i), int32(p.panelWidth/i), int32(p.panelHeight/i), goboard.Overlay)
			rl.DrawRectangleLines(int32(p.posX/i), int32(p.posY/i), int32(p.panelWidth/i), int32(p.panelHeight/i), goboard.Muted)

			rl.DrawRectangle(int32(p.posX/i), int32(p.inputBoxY/i), int32(p.inputBoxWidth/i), int32(p.inputBoxHeight/i), goboard.Surface)
			rl.DrawRectangleLines(int32(p.posX/i), int32(p.inputBoxY/i), int32(p.inputBoxWidth/i), int32(p.inputBoxHeight/i), rl.White)
			time.Sleep(10 * time.Millisecond)
			rl.EndDrawing()
		}
	}
	rl.BeginDrawing()
	p.DrawTextInputPanel()
}

func (p *TextInputPanel) DrawTextInputPanel() {
	var title string = "New Task"

	rl.DrawRectangle(int32(p.posX), int32(p.posY), int32(p.panelWidth), int32(p.panelHeight), goboard.Overlay)
	rl.DrawRectangleLines(int32(p.posX), int32(p.posY), int32(p.panelWidth), int32(p.panelHeight), goboard.Muted)

	if len(p.buffer) > 0 {
		if len(p.buffer[0]) > 0 {
			title = p.buffer[0]
		}
	}
	panelTitleX := p.posX + (p.panelWidth-int(rl.MeasureText(title, 14)))/2
	rl.DrawTextEx(goboard.Font, title, rl.Vector2{X: float32(panelTitleX), Y: float32(p.posY + 6)}, 16, 1, goboard.Pine)

	// Draw the text box rectangle (centered)
	rl.DrawRectangle(int32(p.posX), int32(p.inputBoxY), int32(p.inputBoxWidth), int32(p.inputBoxHeight), goboard.Surface)
	rl.DrawRectangleLines(int32(p.posX), int32(p.inputBoxY), int32(p.inputBoxWidth), int32(p.inputBoxHeight), rl.White)

	for i, line := range p.buffer {
		rl.DrawTextEx(goboard.Font, line, rl.Vector2{X: float32((p.posX + 10)), Y: float32((p.inputBoxY + 10) + (18 * i))}, float32(20), 1, goboard.Text)
	}

	p.cursorBlink += rl.GetFrameTime()
	if p.cursorBlink >= 1.0 {
		p.cursorBlink = 0.0
	}

	if p.cursorBlink < 0.5 {
		if p.cursorY >= len(p.buffer) {
			p.buffer = append(p.buffer, "")
		}
		cursorPosX := int32(p.posX) + 10 + int32(rl.MeasureTextEx(goboard.Font, p.buffer[p.cursorY][:p.cursorX], float32(20), float32(1)).X)
		rl.DrawLine(cursorPosX, int32(p.inputBoxY+10+(p.cursorY*20)), cursorPosX, int32(p.inputBoxY+30+(p.cursorY*20)), goboard.Subtle)
	}
}

func (p *TextInputPanel) Write() {

	key := rl.GetCharPressed()
	for key > 0 {

		if p.cursorY >= len(p.buffer) {
			p.buffer = append(p.buffer, "")
		}

		p.buffer[p.cursorY] = p.buffer[p.cursorY][:p.cursorX] + string(rune(key)) + p.buffer[p.cursorY][p.cursorX:]
		p.cursorX++
		key = rl.GetCharPressed()
	}

	// Handle backspace
	if rl.IsKeyDown(rl.KeyBackspace) && p.cursorX > 0 {
		p.buffer[p.cursorY] = p.buffer[p.cursorY][:p.cursorX-1] + p.buffer[p.cursorY][p.cursorX:]
		p.cursorX--
		time.Sleep(70 * time.Millisecond)
	} else if rl.IsKeyDown(rl.KeyBackspace) && p.cursorX == 0 && p.cursorY > 0 {
		p.cursorY--
		p.cursorX = len(p.buffer[p.cursorY])
		time.Sleep(70 * time.Millisecond)
	}

	// Handle cursor movement
	if rl.IsKeyPressed(rl.KeyLeft) && p.cursorX > 0 {
		p.cursorX--
	}
	if rl.IsKeyPressed(rl.KeyRight) && p.cursorX < len(p.buffer[p.cursorY]) {
		p.cursorX++
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		p.cursorY++
		p.cursorX = 0
		if p.cursorY > len(p.buffer) {
			p.buffer = append(p.buffer, "")
		}
	}
}

func (p *TextInputPanel) Reset() {
	if len(p.buffer) > 2 {
		if t, err := p.ParseInput(); err != nil {
			fmt.Println(err)
		} else {
			goboard.InsertTask(t)
		}

	}
	p.cursorX = 0
	p.cursorY = 0
	p.buffer = []string{}
	p.initFlag = false
}

func (p *TextInputPanel) ParseInput() (Task, error) {

	if len(p.buffer) < 2 {
		return Task{}, fmt.Errorf("invalid task format: must atleast have a name and a description for the task")
	}

	// Parse start date & end date ie line 3
	var startDate *time.Time
	var endDate *time.Time
	dates := strings.Split(p.buffer[2], " ")
	if strings.TrimSpace(p.buffer[2]) != "" {
		parsedStartDate, err := time.Parse("02-01-2006", strings.TrimSpace(dates[0]))
		if err != nil {
			return Task{}, fmt.Errorf("invalid start date format: %w", err)
		}
		startDate = &parsedStartDate

		parsedEndDate, err := time.Parse("02-01-2006", strings.TrimSpace(dates[1]))
		if err != nil {
			return Task{}, fmt.Errorf("invalid end date format: %w", err)
		}
		endDate = &parsedEndDate
	}

	timings := strings.Split(p.buffer[3], " ")
	// Parse start time and end time
	var startTime *time.Time
	var endTime *time.Time
	if strings.TrimSpace(p.buffer[3]) != "" {
		parsedStartTime, err := time.Parse("15:04", strings.TrimSpace(timings[0]))
		if err != nil {
			return Task{}, fmt.Errorf("invalid start time format: %w", err)
		}
		startTime = &parsedStartTime

		parsedEndTime, err := time.Parse("15:04", strings.TrimSpace(timings[1]))
		if err != nil {
			return Task{}, fmt.Errorf("invalid end time format: %w", err)
		}
		endTime = &parsedEndTime
	}

	fmt.Println(timings[1], timings[2])
	rpeatDays := timings[2]

	return Task{
		Title:       strings.TrimSpace(p.buffer[0]),
		Description: strings.TrimSpace(p.buffer[1]),
		StartDate:   startDate,
		EndDate:     endDate,
		StartTime:   startTime,
		EndTime:     endTime,
		RepeatDays:  rpeatDays,
		Cancel:      false,
	}, nil
}
