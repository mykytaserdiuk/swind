package bus

type MessageBus struct {
	subs map[EventType][]func(Event)
}

func NewMessageBus() *MessageBus {
	return &MessageBus{
		subs: make(map[EventType][]func(Event)),
	}
}

func (b *MessageBus) Subscribe(t EventType, h func(Event)) {
	b.subs[t] = append(b.subs[t], h)
}

func (b *MessageBus) Emit(e Event) {
	if handlers, ok := b.subs[e.Type]; ok {
		for _, h := range handlers {
			h(e)
		}
	}
}
