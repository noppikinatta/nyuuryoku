package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GamepadConnection struct {
	appendIDsFn              func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID
	appendJustConnectedIDsFn func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID
	justDisconnectedFn       func(gamepadID ebiten.GamepadID) bool
}

func NewGamepadConnection() *GamepadConnection {
	c := &GamepadConnection{}
	s := GamepadConnectionSetter{GamepadConnection: c}

	s.SetAppendIDsFunc(ebiten.AppendGamepadIDs)
	s.SetAppendJustConnectedIDsFunc(inpututil.AppendJustConnectedGamepadIDs)

	return c
}

func (c *GamepadConnection) AppendIDs(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID {
	return c.appendIDsFn(gamepadIDs)
}

func (c *GamepadConnection) AppendJustConnectedIDs(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID {
	return c.appendJustConnectedIDsFn(gamepadIDs)
}

func (c *GamepadConnection) JustDisconnected(gamepadID ebiten.GamepadID) bool {
	return c.justDisconnectedFn(gamepadID)
}

type GamepadConnectionSetter struct {
	GamepadConnection *GamepadConnection
}

func (s *GamepadConnectionSetter) SetAppendIDsFunc(appendIDsFn func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID) {
	s.GamepadConnection.appendIDsFn = appendIDsFn
}

func (s *GamepadConnectionSetter) SetAppendJustConnectedIDsFunc(appendJustConnectedIDsFn func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID) {
	s.GamepadConnection.appendJustConnectedIDsFn = appendJustConnectedIDsFn
}

type gamepadGeneric[TButton ebiten.GamepadButton | ebiten.StandardGamepadButton, TAxis ebiten.GamepadAxisType | ebiten.StandardGamepadAxis] struct {
	ID                   ebiten.GamepadID
	sdlidFn              func(id ebiten.GamepadID) string
	nameFn               func(id ebiten.GamepadID) string
	axisCountFn          func(id ebiten.GamepadID) int
	axisValueFn          func(id ebiten.GamepadID, axis TAxis) float64
	buttonCountFn        func(id ebiten.GamepadID) int
	pressedFn            func(id ebiten.GamepadID, button TButton) bool
	justPressedFn        func(id ebiten.GamepadID, button TButton) bool
	justReleaseedFn      func(id ebiten.GamepadID, button TButton) bool
	pressedDurationFn    func(id ebiten.GamepadID, button TButton) int
	appendPressedFn      func(id ebiten.GamepadID, buttons []TButton) []TButton
	appendJustPressedFn  func(id ebiten.GamepadID, buttons []TButton) []TButton
	appendJustReleasedFn func(id ebiten.GamepadID, buttons []TButton) []TButton
}

func (g *gamepadGeneric[TButton, TAxis]) SDLID() string {
	return g.sdlidFn(g.ID)
}

func (g *gamepadGeneric[TButton, TAxis]) Name() string {
	return g.nameFn(g.ID)
}

func (g *gamepadGeneric[TButton, TAxis]) AxisCount() int {
	return g.axisCountFn(g.ID)
}

func (g *gamepadGeneric[TButton, TAxis]) AxisValue(axis TAxis) float64 {
	return g.axisValueFn(g.ID, axis)
}

func (g *gamepadGeneric[TButton, TAxis]) ButtonCount() int {
	return g.buttonCountFn(g.ID)
}

func (g *gamepadGeneric[TButton, TAxis]) Pressed(button TButton) bool {
	return g.pressedFn(g.ID, button)
}

func (g *gamepadGeneric[TButton, TAxis]) JustPressed(button TButton) bool {
	return g.justPressedFn(g.ID, button)
}

func (g *gamepadGeneric[TButton, TAxis]) JustReleased(button TButton) bool {
	return g.justReleaseedFn(g.ID, button)
}

func (g *gamepadGeneric[TButton, TAxis]) PressedDuration(button TButton) int {
	return g.pressedDurationFn(g.ID, button)
}

func (g *gamepadGeneric[TButton, TAxis]) AppendPressed(buttons []TButton) []TButton {
	return g.appendPressedFn(g.ID, buttons)
}

func (g *gamepadGeneric[TButton, TAxis]) AppendJustPressed(buttons []TButton) []TButton {
	return g.appendJustPressedFn(g.ID, buttons)
}

func (g *gamepadGeneric[TButton, TAxis]) AppendJustReleased(buttons []TButton) []TButton {
	return g.appendJustReleasedFn(g.ID, buttons)
}

type gamepadGenericSetter[TButton ebiten.GamepadButton | ebiten.StandardGamepadButton, TAxis ebiten.GamepadAxisType | ebiten.StandardGamepadAxis] struct {
	gamepadGeneric *gamepadGeneric[TButton, TAxis]
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetSDLIDFunc(sdlidFn func(id ebiten.GamepadID) string) {
	s.gamepadGeneric.sdlidFn = sdlidFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetNameFunc(nameFn func(id ebiten.GamepadID) string) {
	s.gamepadGeneric.nameFn = nameFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetAxisCountFunc(axisCountFn func(id ebiten.GamepadID) int) {
	s.gamepadGeneric.axisCountFn = axisCountFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetAxisValueFunc(axisValueFn func(id ebiten.GamepadID, axis TAxis) float64) {
	s.gamepadGeneric.axisValueFn = axisValueFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetButtonCountFunc(buttonCountFn func(id ebiten.GamepadID) int) {
	s.gamepadGeneric.buttonCountFn = buttonCountFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetPressedFunc(pressedFn func(id ebiten.GamepadID, button TButton) bool) {
	s.gamepadGeneric.pressedFn = pressedFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetJustPressedFunc(justPressedFn func(id ebiten.GamepadID, button TButton) bool) {
	s.gamepadGeneric.justPressedFn = justPressedFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetJustReleaseedFunc(justReleaseedFn func(id ebiten.GamepadID, button TButton) bool) {
	s.gamepadGeneric.justReleaseedFn = justReleaseedFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetPressedDurationFunc(pressedDurationFn func(id ebiten.GamepadID, button TButton) int) {
	s.gamepadGeneric.pressedDurationFn = pressedDurationFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetAppendPressedFunc(appendPressedFn func(id ebiten.GamepadID, buttons []TButton) []TButton) {
	s.gamepadGeneric.appendPressedFn = appendPressedFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetAppendJustPressedFunc(appendJustPressedFn func(id ebiten.GamepadID, buttons []TButton) []TButton) {
	s.gamepadGeneric.appendJustPressedFn = appendJustPressedFn
}

func (s *gamepadGenericSetter[TButton, TAxis]) SetAppendJustReleasedFunc(appendJustReleasedFn func(id ebiten.GamepadID, buttons []TButton) []TButton) {
	s.gamepadGeneric.appendJustReleasedFn = appendJustReleasedFn
}

type Gamepad struct {
	gamepadGeneric[ebiten.GamepadButton, ebiten.GamepadAxisType]
}

func NewGamepad(id ebiten.GamepadID) *Gamepad {
	g := &Gamepad{
		gamepadGeneric: gamepadGeneric[ebiten.GamepadButton, ebiten.GamepadAxisType]{
			ID: id,
		},
	}
	s := NewGamepadSetter(g)

	s.SetSDLIDFunc(ebiten.GamepadSDLID)
	s.SetNameFunc(ebiten.GamepadName)
	s.SetAxisCountFunc(ebiten.GamepadAxisCount)
	s.SetAxisValueFunc(ebiten.GamepadAxisValue)
	s.SetButtonCountFunc(ebiten.GamepadButtonCount)
	s.SetPressedFunc(ebiten.IsGamepadButtonPressed)
	s.SetJustPressedFunc(inpututil.IsGamepadButtonJustPressed)
	s.SetJustReleaseedFunc(inpututil.IsGamepadButtonJustReleased)
	s.SetPressedDurationFunc(inpututil.GamepadButtonPressDuration)
	s.SetAppendPressedFunc(inpututil.AppendPressedGamepadButtons)
	s.SetAppendJustPressedFunc(inpututil.AppendJustPressedGamepadButtons)
	s.SetAppendJustReleasedFunc(inpututil.AppendJustReleasedGamepadButtons)

	return g
}

type GamepadSetter struct {
	gamepadGenericSetter[ebiten.GamepadButton, ebiten.GamepadAxisType]
}

func NewGamepadSetter(g *Gamepad) *GamepadSetter {
	s := GamepadSetter{}
	s.gamepadGeneric = &g.gamepadGeneric
	return &s
}

type StandardGamepad struct {
	gamepadGeneric[ebiten.StandardGamepadButton, ebiten.StandardGamepadAxis]
	layoutAvailableFn func(id ebiten.GamepadID) bool
	axisAvailableFn   func(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) bool
	buttonAvailableFn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
}

func NewStandardGamepad(id ebiten.GamepadID) *StandardGamepad {
	g := &StandardGamepad{
		gamepadGeneric: gamepadGeneric[ebiten.StandardGamepadButton, ebiten.StandardGamepadAxis]{
			ID: id,
		},
	}
	s := NewStandardgamepadSetter(g)

	s.SetSDLIDFunc(ebiten.GamepadSDLID)
	s.SetNameFunc(ebiten.GamepadName)
	s.SetAxisCountFunc(ebiten.GamepadAxisCount)
	s.SetAxisValueFunc(ebiten.StandardGamepadAxisValue)
	s.SetButtonCountFunc(ebiten.GamepadButtonCount)
	s.SetPressedFunc(ebiten.IsStandardGamepadButtonPressed)
	s.SetJustPressedFunc(inpututil.IsStandardGamepadButtonJustPressed)
	s.SetJustReleaseedFunc(inpututil.IsStandardGamepadButtonJustReleased)
	s.SetPressedDurationFunc(inpututil.StandardGamepadButtonPressDuration)
	s.SetAppendPressedFunc(inpututil.AppendPressedStandardGamepadButtons)
	s.SetAppendJustPressedFunc(inpututil.AppendJustPressedStandardGamepadButtons)
	s.SetAppendJustReleasedFunc(inpututil.AppendJustReleasedStandardGamepadButtons)
	s.SetAvailableFunc(ebiten.IsStandardGamepadLayoutAvailable)
	s.SetAxisAvailableFunc(ebiten.IsStandardGamepadAxisAvailable)
	s.SetButtonAvailableFunc(ebiten.IsStandardGamepadButtonAvailable)

	return g
}

func (g *StandardGamepad) Available() bool {
	return g.layoutAvailableFn(g.ID)
}

func (g *StandardGamepad) AxisAvailable(axis ebiten.StandardGamepadAxis) bool {
	return g.axisAvailableFn(g.ID, axis)
}

func (g *StandardGamepad) ButtonAvailable(button ebiten.StandardGamepadButton) bool {
	return g.buttonAvailableFn(g.ID, button)
}

type StandardGamepadSetter struct {
	gamepadGenericSetter[ebiten.StandardGamepadButton, ebiten.StandardGamepadAxis]
	standardGamepad *StandardGamepad
}

func NewStandardgamepadSetter(g *StandardGamepad) *StandardGamepadSetter {
	s := StandardGamepadSetter{}
	s.gamepadGeneric = &g.gamepadGeneric
	s.standardGamepad = g
	return &s
}

func (s *StandardGamepadSetter) SetAvailableFunc(layoutAvailableFn func(id ebiten.GamepadID) bool) {
	s.standardGamepad.layoutAvailableFn = layoutAvailableFn
}

func (s *StandardGamepadSetter) SetAxisAvailableFunc(axisAvailableFn func(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) bool) {
	s.standardGamepad.axisAvailableFn = axisAvailableFn
}

func (s *StandardGamepadSetter) SetButtonAvailableFunc(buttonAvailableFn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool) {
	s.standardGamepad.buttonAvailableFn = buttonAvailableFn
}
