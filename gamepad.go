// CODE GENERATED BY genapis.go. DO NOT EDIT.

package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Gamepad struct {
	appendGamepadIDsFn                         func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID
	gamepadAxisCountFn                         func(id ebiten.GamepadID) int
	gamepadAxisValueFn                         func(id ebiten.GamepadID, axis int) float64
	gamepadButtonCountFn                       func(id ebiten.GamepadID) int
	gamepadNameFn                              func(id ebiten.GamepadID) string
	gamepadSDLIDFn                             func(id ebiten.GamepadID) string
	isGamepadButtonPressedFn                   func(id ebiten.GamepadID, button ebiten.GamepadButton) bool
	isStandardGamepadAxisAvailableFn           func(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) bool
	isStandardGamepadButtonAvailableFn         func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
	isStandardGamepadButtonPressedFn           func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
	isStandardGamepadLayoutAvailableFn         func(id ebiten.GamepadID) bool
	standardGamepadAxisValueFn                 func(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64
	appendJustConnectedGamepadIDsFn            func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID
	appendJustPressedGamepadButtonsFn          func(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton
	appendJustPressedStandardGamepadButtonsFn  func(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton
	appendJustReleasedGamepadButtonsFn         func(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton
	appendJustReleasedStandardGamepadButtonsFn func(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton
	appendPressedGamepadButtonsFn              func(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton
	appendPressedStandardGamepadButtonsFn      func(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton
	gamepadButtonPressDurationFn               func(id ebiten.GamepadID, button ebiten.GamepadButton) int
	isGamepadButtonJustPressedFn               func(id ebiten.GamepadID, button ebiten.GamepadButton) bool
	isGamepadButtonJustReleasedFn              func(id ebiten.GamepadID, button ebiten.GamepadButton) bool
	isGamepadJustDisconnectedFn                func(id ebiten.GamepadID) bool
	isStandardGamepadButtonJustPressedFn       func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
	isStandardGamepadButtonJustReleasedFn      func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
	standardGamepadButtonPressDurationFn       func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) int
}

func NewGamepad() *Gamepad {
	g := &Gamepad{}
	s := NewGamepadSetter(g)
	s.SetDefault()

	return g
}

func (g *Gamepad) AppendIDs(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID {
	return g.appendGamepadIDsFn(gamepadIDs)
}
func (g *Gamepad) AxisCount(id ebiten.GamepadID) int {
	return g.gamepadAxisCountFn(id)
}
func (g *Gamepad) AxisValue(id ebiten.GamepadID, axis int) float64 {
	return g.gamepadAxisValueFn(id, axis)
}
func (g *Gamepad) ButtonCount(id ebiten.GamepadID) int {
	return g.gamepadButtonCountFn(id)
}
func (g *Gamepad) Name(id ebiten.GamepadID) string {
	return g.gamepadNameFn(id)
}
func (g *Gamepad) SDLID(id ebiten.GamepadID) string {
	return g.gamepadSDLIDFn(id)
}
func (g *Gamepad) IsButtonPressed(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	return g.isGamepadButtonPressedFn(id, button)
}
func (g *Gamepad) IsStandardAxisAvailable(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) bool {
	return g.isStandardGamepadAxisAvailableFn(id, axis)
}
func (g *Gamepad) IsStandardButtonAvailable(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return g.isStandardGamepadButtonAvailableFn(id, button)
}
func (g *Gamepad) IsStandardButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return g.isStandardGamepadButtonPressedFn(id, button)
}
func (g *Gamepad) IsStandardLayoutAvailable(id ebiten.GamepadID) bool {
	return g.isStandardGamepadLayoutAvailableFn(id)
}
func (g *Gamepad) StandardAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64 {
	return g.standardGamepadAxisValueFn(id, axis)
}
func (g *Gamepad) AppendJustConnectedIDs(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID {
	return g.appendJustConnectedGamepadIDsFn(gamepadIDs)
}
func (g *Gamepad) AppendJustPressedButtons(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton {
	return g.appendJustPressedGamepadButtonsFn(id, buttons)
}
func (g *Gamepad) AppendJustPressedStandardButtons(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton {
	return g.appendJustPressedStandardGamepadButtonsFn(id, buttons)
}
func (g *Gamepad) AppendJustReleasedButtons(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton {
	return g.appendJustReleasedGamepadButtonsFn(id, buttons)
}
func (g *Gamepad) AppendJustReleasedStandardButtons(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton {
	return g.appendJustReleasedStandardGamepadButtonsFn(id, buttons)
}
func (g *Gamepad) AppendPressedButtons(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton {
	return g.appendPressedGamepadButtonsFn(id, buttons)
}
func (g *Gamepad) AppendPressedStandardButtons(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton {
	return g.appendPressedStandardGamepadButtonsFn(id, buttons)
}
func (g *Gamepad) ButtonPressDuration(id ebiten.GamepadID, button ebiten.GamepadButton) int {
	return g.gamepadButtonPressDurationFn(id, button)
}
func (g *Gamepad) IsButtonJustPressed(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	return g.isGamepadButtonJustPressedFn(id, button)
}
func (g *Gamepad) IsButtonJustReleased(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	return g.isGamepadButtonJustReleasedFn(id, button)
}
func (g *Gamepad) IsJustDisconnected(id ebiten.GamepadID) bool {
	return g.isGamepadJustDisconnectedFn(id)
}
func (g *Gamepad) IsStandardButtonJustPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return g.isStandardGamepadButtonJustPressedFn(id, button)
}
func (g *Gamepad) IsStandardButtonJustReleased(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return g.isStandardGamepadButtonJustReleasedFn(id, button)
}
func (g *Gamepad) StandardButtonPressDuration(id ebiten.GamepadID, button ebiten.StandardGamepadButton) int {
	return g.standardGamepadButtonPressDurationFn(id, button)
}

type GamepadSetter struct {
	gamepad *Gamepad
}

func NewGamepadSetter(g *Gamepad) *GamepadSetter {
	return &GamepadSetter{}
}

func (s *GamepadSetter) SetDefault() {
	s.SetAppendIDsFunc(ebiten.AppendGamepadIDs)
	s.SetAxisCountFunc(ebiten.GamepadAxisCount)
	s.SetAxisValueFunc(ebiten.GamepadAxisValue)
	s.SetButtonCountFunc(ebiten.GamepadButtonCount)
	s.SetNameFunc(ebiten.GamepadName)
	s.SetSDLIDFunc(ebiten.GamepadSDLID)
	s.SetIsButtonPressedFunc(ebiten.IsGamepadButtonPressed)
	s.SetIsStandardAxisAvailableFunc(ebiten.IsStandardGamepadAxisAvailable)
	s.SetIsStandardButtonAvailableFunc(ebiten.IsStandardGamepadButtonAvailable)
	s.SetIsStandardButtonPressedFunc(ebiten.IsStandardGamepadButtonPressed)
	s.SetIsStandardLayoutAvailableFunc(ebiten.IsStandardGamepadLayoutAvailable)
	s.SetStandardAxisValueFunc(ebiten.StandardGamepadAxisValue)
	s.SetAppendJustConnectedIDsFunc(inpututil.AppendJustConnectedGamepadIDs)
	s.SetAppendJustPressedButtonsFunc(inpututil.AppendJustPressedGamepadButtons)
	s.SetAppendJustPressedStandardButtonsFunc(inpututil.AppendJustPressedStandardGamepadButtons)
	s.SetAppendJustReleasedButtonsFunc(inpututil.AppendJustReleasedGamepadButtons)
	s.SetAppendJustReleasedStandardButtonsFunc(inpututil.AppendJustReleasedStandardGamepadButtons)
	s.SetAppendPressedButtonsFunc(inpututil.AppendPressedGamepadButtons)
	s.SetAppendPressedStandardButtonsFunc(inpututil.AppendPressedStandardGamepadButtons)
	s.SetButtonPressDurationFunc(inpututil.GamepadButtonPressDuration)
	s.SetIsButtonJustPressedFunc(inpututil.IsGamepadButtonJustPressed)
	s.SetIsButtonJustReleasedFunc(inpututil.IsGamepadButtonJustReleased)
	s.SetIsJustDisconnectedFunc(inpututil.IsGamepadJustDisconnected)
	s.SetIsStandardButtonJustPressedFunc(inpututil.IsStandardGamepadButtonJustPressed)
	s.SetIsStandardButtonJustReleasedFunc(inpututil.IsStandardGamepadButtonJustReleased)
	s.SetStandardButtonPressDurationFunc(inpututil.StandardGamepadButtonPressDuration)

}

func (s *GamepadSetter) SetAppendIDsFunc(fn func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID) {
	s.gamepad.appendGamepadIDsFn = fn
}
func (s *GamepadSetter) SetAxisCountFunc(fn func(id ebiten.GamepadID) int) {
	s.gamepad.gamepadAxisCountFn = fn
}
func (s *GamepadSetter) SetAxisValueFunc(fn func(id ebiten.GamepadID, axis int) float64) {
	s.gamepad.gamepadAxisValueFn = fn
}
func (s *GamepadSetter) SetButtonCountFunc(fn func(id ebiten.GamepadID) int) {
	s.gamepad.gamepadButtonCountFn = fn
}
func (s *GamepadSetter) SetNameFunc(fn func(id ebiten.GamepadID) string) {
	s.gamepad.gamepadNameFn = fn
}
func (s *GamepadSetter) SetSDLIDFunc(fn func(id ebiten.GamepadID) string) {
	s.gamepad.gamepadSDLIDFn = fn
}
func (s *GamepadSetter) SetIsButtonPressedFunc(fn func(id ebiten.GamepadID, button ebiten.GamepadButton) bool) {
	s.gamepad.isGamepadButtonPressedFn = fn
}
func (s *GamepadSetter) SetIsStandardAxisAvailableFunc(fn func(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) bool) {
	s.gamepad.isStandardGamepadAxisAvailableFn = fn
}
func (s *GamepadSetter) SetIsStandardButtonAvailableFunc(fn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool) {
	s.gamepad.isStandardGamepadButtonAvailableFn = fn
}
func (s *GamepadSetter) SetIsStandardButtonPressedFunc(fn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool) {
	s.gamepad.isStandardGamepadButtonPressedFn = fn
}
func (s *GamepadSetter) SetIsStandardLayoutAvailableFunc(fn func(id ebiten.GamepadID) bool) {
	s.gamepad.isStandardGamepadLayoutAvailableFn = fn
}
func (s *GamepadSetter) SetStandardAxisValueFunc(fn func(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64) {
	s.gamepad.standardGamepadAxisValueFn = fn
}
func (s *GamepadSetter) SetAppendJustConnectedIDsFunc(fn func(gamepadIDs []ebiten.GamepadID) []ebiten.GamepadID) {
	s.gamepad.appendJustConnectedGamepadIDsFn = fn
}
func (s *GamepadSetter) SetAppendJustPressedButtonsFunc(fn func(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton) {
	s.gamepad.appendJustPressedGamepadButtonsFn = fn
}
func (s *GamepadSetter) SetAppendJustPressedStandardButtonsFunc(fn func(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton) {
	s.gamepad.appendJustPressedStandardGamepadButtonsFn = fn
}
func (s *GamepadSetter) SetAppendJustReleasedButtonsFunc(fn func(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton) {
	s.gamepad.appendJustReleasedGamepadButtonsFn = fn
}
func (s *GamepadSetter) SetAppendJustReleasedStandardButtonsFunc(fn func(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton) {
	s.gamepad.appendJustReleasedStandardGamepadButtonsFn = fn
}
func (s *GamepadSetter) SetAppendPressedButtonsFunc(fn func(id ebiten.GamepadID, buttons []ebiten.GamepadButton) []ebiten.GamepadButton) {
	s.gamepad.appendPressedGamepadButtonsFn = fn
}
func (s *GamepadSetter) SetAppendPressedStandardButtonsFunc(fn func(id ebiten.GamepadID, buttons []ebiten.StandardGamepadButton) []ebiten.StandardGamepadButton) {
	s.gamepad.appendPressedStandardGamepadButtonsFn = fn
}
func (s *GamepadSetter) SetButtonPressDurationFunc(fn func(id ebiten.GamepadID, button ebiten.GamepadButton) int) {
	s.gamepad.gamepadButtonPressDurationFn = fn
}
func (s *GamepadSetter) SetIsButtonJustPressedFunc(fn func(id ebiten.GamepadID, button ebiten.GamepadButton) bool) {
	s.gamepad.isGamepadButtonJustPressedFn = fn
}
func (s *GamepadSetter) SetIsButtonJustReleasedFunc(fn func(id ebiten.GamepadID, button ebiten.GamepadButton) bool) {
	s.gamepad.isGamepadButtonJustReleasedFn = fn
}
func (s *GamepadSetter) SetIsJustDisconnectedFunc(fn func(id ebiten.GamepadID) bool) {
	s.gamepad.isGamepadJustDisconnectedFn = fn
}
func (s *GamepadSetter) SetIsStandardButtonJustPressedFunc(fn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool) {
	s.gamepad.isStandardGamepadButtonJustPressedFn = fn
}
func (s *GamepadSetter) SetIsStandardButtonJustReleasedFunc(fn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool) {
	s.gamepad.isStandardGamepadButtonJustReleasedFn = fn
}
func (s *GamepadSetter) SetStandardButtonPressDurationFunc(fn func(id ebiten.GamepadID, button ebiten.StandardGamepadButton) int) {
	s.gamepad.standardGamepadButtonPressDurationFn = fn
}
