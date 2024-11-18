package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct {
	pressedFn            func(key ebiten.Key) bool
	justPressedFn        func(key ebiten.Key) bool
	justReleasedFn       func(key ebiten.Key) bool
	pressDurationFn      func(key ebiten.Key) int
	keyNameFn            func(key ebiten.Key) string
	appendPressedFn      func(keys []ebiten.Key) []ebiten.Key
	appendJustPressedFn  func(keys []ebiten.Key) []ebiten.Key
	appendJustReleasedFn func(keys []ebiten.Key) []ebiten.Key
	appendInputCharsFn   func(runes []rune) []rune
}

func NewKeyboard() *Keyboard {
	k := &Keyboard{}
	s := KeyboardSetter{keyboard: k}
	s.SetDefault()

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

func (k *Keyboard) PressDuration(key ebiten.Key) int {
	return k.pressDurationFn(key)
}

func (k *Keyboard) KeyName(key ebiten.Key) string {
	return k.keyNameFn(key)
}

func (k *Keyboard) AppendPressed(keys []ebiten.Key) []ebiten.Key {
	return k.appendPressedFn(keys)
}

func (k *Keyboard) AppendJustPressed(keys []ebiten.Key) []ebiten.Key {
	return k.appendJustPressedFn(keys)
}

func (k *Keyboard) AppendJustReleased(keys []ebiten.Key) []ebiten.Key {
	return k.appendJustReleasedFn(keys)
}

func (k *Keyboard) AppendInputChars(runes []rune) []rune {
	return k.appendInputCharsFn(runes)
}

type KeyboardSetter struct {
	keyboard *Keyboard
}

func NewKeyboardSetter(keyboard *Keyboard) *KeyboardSetter {
	return &KeyboardSetter{keyboard: keyboard}
}

func (s *KeyboardSetter) SetDefault() {
	s.SetPressedFunc(ebiten.IsKeyPressed)
	s.SetJustPressedFunc(inpututil.IsKeyJustPressed)
	s.SetJustReleasedFunc(inpututil.IsKeyJustReleased)
	s.SetPressDurationFunc(inpututil.KeyPressDuration)
	s.SetKeyNameFunc(ebiten.KeyName)
	s.SetAppendPressedFunc(inpututil.AppendPressedKeys)
	s.SetAppendJustPressedFunc(inpututil.AppendJustPressedKeys)
	s.SetAppendJustReleasedFunc(inpututil.AppendJustReleasedKeys)
	s.SetAppendInputCharsFunc(ebiten.AppendInputChars)
}

func (s *KeyboardSetter) SetPressedFunc(pressedFn func(key ebiten.Key) bool) {
	s.keyboard.pressedFn = pressedFn
}

func (s *KeyboardSetter) SetJustPressedFunc(justPressedFn func(key ebiten.Key) bool) {
	s.keyboard.justPressedFn = justPressedFn
}

func (s *KeyboardSetter) SetJustReleasedFunc(justReleasedFn func(key ebiten.Key) bool) {
	s.keyboard.justReleasedFn = justReleasedFn
}

func (s *KeyboardSetter) SetPressDurationFunc(pressDurationFn func(key ebiten.Key) int) {
	s.keyboard.pressDurationFn = pressDurationFn
}

func (s *KeyboardSetter) SetKeyNameFunc(keyNameFn func(key ebiten.Key) string) {
	s.keyboard.keyNameFn = keyNameFn
}

func (s *KeyboardSetter) SetAppendPressedFunc(appendPressedFn func(keys []ebiten.Key) []ebiten.Key) {
	s.keyboard.appendPressedFn = appendPressedFn
}

func (s *KeyboardSetter) SetAppendJustPressedFunc(appendJustPressedFn func(keys []ebiten.Key) []ebiten.Key) {
	s.keyboard.appendJustPressedFn = appendJustPressedFn
}

func (s *KeyboardSetter) SetAppendJustReleasedFunc(appendJustReleasedFn func(keys []ebiten.Key) []ebiten.Key) {
	s.keyboard.appendJustReleasedFn = appendJustReleasedFn
}

func (s *KeyboardSetter) SetAppendInputCharsFunc(appendInputCharsFn func(runes []rune) []rune) {
	s.keyboard.appendInputCharsFn = appendInputCharsFn
}
