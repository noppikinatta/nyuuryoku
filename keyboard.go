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
	KeyName(key ebiten.Key) string
	AppendPressed(keys []ebiten.Key) []ebiten.Key
	AppendInputChars(runes []rune) []rune
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

func (k *keyboardImpl) KeyName(key ebiten.Key) string {
	return ebiten.KeyName(key)
}

func (k *keyboardImpl) AppendPressed(keys []ebiten.Key) []ebiten.Key {
	return inpututil.AppendPressedKeys(keys)
}

func (k *keyboardImpl) AppendInputChars(runes []rune) []rune {
	return ebiten.AppendInputChars(runes)
}
