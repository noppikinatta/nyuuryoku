package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	cursorPositionFn  func() (int, int)
	pressedFn         func(button ebiten.MouseButton) bool
	justPressedFn     func(button ebiten.MouseButton) bool
	justReleasedFn    func(button ebiten.MouseButton) bool
	pressedDurationFn func(button ebiten.MouseButton) int
	wheelFn           func() (xoff, yoff float64)
}

func NewMouse() *Mouse {
	m := &Mouse{}
	s := MouseSetter{Mouse: m}

	s.SetCursorPositionFunc(ebiten.CursorPosition)
	s.SetPressedFunc(ebiten.IsMouseButtonPressed)
	s.SetJustPressedFunc(inpututil.IsMouseButtonJustPressed)
	s.SetJustReleasedFunc(inpututil.IsMouseButtonJustReleased)
	s.SetPressedDurationFunc(inpututil.MouseButtonPressDuration)
	s.SetWheelFunc(ebiten.Wheel)

	return m
}

func (m *Mouse) CursorPosition() (int, int) {
	return m.cursorPositionFn()
}

func (m *Mouse) Pressed(button ebiten.MouseButton) bool {
	return m.pressedFn(button)
}

func (m *Mouse) JustPressed(button ebiten.MouseButton) bool {
	return m.justPressedFn(button)
}

func (m *Mouse) JustReleased(button ebiten.MouseButton) bool {
	return m.justReleasedFn(button)
}

func (m *Mouse) PressedDuration(button ebiten.MouseButton) int {
	return m.pressedDurationFn(button)
}

func (m *Mouse) Wheel() (xoff, yoff float64) {
	return m.wheelFn()
}

type MouseSetter struct {
	Mouse *Mouse
}

func (m *MouseSetter) SetCursorPositionFunc(cursorPositionFn func() (int, int)) {
	m.Mouse.cursorPositionFn = cursorPositionFn
}

func (m *MouseSetter) SetPressedFunc(pressedFn func(button ebiten.MouseButton) bool) {
	m.Mouse.pressedFn = pressedFn
}

func (m *MouseSetter) SetJustPressedFunc(justPressedFn func(button ebiten.MouseButton) bool) {
	m.Mouse.justPressedFn = justPressedFn
}

func (m *MouseSetter) SetJustReleasedFunc(justReleasedFn func(button ebiten.MouseButton) bool) {
	m.Mouse.justReleasedFn = justReleasedFn
}

func (m *MouseSetter) SetPressedDurationFunc(pressedDurationFn func(button ebiten.MouseButton) int) {
	m.Mouse.pressedDurationFn = pressedDurationFn
}

func (m *MouseSetter) SetWheelFunc(wheelFn func() (xoff, yoff float64)) {
	m.Mouse.wheelFn = wheelFn
}
