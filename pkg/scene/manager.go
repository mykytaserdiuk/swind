package scene

import (
	"fmt"

	"github.com/mykytaserdiuk/fluxo"
	"github.com/nikitaserdiuk9/swind/pkg/models"
	"github.com/nikitaserdiuk9/swind/pkg/render"
)

type SceneManager struct {
	scenes       map[string]Scene
	currentScene Scene
	bus          fluxo.Bus
}

func NewSceneManager(r render.Renderer, b fluxo.Bus) *SceneManager {
	menuScene := NewMenuScene(r, b)
	gameScene := NewGameScene(r, b)
	sceneManager := &SceneManager{
		currentScene: nil,
		bus:          b,
		scenes: map[string]Scene{
			menuScene.Name(): menuScene,
			gameScene.Name(): gameScene,
		},
	}

	sceneManager.bus.Subscribe(models.SwitchScene, func(e models.Event) {
		if name, ok := e.Data.(string); ok {
			sceneManager.SwitchScene(name)
		} else {
			fmt.Println("Unvalid scene name: ", e.Data)
		}
	})
	sceneManager.bus.Subscribe(models.StateUpdate, func(e models.Event) {
		if event, ok := e.Data.(GameoverEvent); ok {
			fmt.Println(event)
		}
	})

	return sceneManager
}

func (sm *SceneManager) Update(dt float32) {
	sm.currentScene.Update(dt)
}

func (sm *SceneManager) Draw() {
	sm.currentScene.Draw()
}

func (sm *SceneManager) SwitchScene(name string) {
	if sm.currentScene != nil {
		sm.currentScene.OnExit()
	}
	sm.currentScene = sm.scenes[name]
	sm.currentScene.OnEnter()
}
