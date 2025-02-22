package main

import (
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
	}
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
	goboard.InsertTask(p.buffer)
	p.cursorX = 0
	p.cursorY = 0
	p.buffer = []string{}
}
