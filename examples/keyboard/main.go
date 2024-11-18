package main

import (
	"fmt"
	"image/color"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/noppikinatta/nyuuryoku"
)

const (
	screenW = 800
	screenH = 600
)

func main() {
	ebiten.SetWindowSize(screenW, screenH)

	g := game{}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

type game struct {
	keyboard          *nyuuryoku.Keyboard
	tmpKeys           []ebiten.Key
	log               []string
	logIdx            int
	pressedInfo       []string
	keyboardIsVirtual bool
	virtualKeyboard   *virtualKeyboard
}

func (g *game) Update() error {
	if g.keyboard == nil {
		g.keyboard = nyuuryoku.NewKeyboard()
	}
	if g.virtualKeyboard == nil {
		g.virtualKeyboard = newVirtualKeyboard()
	}

	if g.keyboardIsVirtual {
		g.virtualKeyboard.Update()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.switchKeyboard()
	}

	g.appendKeysLog()
	g.updatePressedInfo()

	return nil
}

func (g *game) appendKeysLog() {
	g.tmpKeys = g.tmpKeys[:0]
	g.tmpKeys = g.keyboard.AppendJustPressed(g.tmpKeys)
	for _, k := range g.tmpKeys {
		g.appendJustLog(k, "Just Pressed ")
	}

	g.tmpKeys = g.tmpKeys[:0]
	g.tmpKeys = g.keyboard.AppendJustReleased(g.tmpKeys)
	for _, k := range g.tmpKeys {
		g.appendJustLog(k, "Just Released")
	}
}

func (g *game) appendJustLog(key ebiten.Key, prefix string) {
	line := fmt.Sprintf("%s:%s", prefix, key.String())
	g.appendLog(line)
}

func (g *game) updatePressedInfo() {
	g.tmpKeys = g.tmpKeys[:0]
	g.tmpKeys = g.keyboard.AppendPressed(g.tmpKeys)
	g.pressedInfo = g.pressedInfo[:0]
	for _, k := range g.tmpKeys {
		line := fmt.Sprintf("%s pressed: %d frames", k.String(), g.keyboard.PressDuration(k))
		g.pressedInfo = append(g.pressedInfo, line)
	}
}

func (g *game) appendLog(line string) {
	const max = 48
	if len(g.log) < max {
		g.log = append(g.log, line)
	} else {
		g.log[g.logIdx] = line
		g.logIdx++
		g.logIdx = g.logIdx % max
	}
}

func (g *game) switchKeyboard() {
	setter := nyuuryoku.KeyboardSetter{Keyboard: g.keyboard}

	if g.keyboardIsVirtual {
		setter.SetDefault()
	} else {
		setter.SetAppendPressedFunc(g.virtualKeyboard.appendPressed)
		setter.SetAppendJustPressedFunc(g.virtualKeyboard.appendJustPressed)
		setter.SetAppendJustReleasedFunc(g.virtualKeyboard.appendJustReleased)
		setter.SetPressDurationFunc(g.virtualKeyboard.pressedDuration)
	}

	g.keyboardIsVirtual = !g.keyboardIsVirtual
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 48})
	if g.keyboard == nil {
		return
	}

	g.drawPressedInfo(screen)
	g.drawLog(screen)

	ebitenutil.DebugPrintAt(screen, "left click to switch to virtual keyboard", 10, 576)
}

func (g *game) drawPressedInfo(screen *ebiten.Image) {
	for i := range g.pressedInfo {
		info := g.pressedInfo[i]
		y := i * 12

		ebitenutil.DebugPrintAt(screen, info, 400, y)
	}
}

func (g *game) drawLog(screen *ebiten.Image) {
	for i := range g.log {
		line := g.log[i]

		yIdx := (i - g.logIdx)
		if yIdx < 0 {
			yIdx += len(g.log)
		}
		y := yIdx * 12

		ebitenutil.DebugPrintAt(screen, line, 200, y)
	}
}

type keyState struct {
	StateChangedCount int
	Pressed           bool
}

type virtualKeyboard struct {
	tmpKeys []ebiten.Key
	states  map[ebiten.Key]keyState
	count   int
}

func newVirtualKeyboard() *virtualKeyboard {
	return &virtualKeyboard{
		states: make(map[ebiten.Key]keyState),
	}
}

func (k *virtualKeyboard) Update() {
	k.count++
	for key := range ebiten.KeyMax {
		state, ok := k.states[key]
		if !ok {
			state = keyState{}
		}

		newPressed := (((k.count + int(key)*30) / (100 + int(key))) % 4) == 1
		if state.Pressed != newPressed {
			state.Pressed = newPressed
			state.StateChangedCount = k.count
			k.states[key] = state
		}
	}
}

func (k *virtualKeyboard) keys() []ebiten.Key {
	k.tmpKeys = k.tmpKeys[:0]
	for key := range k.states {
		k.tmpKeys = append(k.tmpKeys, key)
	}

	slices.Sort(k.tmpKeys)

	return k.tmpKeys
}

func (k *virtualKeyboard) appendPressed(keys []ebiten.Key) []ebiten.Key {
	return k.appendCondition(keys, func(state keyState) bool {
		return state.Pressed
	})
}

func (k *virtualKeyboard) appendJustPressed(keys []ebiten.Key) []ebiten.Key {
	return k.appendCondition(keys, func(state keyState) bool {
		return state.Pressed && state.StateChangedCount == k.count
	})
}

func (k *virtualKeyboard) appendJustReleased(keys []ebiten.Key) []ebiten.Key {
	return k.appendCondition(keys, func(state keyState) bool {
		return !state.Pressed && state.StateChangedCount == k.count
	})
}

func (k *virtualKeyboard) appendCondition(keys []ebiten.Key, cond func(state keyState) bool) []ebiten.Key {
	for _, key := range k.keys() {
		state := k.states[key]
		if cond(state) {
			keys = append(keys, key)
		}
	}

	return keys
}

func (k *virtualKeyboard) pressedDuration(key ebiten.Key) int {
	state, ok := k.states[key]
	if !ok {
		return 0
	}
	if !state.Pressed {
		return 0
	}

	return k.count - state.StateChangedCount + 1
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}
