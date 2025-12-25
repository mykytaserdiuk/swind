package models

type Layer int

const (
	LayerBackground Layer = iota
	LayerContent
	LayerUI
	LayerOverlay
)
