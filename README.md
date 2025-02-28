# Nyuuryoku

A flexible input abstraction layer for [Ebitengine](https://ebitengine.org/) game development.

## Overview

Nyuuryoku is an input abstraction library for the Ebiten game engine. It provides a wrapper around Ebiten's input APIs, allowing for easier testing and more flexible input handling in your games.

The main goal of this library is to enable testing without requiring actual human input, making automated testing of game input systems possible.

## Features

- **Complete abstraction** of Ebiten's input APIs (Gamepad, Mouse, Keyboard)
- **Easy switching** between real and virtual input sources
- **Same interface** as Ebiten's built-in input functions
- **Test-friendly** design that allows for automated testing of input-dependent code

## Installation

```
go get github.com/noppikinatta/nyuuryoku
```

## Basic Usage

```go
import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/noppikinatta/nyuuryoku"
)

// Create input handlers
keyboard := nyuuryoku.NewKeyboard()
mouse := nyuuryoku.NewMouse()
gamepad := nyuuryoku.NewGamepad()

// Use them similar to Ebiten's native input
func (g *Game) Update() error {
    // Keyboard example
    if keyboard.IsKeyPressed(ebiten.KeySpace) {
        // Handle space key pressed
    }
    
    // Mouse example
    x, y := mouse.Position()
    if mouse.IsButtonPressed(ebiten.MouseButtonLeft) {
        // Handle mouse click at (x,y)
    }
    
    // Gamepad example
    var ids []ebiten.GamepadID
    ids = gamepad.AppendIDs(ids)
    for _, id := range ids {
        if gamepad.IsButtonPressed(id, 0) {
            // Handle gamepad button press
        }
    }
    
    return nil
}
```

## Working with Virtual Input

Nyuuryoku allows you to easily switch between real input devices and virtual ones for testing:

```go
// Create and configure input handlers
keyboard := nyuuryoku.NewKeyboard()
gamepad := nyuuryoku.NewGamepad()

// Switch keyboard to virtual input
setter := nyuuryoku.NewKeyboardSetter(keyboard)
setter.SetIsKeyPressedFunc(myVirtualKeyboard.isKeyPressed)
setter.SetAppendPressedFunc(myVirtualKeyboard.appendPressed)
// ... set other virtual functions

// Switch gamepad to virtual input
gameSetter := nyuuryoku.NewGamepadSetter(gamepad)
gameSetter.SetAppendIDsFunc(myVirtualGamepad.appendIDs)
gameSetter.SetIsButtonPressedFunc(myVirtualGamepad.isButtonPressed)
// ... set other virtual functions

// Reset to real input when needed
setter.SetDefault()
gameSetter.SetDefault()
```

## Examples

The repository includes examples for each input type:

- `examples/keyboard/` - Keyboard input examples with virtual keyboard support
- `examples/mouse/` - Mouse input examples
- `examples/gamepad/` - Gamepad input examples with virtual gamepad support

Run the examples to see how the library works:

```
cd examples/gamepad
go run .
```

## Testing Games with Virtual Input

Nyuuryoku is particularly useful for testing game input systems without requiring human interaction:

```go
func TestPlayerMovement(t *testing.T) {
    game := NewGame()
    keyboard := nyuuryoku.NewKeyboard()
    game.SetKeyboard(keyboard)
    
    // Create a virtual keyboard
    vk := NewTestKeyboard()
    setter := nyuuryoku.NewKeyboardSetter(keyboard)
    setter.SetIsKeyPressedFunc(func(key ebiten.Key) bool {
        return key == ebiten.KeyRight // Simulate pressing right key
    })
    
    // Run game update
    game.Update()
    
    // Assert player moved to the right
    if game.player.X <= initialX {
        t.Errorf("Player didn't move right")
    }
}
```

## API Documentation

The library provides three main input handlers:

- `Keyboard` - Handles keyboard input
- `Mouse` - Handles mouse input and cursor position
- `Gamepad` - Handles gamepad/controller input

Each comes with a corresponding `*Setter` type that allows switching between real and virtual input sources.

## License

MIT License - See LICENSE file for details
