package element

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/bus"
	"github.com/nikitaserdiuk9/swind/pkg/input"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type Button struct {
	rect    rl.Rectangle
	bus     bus.Bus
	focused bool
}

func NewButton(rect rl.Rectangle, bus bus.Bus) *Button {
	return &Button{rect: rect, bus: bus}
}

func (b *Button) HandleInput() bool {
	over := rl.CheckCollisionPointRec(input.MousePos, b.rect)
	if b.focused && !over {
		b.focused = false
		return true
	}

	return over
}

func (b *Button) Update(dt float32) {

}

func (b *Button) Draw(r render.Renderer) {
	r.Submit(render.DrawCmd{
		Layer: models.LayerContent,
		Fn: func() {
			if rg.Button(b.rect, "Exit") {
				b.bus.Emit(bus.Event{Type: bus.UIEvent, Data: models.UIEvent{ID: "Exit", Type: "button_click"}})
			}
		},
	})
}
