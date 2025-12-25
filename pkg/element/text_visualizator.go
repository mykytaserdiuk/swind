package element

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/bus"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type TextVisualizator struct {
	rect        rl.Rectangle
	text        string
	currentText string

	bus bus.Bus
}

func NewTextVisualizator(rect rl.Rectangle, b bus.Bus, txt string) *TextVisualizator {
	tv := &TextVisualizator{
		rect: rect,
		bus:  b,
		text: strings.ToUpper(txt),
	}

	tv.bus.Subscribe(bus.StateUpdate, func(e bus.Event) {
		if data, ok := e.Data.(models.UIEvent); ok {
			if data.Type == "input_text" {
				if len(data.ID) > len(tv.text) {
					return
				}
				tv.currentText = strings.ToUpper(data.ID)
			}
		}
	})

	return tv
}

func (v *TextVisualizator) Update(dt float32) {

}

func (v *TextVisualizator) Draw(r render.Renderer) {
	x := v.rect.X
	font := rl.GetFontDefault()
	r.Submit(render.DrawCmd{
		Layer: models.LayerContent,
		Fn: func() {
			for i, seg := range v.currentText {
				color := rl.Red
				if string(seg) == string(v.text[i]) {
					color = rl.DarkGreen
				}
				rl.DrawTextEx(font, string(seg), rl.NewVector2(x, v.rect.Y), 20, 30, color)
				x += rl.MeasureTextEx(font, string(seg), 20, 30).X + 2
			}
		},
	})
}

func (v *TextVisualizator) HandleInput() bool {
	return false
}
