package element

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mykytaserdiuk/fluxo"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type TextVisualizator struct {
	rect        rl.Rectangle
	text        string
	currentText string
	fontSize    float32
	bus         fluxo.Bus
}

func NewTextVisualizator(rect rl.Rectangle, b fluxo.Bus, txt string) *TextVisualizator {
	tv := &TextVisualizator{
		rect:     rect,
		bus:      b,
		fontSize: 45,
		text:     strings.ToUpper(txt),
	}

	tv.bus.Subscribe(models.StateUpdate, func(e models.Event) {
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
				rl.DrawTextEx(font, string(seg), rl.NewVector2(x, v.rect.Y), v.fontSize, 30, color)
				x += rl.MeasureTextEx(font, string(seg), v.fontSize, 30).X + 2.5
			}
		},
	})
}

func (v *TextVisualizator) HandleInput() bool {
	return false
}
