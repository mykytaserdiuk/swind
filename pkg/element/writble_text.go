package element

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/bus"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type WritableText struct {
	rect        rl.Rectangle
	text        string
	currentText string
	fontSize    float32

	bus bus.Bus
}

func NewWritableText(rect rl.Rectangle, b bus.Bus, text string) *WritableText {
	wt := &WritableText{
		bus:      b,
		rect:     rect,
		text:     text,
		fontSize: 20,
	}
	wt.bus.Subscribe(bus.StateUpdate, func(e bus.Event) {
		if data, ok := e.Data.(models.UIEvent); ok {
			if data.Type == "input_text" {
				wt.currentText = data.ID
			}
		}
	})
	return wt
}

func (t *WritableText) Update(dt float32) {

}

func (t *WritableText) Draw(r render.Renderer) {
	x := t.rect.X
	font := rl.GetFontDefault()
	r.Submit(render.DrawCmd{
		Layer: models.LayerContent,
		Fn: func() {
			for i, seg := range t.text {
				color := rl.Black
				if i < len(t.currentText) {
					color = rl.Lime
				}
				rl.DrawTextEx(font, string(seg), rl.NewVector2(x, t.rect.Y), t.fontSize, 30, color)
				x += rl.MeasureTextEx(font, string(seg), t.fontSize, 30).X + 2
			}
		},
	})
}

func (t *WritableText) HandleInput() bool {
	return false
}
