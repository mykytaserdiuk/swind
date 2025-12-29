package scene

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mykytaserdiuk/fluxo"
	"github.com/nikitaserdiuk9/swind/pkg/element"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type MenuScene struct {
	name string

	r        render.Renderer
	elements []element.Base
	b        fluxo.Bus

	workArea rl.Rectangle
}

func NewMenuScene(renderer render.Renderer, b fluxo.Bus) *MenuScene {
	menu := &MenuScene{
		name: "menu",
		b:    b,
		r:    renderer,
		elements: []element.Base{element.NewButton(
			rl.Rectangle{
				X:      float32(1280 - 222),
				Y:      0,
				Width:  222,
				Height: 222,
			},
			b,
			// utils.ClampToWorkArea(rl.Rectangle{
			// 	X:      25,
			// 	Y:      25,
			// 	Width:  100,
			// 	Height: 100,
			// }, workArea),
		)},
	}
	return menu
}

func (s *MenuScene) Name() string {
	return s.name
}

func (s *MenuScene) OnEnter() {
	fmt.Println("Menu Enter")

	s.b.Subscribe(models.UIevent, func(e models.Event) {
		fmt.Println("EEE")
		s.onEvent(e)
	})

}

func (s *MenuScene) OnExit() {
	fmt.Println("Menu Exit")
	// Unsubscribe
}

func (s *MenuScene) Update(dt float32) {
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

func (s *MenuScene) Draw() {
	for _, el := range s.elements {
		el.Draw(s.r)
	}

	s.r.Flush()
}

func (sm *MenuScene) onEvent(e models.Event) {
	switch ev := e.Data.(type) {
	case models.UIEvent:
		switch ev.Type {
		case "button_click":
			if ev.ID == "Exit" {
				sm.b.Emit(models.SwitchScene, models.Event{Type: models.SwitchScene, Data: "game"})
			}
		}
	}
}
