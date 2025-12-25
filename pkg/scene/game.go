package scene

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/nikitaserdiuk9/swind/pkg/bus"
	"github.com/nikitaserdiuk9/swind/pkg/element"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type GameScene struct {
	name string

	r        render.Renderer
	elements []element.Base
	b        bus.Bus

	text     string
	workArea rl.Rectangle
}

func NewGameScene(renderer render.Renderer, b bus.Bus) *GameScene {
	text := "WACHALA"
	menu := &GameScene{
		name: "game",
		b:    b,
		text: text,
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
				len(text)+1,
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

func (s *GameScene) matchPercent(input string) (percent int, ok bool) {
	runesA := []rune(s.text)
	runesB := []rune(input)
	if len(runesA) == len(runesB) && len(runesA) > 0 {
		matches := 0
		for i := range runesA {
			if runesA[i] == runesB[i] {
				matches++
			}
		}
		percent = int(float64(matches) * 100.0 / float64(len(runesA)))
		return percent, true
	}
	return 0, false
}

func (s *GameScene) OnEnter() {
	fmt.Println("Game Enter")
	s.b.Subscribe(bus.StateUpdate, func(e bus.Event) {
		if data, ok := e.Data.(models.UIEvent); ok {
			if data.Type == "input_text" {
				if percent, ok := s.matchPercent(data.ID); ok {
					s.b.Emit(bus.Event{
						Type: bus.StateUpdate,
						Data: GameoverEvent{Score: percent},
					})
				}
			}
		}
	})
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
