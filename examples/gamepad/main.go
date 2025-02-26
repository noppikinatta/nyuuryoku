package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/noppikinatta/nyuuryoku"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	gamepad          *nyuuryoku.Gamepad
	log              []string
	logIdx           int
	pressedDurations map[ebiten.GamepadID]map[ebiten.GamepadButton]int
}

func NewGame() *Game {
	return &Game{
		gamepad:          nyuuryoku.NewGamepad(),
		pressedDurations: make(map[ebiten.GamepadID]map[ebiten.GamepadButton]int),
	}
}

func (g *Game) Update() error {
	// Get connected gamepad IDs
	var ids []ebiten.GamepadID
	ids = g.gamepad.AppendIDs(ids)

	// Check for newly connected gamepads
	var justConnectedIDs []ebiten.GamepadID
	justConnectedIDs = g.gamepad.AppendJustConnectedIDs(justConnectedIDs)
	for _, id := range justConnectedIDs {
		g.appendLog(fmt.Sprintf("New gamepad connected: %s (%s)", g.gamepad.Name(id), g.gamepad.SDLID(id)))
		g.pressedDurations[id] = make(map[ebiten.GamepadButton]int)
	}

	// Check for disconnected gamepads
	for id := range g.pressedDurations {
		if g.gamepad.IsJustDisconnected(id) {
			g.appendLog(fmt.Sprintf("Gamepad disconnected: %d", id))
			delete(g.pressedDurations, id)
		}
	}

	// Log gamepad states
	for _, id := range ids {
		// Log axis values when they change significantly
		axisCount := g.gamepad.AxisCount(id)
		for a := 0; a < axisCount; a++ {
			value := g.gamepad.AxisValue(id, a)
			if value > 0.5 || value < -0.5 {
				g.appendLog(fmt.Sprintf("Gamepad %d: Axis %d = %.2f", id, a, value))
			}
		}

		// Log button states
		if g.gamepad.IsStandardLayoutAvailable(id) {
			// For standard layout
			buttons := []ebiten.StandardGamepadButton{
				ebiten.StandardGamepadButtonRightBottom,
				ebiten.StandardGamepadButtonRightRight,
				ebiten.StandardGamepadButtonRightTop,
				ebiten.StandardGamepadButtonRightLeft,
			}
			for _, b := range buttons {
				if g.gamepad.IsStandardButtonJustPressed(id, b) {
					g.appendLog(fmt.Sprintf("Gamepad %d: Standard Button %d just pressed", id, b))
				}
				if g.gamepad.IsStandardButtonJustReleased(id, b) {
					g.appendLog(fmt.Sprintf("Gamepad %d: Standard Button %d just released", id, b))
				}
			}
		} else {
			// For custom layout
			buttonCount := g.gamepad.ButtonCount(id)
			for b := ebiten.GamepadButton(0); b < ebiten.GamepadButton(buttonCount); b++ {
				if g.gamepad.IsButtonJustPressed(id, b) {
					g.appendLog(fmt.Sprintf("Gamepad %d: Button %d just pressed", id, b))
				}
				if g.gamepad.IsButtonJustReleased(id, b) {
					g.appendLog(fmt.Sprintf("Gamepad %d: Button %d just released", id, b))
				}
			}
		}
	}

	return nil
}

func (g *Game) appendLog(line string) {
	const maxLines = 20
	if len(g.log) < maxLines {
		g.log = append(g.log, line)
	} else {
		g.log[g.logIdx] = line
		g.logIdx++
		g.logIdx = g.logIdx % maxLines
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Display basic information
	var ids []ebiten.GamepadID
	ids = g.gamepad.AppendIDs(ids)

	y := 16
	if len(ids) == 0 {
		ebitenutil.DebugPrintAt(screen, "No gamepad is connected.", 16, y)
		y += 16
	} else {
		for _, id := range ids {
			name := g.gamepad.Name(id)
			isStandard := g.gamepad.IsStandardLayoutAvailable(id)
			layoutStr := "Standard"
			if !isStandard {
				layoutStr = "Custom"
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Gamepad %d: %s (%s)", id, name, layoutStr), 16, y)
			y += 16
		}
	}

	y += 16

	// Display operation log
	ebitenutil.DebugPrintAt(screen, "Operation Log:", 16, y)
	y += 16

	for i := range g.log {
		line := g.log[i]
		yIdx := (i - g.logIdx)
		if yIdx < 0 {
			yIdx += len(g.log)
		}
		ebitenutil.DebugPrintAt(screen, line, 16, y+(yIdx*16))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gamepad Example")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
