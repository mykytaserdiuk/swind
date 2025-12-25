package element

import "github.com/nikitaserdiuk9/swind/pkg/render"

type Base interface {
	Update(dt float32)
	Draw(r render.Renderer)
	HandleInput() bool
}
