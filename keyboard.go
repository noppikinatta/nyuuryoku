package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	pressedFn          func(key ebiten.Key) bool
	justPressedFn      func(key ebiten.Key) bool
	justReleasedFn     func(key ebiten.Key) bool
	pressedDurationFn  func(key ebiten.Key) int
	keyNameFn          func(key ebiten.Key) string
	appendPressedFn    func(keys []ebiten.Key) []ebiten.Key
	appendInputCharsFn func(runes []rune) []rune
}

func NewKeyboard2() *Keyboard {
	k := &Keyboard{}
	s := KeyboardSetter{Keyboard: k}

	s.SetPressed(ebiten.IsKeyPressed)
	s.SetJustPressed(inpututil.IsKeyJustPressed)
	s.SetJustReleased(inpututil.IsKeyJustReleased)
	s.SetPressedDuration(inpututil.KeyPressDuration)
	s.SetKeyName(ebiten.KeyName)
	s.SetAppendPressed(inpututil.AppendPressedKeys)
	s.SetAppendInputChars(ebiten.AppendInputChars)

	return k
}

func (k *Keyboard) Pressed(key ebiten.Key) bool {
	return k.pressedFn(key)
}

func (k *Keyboard) JustPressed(key ebiten.Key) bool {
	return k.justPressedFn(key)
}

func (k *Keyboard) JustReleased(key ebiten.Key) bool {
	return k.justReleasedFn(key)
}

func (k *Keyboard) PressedDuration(key ebiten.Key) int {
	return k.pressedDurationFn(key)
}

func (k *Keyboard) KeyName(key ebiten.Key) string {
	return k.keyNameFn(key)
}

func (k *Keyboard) AppendPressed(keys []ebiten.Key) []ebiten.Key {
	return k.appendPressedFn(keys)
}

func (k *Keyboard) AppendInputChars(runes []rune) []rune {
	return k.appendInputCharsFn(runes)
}

type KeyboardSetter struct {
	Keyboard *Keyboard
}

func (s *KeyboardSetter) SetPressed(pressedFn func(key ebiten.Key) bool) {
	s.Keyboard.pressedFn = pressedFn
}

func (s *KeyboardSetter) SetJustPressed(justPressedFn func(key ebiten.Key) bool) {
	s.Keyboard.justPressedFn = justPressedFn
}

func (s *KeyboardSetter) SetJustReleased(justReleasedFn func(key ebiten.Key) bool) {
	s.Keyboard.justReleasedFn = justReleasedFn
}

func (s *KeyboardSetter) SetPressedDuration(pressedDurationFn func(key ebiten.Key) int) {
	s.Keyboard.pressedDurationFn = pressedDurationFn
}

func (s *KeyboardSetter) SetKeyName(keyNameFn func(key ebiten.Key) string) {
	s.Keyboard.keyNameFn = keyNameFn
}

func (s *KeyboardSetter) SetAppendPressed(appendPressedFn func(keys []ebiten.Key) []ebiten.Key) {
	s.Keyboard.appendPressedFn = appendPressedFn
}

func (s *KeyboardSetter) SetAppendInputChars(appendInputCharsFn func(runes []rune) []rune) {
	s.Keyboard.appendInputCharsFn = appendInputCharsFn
}
