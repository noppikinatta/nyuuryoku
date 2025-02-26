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

// GamepadState stores the state of a gamepad
type GamepadState struct {
	ID                   ebiten.GamepadID
	Name                 string
	SDLID                string
	IsStandard           bool
	Connected            bool
	JustConnected        bool
	JustDisconnected     bool
	AxisValues           []float64
	ButtonStates         map[ebiten.GamepadButton]bool
	StandardButtonStates map[ebiten.StandardGamepadButton]bool
}

// LogEntry represents a single log entry
type LogEntry struct {
	Message string
}

type Game struct {
	gamepad       *nyuuryoku.Gamepad
	gamepadStates map[ebiten.GamepadID]*GamepadState
	log           []LogEntry
	logIdx        int
}

func NewGame() *Game {
	return &Game{
		gamepad:       nyuuryoku.NewGamepad(),
		gamepadStates: make(map[ebiten.GamepadID]*GamepadState),
		log:           make([]LogEntry, 0, 20),
	}
}

func (g *Game) Update() error {
	// Get connected gamepad IDs
	var ids []ebiten.GamepadID
	ids = g.gamepad.AppendIDs(ids)

	// Mark all gamepads as not connected initially
	for _, state := range g.gamepadStates {
		state.Connected = false
		state.JustConnected = false
		state.JustDisconnected = false
	}

	// Update states for connected gamepads
	for _, id := range ids {
		state, exists := g.gamepadStates[id]
		if !exists {
			// Create new state for newly connected gamepad
			state = &GamepadState{
				ID:                   id,
				Name:                 g.gamepad.Name(id),
				SDLID:                g.gamepad.SDLID(id),
				IsStandard:           g.gamepad.IsStandardLayoutAvailable(id),
				Connected:            true,
				JustConnected:        true,
				ButtonStates:         make(map[ebiten.GamepadButton]bool),
				StandardButtonStates: make(map[ebiten.StandardGamepadButton]bool),
			}
			g.gamepadStates[id] = state
			g.appendLog(fmt.Sprintf("New gamepad connected: %s (%s)", state.Name, state.SDLID))
		} else {
			// Update existing gamepad state
			state.Connected = true
			state.Name = g.gamepad.Name(id)
			state.SDLID = g.gamepad.SDLID(id)
			state.IsStandard = g.gamepad.IsStandardLayoutAvailable(id)
		}

		// Update axis values
		axisCount := g.gamepad.AxisCount(id)
		if len(state.AxisValues) != axisCount {
			state.AxisValues = make([]float64, axisCount)
		}

		for a := 0; a < axisCount; a++ {
			newValue := g.gamepad.AxisValue(id, a)
			oldValue := state.AxisValues[a]
			state.AxisValues[a] = newValue

			// Log significant changes in axis values
			if (newValue > 0.5 && oldValue <= 0.5) || (newValue < -0.5 && oldValue >= -0.5) {
				g.appendLog(fmt.Sprintf("Gamepad %d: Axis %d = %.2f", id, a, newValue))
			}
		}

		// Update button states
		if state.IsStandard {
			// For standard layout
			buttons := []ebiten.StandardGamepadButton{
				ebiten.StandardGamepadButtonRightBottom,
				ebiten.StandardGamepadButtonRightRight,
				ebiten.StandardGamepadButtonRightTop,
				ebiten.StandardGamepadButtonRightLeft,
			}

			for _, b := range buttons {
				oldPressed := state.StandardButtonStates[b]
				newPressed := g.gamepad.IsStandardButtonPressed(id, b)
				state.StandardButtonStates[b] = newPressed

				if newPressed && !oldPressed {
					g.appendLog(fmt.Sprintf("Gamepad %d: Standard Button %d just pressed", id, b))
				} else if !newPressed && oldPressed {
					g.appendLog(fmt.Sprintf("Gamepad %d: Standard Button %d just released", id, b))
				}
			}
		} else {
			// For custom layout
			buttonCount := g.gamepad.ButtonCount(id)
			for b := ebiten.GamepadButton(0); b < ebiten.GamepadButton(buttonCount); b++ {
				oldPressed := state.ButtonStates[b]
				newPressed := g.gamepad.IsButtonPressed(id, b)
				state.ButtonStates[b] = newPressed

				if newPressed && !oldPressed {
					g.appendLog(fmt.Sprintf("Gamepad %d: Button %d just pressed", id, b))
				} else if !newPressed && oldPressed {
					g.appendLog(fmt.Sprintf("Gamepad %d: Button %d just released", id, b))
				}
			}
		}
	}

	// Check for disconnected gamepads
	for id, state := range g.gamepadStates {
		if !state.Connected && g.gamepad.IsJustDisconnected(id) {
			state.JustDisconnected = true
			g.appendLog(fmt.Sprintf("Gamepad disconnected: %d", id))
		}
	}

	return nil
}

func (g *Game) appendLog(message string) {
	const maxLines = 20
	entry := LogEntry{Message: message}

	if len(g.log) < maxLines {
		g.log = append(g.log, entry)
	} else {
		g.log[g.logIdx] = entry
		g.logIdx++
		g.logIdx = g.logIdx % maxLines
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	y := 16

	// Display connected gamepads
	connectedCount := 0
	for _, state := range g.gamepadStates {
		if state.Connected {
			connectedCount++
		}
	}

	if connectedCount == 0 {
		ebitenutil.DebugPrintAt(screen, "No gamepad is connected.", 16, y)
		y += 16
	} else {
		for _, state := range g.gamepadStates {
			if !state.Connected {
				continue
			}

			layoutStr := "Standard"
			if !state.IsStandard {
				layoutStr = "Custom"
			}

			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Gamepad %d: %s (%s)", state.ID, state.Name, layoutStr), 16, y)
			y += 16
		}
	}

	y += 16

	// Display operation log
	ebitenutil.DebugPrintAt(screen, "Operation Log:", 16, y)
	y += 16

	for i := range g.log {
		yIdx := (i - g.logIdx)
		if yIdx < 0 {
			yIdx += len(g.log)
		}
		ebitenutil.DebugPrintAt(screen, g.log[i].Message, 16, y+(yIdx*16))
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
