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
	gamepad *nyuuryoku.Gamepad
}

func NewGame() *Game {
	return &Game{
		gamepad: nyuuryoku.NewGamepad(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Get connected gamepad IDs
	var ids []ebiten.GamepadID
	ids = g.gamepad.AppendIDs(ids)

	y := 16
	if len(ids) == 0 {
		ebitenutil.DebugPrintAt(screen, "No gamepad is connected.", 16, y)
		return
	}

	// Display information for each gamepad
	for _, id := range ids {
		name := g.gamepad.Name(id)
		isStandard := g.gamepad.IsStandardLayoutAvailable(id)
		layoutStr := "Standard"
		if !isStandard {
			layoutStr = "Custom"
		}

		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Gamepad %d: %s (%s)", id, name, layoutStr), 16, y)
		y += 20

		// Display axis values
		axisCount := g.gamepad.AxisCount(id)
		for a := 0; a < axisCount; a++ {
			value := g.gamepad.AxisValue(id, a)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("  Axis %d: %.2f", a, value), 16, y)
			y += 16
		}

		// Display button states
		if isStandard {
			// For standard layout
			buttons := []ebiten.StandardGamepadButton{
				ebiten.StandardGamepadButtonRightBottom,
				ebiten.StandardGamepadButtonRightRight,
				ebiten.StandardGamepadButtonRightTop,
				ebiten.StandardGamepadButtonRightLeft,
			}
			for _, b := range buttons {
				if g.gamepad.IsStandardButtonPressed(id, b) {
					ebitenutil.DebugPrintAt(screen, fmt.Sprintf("  Button %d: Pressed", b), 16, y)
				} else {
					ebitenutil.DebugPrintAt(screen, fmt.Sprintf("  Button %d: Released", b), 16, y)
				}
				y += 16
			}
		} else {
			// For custom layout
			buttonCount := g.gamepad.ButtonCount(id)
			for b := ebiten.GamepadButton(0); b < ebiten.GamepadButton(buttonCount); b++ {
				if g.gamepad.IsButtonPressed(id, b) {
					ebitenutil.DebugPrintAt(screen, fmt.Sprintf("  Button %d: Pressed", b), 16, y)
				} else {
					ebitenutil.DebugPrintAt(screen, fmt.Sprintf("  Button %d: Released", b), 16, y)
				}
				y += 16
			}
		}

		y += 16 // Space between gamepads
	}

	// Display newly connected gamepads
	var justConnectedIDs []ebiten.GamepadID
	justConnectedIDs = g.gamepad.AppendJustConnectedIDs(justConnectedIDs)
	if len(justConnectedIDs) > 0 {
		for _, id := range justConnectedIDs {
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("New gamepad connected: %d", id), 16, y)
			y += 20
		}
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
