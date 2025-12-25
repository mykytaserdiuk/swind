package element

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/bus"
	"github.com/nikitaserdiuk9/swind/pkg/input"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type TextInput struct {
	rect  rl.Rectangle
	input string

	prevText string

	dragging   bool
	dragOffset rl.Vector2
	b          bus.Bus
}

func NewTextInput(rect rl.Rectangle, b bus.Bus) *TextInput {
	return &TextInput{rect: rect, b: b}
}

func (i *TextInput) HandleInput() bool {
	mx, my := input.MousePos.X, input.MousePos.Y

	over := rl.CheckCollisionPointRec(input.MousePos, i.rect)

	if input.MouseLeftPressed && over {
		i.dragging = true
		i.dragOffset = rl.NewVector2(mx-i.rect.X, my-i.rect.Y)
	}

	if !input.MouseLeftDown && i.dragging {
		i.dragging = false
		return true
	}

	return i.dragging || over
}

func (i *TextInput) Update(dt float32) {
	if i.prevText != i.input {
		i.b.Emit(bus.Event{
			Type: bus.StateUpdate,
			Data: models.UIEvent{
				Type: "input_text",
				ID:   i.input,
			},
		})
		i.prevText = i.input
		fmt.Println(i.input)
	}

	if i.dragging {
		if i.dragging {
			mx, my := input.MousePos.X, input.MousePos.Y
			i.rect.X = mx - i.dragOffset.X
			i.rect.Y = my - i.dragOffset.Y
		}
	}
}

func (i *TextInput) Draw(r render.Renderer) {
	r.Submit(render.DrawCmd{
		Layer: models.LayerContent,
		Fn: func() {
			// rg.TextInputBox(i.rect, "title", "meessage", "OK", &i.input, 100, &i.isSelected)
			rg.TextBox(i.rect, &i.input, 500, true)
		},
	})
}
