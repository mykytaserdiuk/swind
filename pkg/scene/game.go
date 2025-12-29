package scene

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mykytaserdiuk/fluxo"
	"github.com/nikitaserdiuk9/swind/pkg/element"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type GameScene struct {
	name string

	r        render.Renderer
	elements []element.Base
	b        fluxo.Bus

	text     string
	workArea rl.Rectangle
}

func NewGameScene(renderer render.Renderer, b fluxo.Bus) *GameScene {
	text := "WACHALA"
	menu := &GameScene{
		name: "game",
		b:    b,
		text: text,
		r:    renderer,
		// Layout: center main text and input horizontally, visualizer to the right
		elements: func() []element.Base {
			winW := float32(1280)
			winH := float32(720)
			// layout sizes
			mainW := float32(700)
			mainH := float32(120)
			inputW := mainW
			inputH := float32(64)
			visualW := float32(600)
			visualH := float32(160)

			// center everything horizontally
			centerX := (winW - mainW) / 2
			visualX := (winW - visualW) / 2

			// visualizator higher and centered, writable text centered below it, input at bottom
			visualY := float32(80)
			topY := visualY + visualH + 18
			inputY := winH - inputH - 72

			return []element.Base{
				element.NewTextVisualizator(
					rl.Rectangle{X: visualX, Y: visualY, Width: visualW, Height: visualH}, b, text),
				element.NewWritableText(
					rl.Rectangle{X: centerX, Y: topY, Width: mainW, Height: mainH},
					b, text,
				),
				element.NewTextInput(
					rl.Rectangle{X: centerX, Y: inputY, Width: inputW, Height: inputH},
					len(text)+1,
					b,
				),
			}
		}(),
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
	s.b.Subscribe(models.StateUpdate, func(e models.Event) {
		if data, ok := e.Data.(models.UIEvent); ok {
			if data.Type == "input_text" {
				if percent, ok := s.matchPercent(data.ID); ok {
					s.b.Emit(models.StateUpdate, models.Event{
						Type: models.StateUpdate,
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
// 	case models.ButtonClick:
// 		if event.ID == "Exit" {
// 			sm.b.Emit(models.Event{Type: models.SwitchScene, Data: "menu"})
// 		}
// 	}
// }
