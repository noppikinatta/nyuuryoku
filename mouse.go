package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mouse interface {
	CursorPosition() (int, int)
	Clicked(button ebiten.MouseButton) bool
	Pressed(button ebiten.MouseButton) bool
	JustPressed(button ebiten.MouseButton) bool
	JustReleased(button ebiten.MouseButton) bool
	PressedDuration(button ebiten.MouseButton) int
	Wheel() (xoff, yoff float64)
	Idle(button ebiten.MouseButton) bool
}

type mouseImpl struct {
}

func NewMouse() Mouse {
	return &mouseImpl{}
}

func (m *mouseImpl) CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func (m *mouseImpl) Clicked(button ebiten.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(button)
}

func (m *mouseImpl) Pressed(button ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(button)
}

func (m *mouseImpl) JustPressed(button ebiten.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(button)
}

func (m *mouseImpl) JustReleased(button ebiten.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(button)
}

func (m *mouseImpl) PressedDuration(button ebiten.MouseButton) int {
	return inpututil.MouseButtonPressDuration(button)
}

func (m *mouseImpl) Wheel() (xoff, yoff float64) {
	return ebiten.Wheel()
}

func (m *mouseImpl) Idle(button ebiten.MouseButton) bool {
	return !ebiten.IsMouseButtonPressed(button) && !inpututil.IsMouseButtonJustReleased(button)
}
