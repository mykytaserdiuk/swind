package input

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	MousePos         rl.Vector2
	MouseDelta       rl.Vector2
	MouseLeftPressed bool
	MouseLeftDown    bool
)

type InputSystem struct{}

func NewInputSystem() *InputSystem {
	return &InputSystem{}
}

func (i *InputSystem) Update() {
	MousePos = rl.GetMousePosition()
	MouseDelta = rl.GetMouseDelta()
	MouseLeftPressed = rl.IsMouseButtonDown(rl.MouseButtonLeft)
	MouseLeftDown = rl.IsMouseButtonDown(rl.MouseButtonRight)

}
