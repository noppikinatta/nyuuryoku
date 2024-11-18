package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse struct {
	cursorPositionFn func() (int, int)
	pressedFn        func(button ebiten.MouseButton) bool
	justPressedFn    func(button ebiten.MouseButton) bool
	justReleasedFn   func(button ebiten.MouseButton) bool
	pressDurationFn  func(button ebiten.MouseButton) int
	wheelFn          func() (xoff, yoff float64)
}

func NewMouse() *Mouse {
	m := &Mouse{}
	s := MouseSetter{mouse: m}
	s.SetDefault()

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

func (m *Mouse) PressDuration(button ebiten.MouseButton) int {
	return m.pressDurationFn(button)
}

func (m *Mouse) Wheel() (xoff, yoff float64) {
	return m.wheelFn()
}

type MouseSetter struct {
	mouse *Mouse
}

func NewMouseSetter(mouse *Mouse) *MouseSetter {
	return &MouseSetter{mouse: mouse}
}

func (s *MouseSetter) SetDefault() {
	s.SetCursorPositionFunc(ebiten.CursorPosition)
	s.SetPressedFunc(ebiten.IsMouseButtonPressed)
	s.SetJustPressedFunc(inpututil.IsMouseButtonJustPressed)
	s.SetJustReleasedFunc(inpututil.IsMouseButtonJustReleased)
	s.SetPressDurationFunc(inpututil.MouseButtonPressDuration)
	s.SetWheelFunc(ebiten.Wheel)
}

func (s *MouseSetter) SetCursorPositionFunc(cursorPositionFn func() (int, int)) {
	s.mouse.cursorPositionFn = cursorPositionFn
}

func (s *MouseSetter) SetPressedFunc(pressedFn func(button ebiten.MouseButton) bool) {
	s.mouse.pressedFn = pressedFn
}

func (s *MouseSetter) SetJustPressedFunc(justPressedFn func(button ebiten.MouseButton) bool) {
	s.mouse.justPressedFn = justPressedFn
}

func (s *MouseSetter) SetJustReleasedFunc(justReleasedFn func(button ebiten.MouseButton) bool) {
	s.mouse.justReleasedFn = justReleasedFn
}

func (s *MouseSetter) SetPressDurationFunc(pressDurationFn func(button ebiten.MouseButton) int) {
	s.mouse.pressDurationFn = pressDurationFn
}

func (s *MouseSetter) SetWheelFunc(wheelFn func() (xoff, yoff float64)) {
	s.mouse.wheelFn = wheelFn
}
