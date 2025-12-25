package render

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/models"
)

type RaylibRender struct {
	queue []DrawCmd
	cam   *rl.Camera2D
}

func NewRaylibRender(cam *rl.Camera2D) *RaylibRender {
	return &RaylibRender{
		cam:   cam,
		queue: []DrawCmd{},
	}
}

func (r *RaylibRender) Text(layer models.Layer, cmd TextRenderCmd) {
	r.queue = append(r.queue, DrawCmd{
		Layer: layer,
		Fn: func() {
			rl.DrawText(cmd.Text, cmd.PosX, cmd.PosY, cmd.FontSize, cmd.Col)
		},
	})
}

func (r *RaylibRender) Rect(layer models.Layer, cmd RectRenderCmd) {
	r.queue = append(r.queue, DrawCmd{
		Layer: layer,
		Fn: func() {
			// if layer == models.LayerContent {
			rl.DrawRectangle(cmd.PosX, cmd.PosY, cmd.Width, cmd.Height, cmd.Col)
			// } else {
			// 	screenPos := rl.GetWorldToScreen2D(rl.NewVector2(float32(cmd.PosX), float32(cmd.PosY)), *cam)
			// 	rl.DrawRectangle(int32(screenPos.X), int32(screenPos.Y), cmd.Width, cmd.Height, cmd.Col)
			// }
		},
	})
}

func (r *RaylibRender) Flush() {
	sort.Slice(r.queue, func(i, j int) bool {
		return r.queue[i].Layer < r.queue[j].Layer
	})

	// world-space: начинаем BeginMode2D один раз
	rl.BeginMode2D(*r.cam)
	for _, cmd := range r.queue {
		if cmd.Layer == models.LayerContent {
			cmd.Fn()
		}
	}
	rl.EndMode2D()

	for _, cmd := range r.queue {
		if cmd.Layer != models.LayerContent {
			cmd.Fn()
		}
	}

	r.queue = r.queue[:0]
}

func (r *RaylibRender) Submit(cmd DrawCmd) {
	r.queue = append(r.queue, cmd)
}
