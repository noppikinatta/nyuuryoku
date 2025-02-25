package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/noppikinatta/nyuuryoku"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	gamepad   *nyuuryoku.Gamepad
	mplusFont font.Face
}

func NewGame() (*Game, error) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}

	mplusFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	return &Game{
		gamepad:   nyuuryoku.NewGamepad(),
		mplusFont: mplusFont,
	}, nil
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
		text.Draw(screen, "No gamepad is connected.", g.mplusFont, 16, y, color.White)
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

		text.Draw(screen, fmt.Sprintf("Gamepad %d: %s (%s)", id, name, layoutStr), g.mplusFont, 16, y, color.White)
		y += 20

		// Display axis values
		axisCount := g.gamepad.AxisCount(id)
		for a := 0; a < axisCount; a++ {
			value := g.gamepad.AxisValue(id, a)
			text.Draw(screen, fmt.Sprintf("  Axis %d: %.2f", a, value), g.mplusFont, 16, y, color.White)
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
					text.Draw(screen, fmt.Sprintf("  Button %d: Pressed", b), g.mplusFont, 16, y, color.RGBA{0, 255, 0, 255})
				} else {
					text.Draw(screen, fmt.Sprintf("  Button %d: Released", b), g.mplusFont, 16, y, color.White)
				}
				y += 16
			}
		} else {
			// For custom layout
			buttonCount := g.gamepad.ButtonCount(id)
			for b := ebiten.GamepadButton(0); b < ebiten.GamepadButton(buttonCount); b++ {
				if g.gamepad.IsButtonPressed(id, b) {
					text.Draw(screen, fmt.Sprintf("  Button %d: Pressed", b), g.mplusFont, 16, y, color.RGBA{0, 255, 0, 255})
				} else {
					text.Draw(screen, fmt.Sprintf("  Button %d: Released", b), g.mplusFont, 16, y, color.White)
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
			text.Draw(screen, fmt.Sprintf("New gamepad connected: %d", id), g.mplusFont, 16, y, color.RGBA{0, 255, 0, 255})
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

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
