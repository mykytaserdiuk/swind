package models

const (
	SwitchScene = "switch_scene"
	ButtonClick = "button_click"
	UIevent     = "ui_event"
	StateUpdate = "new_state"
)

type Event struct {
	Type string
	Data any
}
