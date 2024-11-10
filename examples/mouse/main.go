package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/noppikinatta/nyuuryoku"
)

func main() {
	ebiten.SetWindowSize(800, 600)

	g := game{}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

type game struct {
	mouse            *nyuuryoku.Mouse
	keyboard         *nyuuryoku.Keyboard
	log              []string
	logIdx           int
	pressedDurations map[ebiten.MouseButton]int
}

func (g *game) Update() error {
	if g.mouse == nil {
		g.mouse = nyuuryoku.NewMouse()
		g.pressedDurations = make(map[ebiten.MouseButton]int)
	}
	if g.keyboard == nil {
		g.keyboard = nyuuryoku.NewKeyboard()
	}

	g.appendButtonLog(ebiten.MouseButtonLeft)
	g.appendButtonLog(ebiten.MouseButtonRight)
	g.appendButtonLog(ebiten.MouseButtonMiddle)
	g.appendWheelLog()

	return nil
}

func (g *game) appendButtonLog(button ebiten.MouseButton) {
	buttonName := g.buttonName(button)

	if g.mouse.JustPressed(button) {
		g.appendLog(fmt.Sprintf("just pressed: %s", buttonName))
	}
	if g.mouse.JustReleased(button) {
		g.appendLog(fmt.Sprintf("just released: %s", buttonName))
		g.appendLog(fmt.Sprintf("%6s pressed durarion: %d", buttonName, g.pressedDurations[button]))
	}
	if g.mouse.Pressed(button) {
		g.pressedDurations[button] = g.mouse.PressDuration(button)
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
	const max = 50
	if len(g.log) < max {
		g.log = append(g.log, line)
	} else {
		g.log[g.logIdx] = line
		g.logIdx++
		g.logIdx = g.logIdx % max
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 48})
	if g.mouse == nil {
		return
	}

	g.drawButton(screen, image.Rect(16, 16, 80, 144), ebiten.MouseButtonLeft)
	g.drawButton(screen, image.Rect(88, 16, 112, 80), ebiten.MouseButtonMiddle)
	g.drawButton(screen, image.Rect(120, 16, 184, 144), ebiten.MouseButtonRight)

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
