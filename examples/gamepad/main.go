package main

import (
	"fmt"
	"image"
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
	conn           *nyuuryoku.GamepadConnection
	gamepad        *nyuuryoku.StandardGamepad
	connIsVirtual  bool
	virtualConn    *virtualGamepadConnection
	connectedIDs   []ebiten.GamepadID
	names          map[ebiten.GamepadID]string
	pressedButtons map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	axisValues     map[ebiten.GamepadID][]ebiten.StandardGamepadAxis
}

func (g *game) Update() error {
	if g.conn == nil {
		g.conn = nyuuryoku.NewGamepadConnection()
	}
	if g.gamepad == nil {
		g.gamepad = nyuuryoku.NewStandardGamepad(0)
	}
	if g.virtualConn == nil {
		g.virtualConn = newVirtualGamepadConnection()
	}
	if g.names == nil {
		g.names = make(map[ebiten.GamepadID]string)
	}
	if g.pressedButtons == nil {
		g.pressedButtons = make(map[ebiten.GamepadID][]ebiten.StandardGamepadButton)
	}
	if g.axisValues == nil {
		g.axisValues = make(map[ebiten.GamepadID][]ebiten.StandardGamepadAxis)
	}

	if g.connIsVirtual {
		g.virtualConn.Update()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.switchGamepad()
	}

	g.updatePressed()

	return nil
}

func (g *game) updatePressed() {
	g.connectedIDs = g.conn.AppendIDs(g.connectedIDs)

	for _, id := range g.connectedIDs {
		g.gamepad.ID = id
		g.names[id] = g.gamepad.Name()

		buttons, ok := g.pressedButtons[id]
		if !ok {
			buttons = make([]ebiten.StandardGamepadButton, 0)
		}
		g.pressedButtons[id] = g.gamepad.AppendPressed(buttons)

		axisValues, ok := g.axisValues[id]
		if !ok {
			axisValues = make([]ebiten.StandardGamepadAxis, 0)
		}
		g.axisValues[id] = g.gamepad.AxisValue(ebiten.s)

	}
}

func (g *game) buttonName(button ebiten.MouseButton) string {
	switch button {
	case ebiten.MouseButtonLeft:
		return "left"
	case ebiten.MouseButtonRight:
		return "right"
	case ebiten.MouseButtonMiddle:
		return "middle"
	}

	return "other button"
}

func (g *game) appendWheelLog() {
	xoff, yoff := g.mouse.Wheel()
	if xoff == 0 && yoff == 0 {
		return
	}
	g.appendLog(fmt.Sprintf("wheel: %0.1f, %0.1f", xoff, yoff))
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

func (g *game) switchGamepad() {
	setter := nyuuryoku.NewGamepadConnectionSetter(g.conn)

	if g.connIsVirtual {
		setter.SetDefault()
	} else {
		setter.SetAppendIDsFunc(g.virtualConn.appendIDs)
	}

	g.connIsVirtual = !g.connIsVirtual
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 48})
	if g.mouse == nil {
		return
	}

	g.drawButton(screen, image.Rect(16, 16, 80, 144), ebiten.MouseButtonLeft)
	g.drawButton(screen, image.Rect(88, 16, 112, 80), ebiten.MouseButtonMiddle)
	g.drawButton(screen, image.Rect(120, 16, 184, 144), ebiten.MouseButtonRight)

	g.drawCursor(screen)

	g.drawLog(screen)

	ebitenutil.DebugPrintAt(screen, "press space to switch to virtual mouse", 10, 576)
}

func (g *game) drawButton(screen *ebiten.Image, bounds image.Rectangle, button ebiten.MouseButton) {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Scale(float64(bounds.Dx()), float64(bounds.Dy()))
	opt.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
	if g.mouse.Pressed(button) {
		opt.ColorScale.SetG(0.5)
		opt.ColorScale.SetB(0.5)
	}

	screen.DrawImage(dummyWhitePixel, &opt)
}

func (g *game) drawCursor(screen *ebiten.Image) {
	x, y := g.mouse.CursorPosition()

	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Scale(9, 9)
	opt.GeoM.Translate(float64(x-4), float64(y-4))

	screen.DrawImage(dummyWhitePixel, &opt)
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

type virtualGamepadConnection struct {
}

func newVirtualGamepadConnection() *virtualGamepadConnection {
	return &virtualGamepadConnection{}
}

func (g *virtualGamepadConnection) Update() {

}

func (g *virtualGamepadConnection) appendIDs(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID {

}

type virtualStandardGamepad struct {
}

func (g *virtualStandardGamepad) name() string {

}

func (g *virtualStandardGamepad) axisValue(axis ebiten.StandardGamepadAxis) float64 {

}

func (g *virtualStandardGamepad) appendPressed(button []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton {

}

func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

var (
	dummyImageBase = ebiten.NewImage(3, 3)

	// dummyWhitePixel is a 1x1 white pixel image.
	dummyWhitePixel = dummyImageBase.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	dummyImageBase.Fill(color.White)
}
