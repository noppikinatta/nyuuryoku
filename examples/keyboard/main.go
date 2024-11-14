package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/noppikinatta/nyuuryoku"
)

const (
	screenW = 800
	screenH = 600
)

func main() {
	ebiten.SetWindowSize(screenW, screenH)

	g := game{}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

type game struct {
	mouse          *nyuuryoku.Mouse
	keyboard       *nyuuryoku.Keyboard
	tmpKeys        []ebiten.Key
	log            []string
	logIdx         int
	pressedInfo    []string
	mouseIsVirtual bool
	virtualMouse   *virtualMouse
}

func (g *game) Update() error {
	if g.mouse == nil {
		g.mouse = nyuuryoku.NewMouse()
	}
	if g.keyboard == nil {
		g.keyboard = nyuuryoku.NewKeyboard()
	}
	if g.virtualMouse == nil {
		g.virtualMouse = newVirtualMouse()
	}

	if g.mouseIsVirtual {
		g.virtualMouse.Update()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.switchMouse()
	}

	g.appendKeysLog()
	g.updatePressedInfo()

	return nil
}

func (g *game) appendKeysLog() {
	g.tmpKeys = g.tmpKeys[:0]
	g.tmpKeys = g.keyboard.AppendJustPressed(g.tmpKeys)
	for _, k := range g.tmpKeys {
		g.appendJustLog(k, "Just Pressed ")
	}

	g.tmpKeys = g.tmpKeys[:0]
	g.tmpKeys = g.keyboard.AppendJustReleased(g.tmpKeys)
	for _, k := range g.tmpKeys {
		g.appendJustLog(k, "Just Released")
	}
}

func (g *game) appendJustLog(key ebiten.Key, prefix string) {
	line := fmt.Sprintf("%s:%s", prefix, key.String())
	g.appendLog(line)
}

func (g *game) updatePressedInfo() {
	g.tmpKeys = g.tmpKeys[:0]
	g.tmpKeys = g.keyboard.AppendPressed(g.tmpKeys)
	g.pressedInfo = g.pressedInfo[:0]
	for _, k := range g.tmpKeys {
		line := fmt.Sprintf("%s pressed: %d frames", k.String(), g.keyboard.PressDuration(k))
		g.pressedInfo = append(g.pressedInfo, line)
	}
}

func (g *game) appendLog(line string) {
	const max = 48
	if len(g.log) < max {
		g.log = append(g.log, line)
	} else {
		g.log[g.logIdx] = line
		g.logIdx++
		g.logIdx = g.logIdx % max
	}
}

func (g *game) switchMouse() {
	setter := nyuuryoku.MouseSetter{Mouse: g.mouse}

	if g.mouseIsVirtual {
		setter.SetDefault()
	} else {
		setter.SetCursorPositionFunc(g.virtualMouse.cursorPosition)
		setter.SetPressedFunc(g.virtualMouse.buttonPressed)
		setter.SetJustPressedFunc(g.virtualMouse.buttonJustPressed)
		setter.SetJustReleasedFunc(g.virtualMouse.buttonJustReleased)
		setter.SetPressDurationFunc(g.virtualMouse.pressDuration)
		setter.SetWheelFunc(g.virtualMouse.Wheel)
	}

	g.mouseIsVirtual = !g.mouseIsVirtual
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 48})
	if g.keyboard == nil {
		return
	}

	g.drawPressedInfo(screen)
	g.drawLog(screen)

	ebitenutil.DebugPrintAt(screen, "press space to switch to virtual mouse", 10, 576)
}

func (g *game) drawPressedInfo(screen *ebiten.Image) {
	for i := range g.pressedInfo {
		info := g.pressedInfo[i]
		y := i * 12

		ebitenutil.DebugPrintAt(screen, info, 400, y)
	}
}

func (g *game) drawLog(screen *ebiten.Image) {
	for i := range g.log {
		line := g.log[i]

		yIdx := (i - g.logIdx)
		if yIdx < 0 {
			yIdx += len(g.log)
		}
		y := yIdx * 12

		ebitenutil.DebugPrintAt(screen, line, 200, y)
	}
}

type virtualMouse struct {
	intervals     map[ebiten.MouseButton]int
	wheelInterval int
	count         int
}

func newVirtualMouse() *virtualMouse {
	return &virtualMouse{
		intervals: map[ebiten.MouseButton]int{
			ebiten.MouseButtonLeft:   60,
			ebiten.MouseButtonRight:  90,
			ebiten.MouseButtonMiddle: 150,
		},
		wheelInterval: 21,
	}
}

func (m *virtualMouse) Update() {
	m.count++
}

func (m *virtualMouse) cursorPosition() (x, y int) {
	x = m.count % screenW
	y = m.count % screenH
	return x, y
}

func (m *virtualMouse) buttonPressed(button ebiten.MouseButton) bool {
	interval := m.intervals[button]
	quotient := m.count / interval
	return quotient%2 == 1
}

func (m *virtualMouse) buttonJustPressed(button ebiten.MouseButton) bool {
	if !m.buttonPressed(button) {
		return false
	}
	return m.duration(button) == 1
}

func (m *virtualMouse) buttonJustReleased(button ebiten.MouseButton) bool {
	if m.buttonPressed(button) {
		return false
	}
	return m.duration(button) == 1
}

func (m *virtualMouse) pressDuration(button ebiten.MouseButton) int {
	if !m.buttonPressed(button) {
		return 0
	}
	return m.duration(button)
}

func (m *virtualMouse) duration(button ebiten.MouseButton) int {
	interval := m.intervals[button]
	return m.count%interval + 1
}

func (m *virtualMouse) Wheel() (xoff, yoff float64) {
	quotient := m.count / m.wheelInterval
	if quotient%10 == 1 {
		return 0.1, 0.1
	}
	return 0, 0
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}
