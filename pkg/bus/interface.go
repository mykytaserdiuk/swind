package bus

type EventType string

const (
	SwitchScene EventType = "switch_scene"
	ButtonClick EventType = "button_click"
	UIEvent     EventType = "ui_event"
	StateUpdate EventType = "new_state"
)

type Event struct {
	Type EventType
	Data any
}
type Bus interface {
	Emit(event Event)
	Subscribe(eventType EventType, handler func(Event))
}
