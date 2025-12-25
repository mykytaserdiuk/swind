package render

import (
	"image/color"

	"github.com/nikitaserdiuk9/swind/pkg/models"
)

type Renderer interface {
	Text(layer models.Layer, cmd TextRenderCmd)
	Rect(layer models.Layer, cmd RectRenderCmd)

	Submit(cmd DrawCmd)
	Flush()
}

type DrawCmd struct {
	Layer models.Layer
	Fn    func()
}

type RectRenderCmd struct {
	PosX, PosY, Width, Height int32
	Col                       color.RGBA
}

type TextRenderCmd struct {
	Text     string
	PosX     int32
	PosY     int32
	FontSize int32
	Col      color.RGBA
}
