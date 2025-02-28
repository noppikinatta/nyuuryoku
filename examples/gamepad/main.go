package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	gamepad          *nyuuryoku.Gamepad
	gamepadStates    map[ebiten.GamepadID]*GamepadState
	log              []LogEntry
	logIdx           int
	gamepadIsVirtual bool
	virtualGamepad   *virtualGamepad
}

func NewGame() *Game {
	return &Game{
		gamepad:          nyuuryoku.NewGamepad(),
		gamepadStates:    make(map[ebiten.GamepadID]*GamepadState),
		log:              make([]LogEntry, 0, 20),
		gamepadIsVirtual: false,
		virtualGamepad:   newVirtualGamepad(),
	}
}

func (g *Game) Update() error {
	// Switch between real and virtual gamepad with mouse click
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.switchGamepad()
	}

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
			for b := ebiten.StandardGamepadButton(0); b < ebiten.StandardGamepadButtonMax; b++ {
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

func (g *Game) switchGamepad() {
	setter := nyuuryoku.NewGamepadSetter(g.gamepad)

	if g.gamepadIsVirtual {
		setter.SetDefault()
		g.appendLog("Switched to real gamepad")
	} else {
		// Only set functions for the virtual gamepad ID
		setter.SetAppendIDsFunc(g.virtualGamepad.appendIDs)
		setter.SetAppendJustConnectedIDsFunc(g.virtualGamepad.appendJustConnectedIDs)
		setter.SetIsJustDisconnectedFunc(g.virtualGamepad.isJustDisconnected)
		setter.SetNameFunc(g.virtualGamepad.name)
		setter.SetSDLIDFunc(g.virtualGamepad.sdlID)
		setter.SetIsStandardLayoutAvailableFunc(g.virtualGamepad.isStandardLayoutAvailable)
		setter.SetAxisCountFunc(g.virtualGamepad.axisCount)
		setter.SetAxisValueFunc(g.virtualGamepad.axisValue)
		setter.SetButtonCountFunc(g.virtualGamepad.buttonCount)
		setter.SetIsButtonPressedFunc(g.virtualGamepad.isButtonPressed)
		setter.SetIsButtonJustPressedFunc(g.virtualGamepad.isButtonJustPressed)
		setter.SetIsButtonJustReleasedFunc(g.virtualGamepad.isButtonJustReleased)
		setter.SetIsStandardButtonPressedFunc(g.virtualGamepad.isStandardButtonPressed)
		setter.SetIsStandardButtonJustPressedFunc(g.virtualGamepad.isStandardButtonJustPressed)
		setter.SetIsStandardButtonJustReleasedFunc(g.virtualGamepad.isStandardButtonJustReleased)

		g.appendLog("Switched to virtual gamepad")
	}

	g.gamepadIsVirtual = !g.gamepadIsVirtual

	// Clear state to start fresh with the new input source
	g.gamepadStates = make(map[ebiten.GamepadID]*GamepadState)
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

	// Display mode information
	if g.gamepadIsVirtual {
		ebitenutil.DebugPrintAt(screen, "VIRTUAL GAMEPAD MODE", 16, y)
	} else {
		ebitenutil.DebugPrintAt(screen, "REAL GAMEPAD MODE", 16, y)
	}
	y += 24

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

	ebitenutil.DebugPrintAt(screen, "left click to switch input source", 16, screenHeight-24)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// buttonState tracks the state of a gamepad button
type buttonState struct {
	StateChangedCount int
	Pressed           bool
}

// virtualGamepad implements a simulated gamepad
type virtualGamepad struct {
	ID                   ebiten.GamepadID
	frame                int
	connected            bool
	justConnected        bool
	justDisconnected     bool
	standardLayout       bool
	buttonStates         map[ebiten.GamepadButton]buttonState
	standardButtonStates map[ebiten.StandardGamepadButton]buttonState
	axisValues           []float64
}

func newVirtualGamepad() *virtualGamepad {
	vg := &virtualGamepad{
		ID:                   42, // Virtual ID
		connected:            true,
		justConnected:        true,
		standardLayout:       true,
		buttonStates:         make(map[ebiten.GamepadButton]buttonState),
		standardButtonStates: make(map[ebiten.StandardGamepadButton]buttonState),
		axisValues:           make([]float64, 4), // Four axes
	}
	return vg
}

func (v *virtualGamepad) Update() {
	v.frame++

	// Reset just connected/disconnected after first frame
	if v.frame > 1 {
		v.justConnected = false
		v.justDisconnected = false
	}

	// Update axis values with some animation
	for i := range v.axisValues {
		period := 120 + i*30 // Different period for each axis
		amplitude := 0.8
		// Sinusoidal-like movement
		if i%2 == 0 {
			v.axisValues[i] = amplitude * float64((v.frame/period)%2*2-1) * float64(v.frame%period) / float64(period)
		} else {
			v.axisValues[i] = amplitude * float64((v.frame/period)%2*2-1) * (1.0 - float64(v.frame%period)/float64(period))
		}
	}

	// Update standard button states
	for btn := ebiten.StandardGamepadButton(0); btn < ebiten.StandardGamepadButtonMax; btn++ {
		state := v.standardButtonStates[btn]

		// Different pattern for each button
		cycleLength := 60 + int(btn)*10
		newPressed := ((v.frame + int(btn)*15) % cycleLength) < cycleLength/3

		if state.Pressed != newPressed {
			state.Pressed = newPressed
			state.StateChangedCount = v.frame
			v.standardButtonStates[btn] = state
		}
	}

	// Update regular button states
	for btn := ebiten.GamepadButton(0); btn < 12; btn++ {
		state := v.buttonStates[btn]

		// Different pattern for each button
		cycleLength := 90 + int(btn)*12
		newPressed := ((v.frame + int(btn)*20) % cycleLength) < cycleLength/4

		if state.Pressed != newPressed {
			state.Pressed = newPressed
			state.StateChangedCount = v.frame
			v.buttonStates[btn] = state
		}
	}
}

// Implementation of gamepad functions
func (v *virtualGamepad) appendIDs(ids []ebiten.GamepadID) []ebiten.GamepadID {
	v.Update() // Update virtual state
	if v.connected {
		return append(ids, v.ID)
	}
	return ids
}

func (v *virtualGamepad) appendJustConnectedIDs(ids []ebiten.GamepadID) []ebiten.GamepadID {
	if v.justConnected {
		return append(ids, v.ID)
	}
	return ids
}

func (v *virtualGamepad) isJustDisconnected(id ebiten.GamepadID) bool {
	return id == v.ID && v.justDisconnected
}

func (v *virtualGamepad) name(id ebiten.GamepadID) string {
	if id == v.ID {
		return "Virtual Gamepad"
	}
	return ""
}

func (v *virtualGamepad) sdlID(id ebiten.GamepadID) string {
	if id == v.ID {
		return "VIRTUAL-SDLID-X123"
	}
	return ""
}

func (v *virtualGamepad) isStandardLayoutAvailable(id ebiten.GamepadID) bool {
	if id == v.ID {
		return v.standardLayout
	}
	return false
}

func (v *virtualGamepad) axisCount(id ebiten.GamepadID) int {
	if id == v.ID {
		return len(v.axisValues)
	}
	return 0
}

func (v *virtualGamepad) axisValue(id ebiten.GamepadID, axis int) float64 {
	if id == v.ID && axis >= 0 && axis < len(v.axisValues) {
		return v.axisValues[axis]
	}
	return 0
}

func (v *virtualGamepad) buttonCount(id ebiten.GamepadID) int {
	if id == v.ID {
		return 12 // Fixed number of buttons
	}
	return 0
}

func (v *virtualGamepad) isButtonPressed(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	if id == v.ID {
		state, exists := v.buttonStates[button]
		return exists && state.Pressed
	}
	return false
}

func (v *virtualGamepad) isButtonJustPressed(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	if id == v.ID {
		state, exists := v.buttonStates[button]
		return exists && state.Pressed && state.StateChangedCount == v.frame
	}
	return false
}

func (v *virtualGamepad) isButtonJustReleased(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	if id == v.ID {
		state, exists := v.buttonStates[button]
		return exists && !state.Pressed && state.StateChangedCount == v.frame
	}
	return false
}

func (v *virtualGamepad) isStandardButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	if id == v.ID {
		state, exists := v.standardButtonStates[button]
		return exists && state.Pressed
	}
	return false
}

func (v *virtualGamepad) isStandardButtonJustPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	if id == v.ID {
		state, exists := v.standardButtonStates[button]
		return exists && state.Pressed && state.StateChangedCount == v.frame
	}
	return false
}

func (v *virtualGamepad) isStandardButtonJustReleased(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	if id == v.ID {
		state, exists := v.standardButtonStates[button]
		return exists && !state.Pressed && state.StateChangedCount == v.frame
	}
	return false
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gamepad Example")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
