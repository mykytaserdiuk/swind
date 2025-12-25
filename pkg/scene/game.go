package scene

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/bus"
	"github.com/nikitaserdiuk9/swind/pkg/element"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type GameScene struct {
	name string

	r        render.Renderer
	elements []element.Base
	b        bus.Bus

	workArea rl.Rectangle
}

func NewGameScene(renderer render.Renderer, b bus.Bus) *GameScene {
	text := "WACHALA"
	menu := &GameScene{
		name: "game",
		b:    b,
		r:    renderer,
		elements: []element.Base{
			element.NewWritableText(
				rl.Rectangle{
					X:      100,
					Y:      100,
					Width:  222,
					Height: 222,
				},
				b, text,
			),
			element.NewTextInput(
				rl.Rectangle{
					X:      100,
					Y:      333,
					Width:  222,
					Height: 100,
				},
				b,
			),
			element.NewTextVisualizator(
				rl.Rectangle{
					X:      float32(300),
					Y:      float32(300),
					Width:  222,
					Height: 222,
				}, b, text),
		},
	}

	return menu
}

func (s *GameScene) Name() string {
	return s.name
}

func (s *GameScene) OnEnter() {
	fmt.Println("Game Enter")
	// s.b.Subscribe(bus.UIEvent, func(e bus.Event) {
	// 	if event, ok := e.Data.(models.UIEvent); ok {
	// 		s.onEvent(event)
	// 	} else {
	// 		fmt.Println("Unvalid event: ", e.Data)
	// 	}
	// })
}

func (s *GameScene) OnExit() {
	fmt.Println("Game Exit")
}

func (s *GameScene) Update(dt float32) {
	for _, e := range s.elements {
		e.Update(dt)
	}

	for i := len(s.elements) - 1; i >= 0; i-- {
		e := s.elements[i]
		if e.HandleInput() {
			break
		}
	}
}

func (s *GameScene) Draw() {
	for _, el := range s.elements {
		el.Draw(s.r)
	}

	s.r.Flush()
}

// func (sm *GameScene) onEvent(event models.UIEvent) {
// 	switch event.EventType {
// 	case bus.ButtonClick:
// 		if event.ID == "Exit" {
// 			sm.b.Emit(bus.Event{Type: bus.SwitchScene, Data: "menu"})
// 		}
// 	}
// }
