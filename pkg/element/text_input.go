package element

import (
	"strings"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mykytaserdiuk/fluxo"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type TextInput struct {
	rect  rl.Rectangle
	input string

	prevText string
	maxText  int

	b fluxo.Bus
}

func NewTextInput(rect rl.Rectangle, max int, b fluxo.Bus) *TextInput {
	return &TextInput{rect: rect, b: b, maxText: max}
}

func (i *TextInput) HandleInput() bool {
	return false
}

func (i *TextInput) Update(dt float32) {
	i.input = strings.ToUpper(i.input)
	if len(i.prevText) > len(i.input) {
		i.input = i.prevText
		return
	}

	if i.prevText != i.input {
		i.b.Emit(models.StateUpdate, models.Event{
			Type: models.StateUpdate,
			Data: models.UIEvent{
				Type: "input_text",
				ID:   i.input,
			},
		})

		i.prevText = i.input
	}
}

func (i *TextInput) Draw(r render.Renderer) {
	r.Submit(render.DrawCmd{
		Layer: models.LayerContent,
		Fn: func() {
			rg.TextBox(i.rect, &i.input, i.maxText, true)
		},
	})
}
