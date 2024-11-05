package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard interface {
	Pressed(key ebiten.Key) bool
	JustPressed(key ebiten.Key) bool
	JustReleased(key ebiten.Key) bool
	PressedDuration(key ebiten.Key) int
	Idle(key ebiten.Key) bool
	AppendInputChars(cc []rune) []rune
}

func NewKeyboard() Keyboard {
	return &keyboardImpl{}
}

type keyboardImpl struct {
}

func (k *keyboardImpl) Pressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (k *keyboardImpl) JustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

func (k *keyboardImpl) JustReleased(key ebiten.Key) bool {
	return inpututil.IsKeyJustReleased(key)
}

func (k *keyboardImpl) PressedDuration(key ebiten.Key) int {
	return inpututil.KeyPressDuration(key)
}

func (k *keyboardImpl) Idle(key ebiten.Key) bool {
	return !ebiten.IsKeyPressed(key) && !inpututil.IsKeyJustReleased(key)
}

func (k *keyboardImpl) AppendInputChars(runes []rune) []rune {
	return ebiten.AppendInputChars(runes)
}
