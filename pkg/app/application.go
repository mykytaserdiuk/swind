package app

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mykytaserdiuk/fluxo"
	"github.com/nikitaserdiuk9/swind/pkg/input"
	"github.com/nikitaserdiuk9/swind/pkg/render"
	"github.com/nikitaserdiuk9/swind/pkg/scene"
)

type Application struct {
	windowWidth, windowHeight int32
	renderer                  *render.RaylibRender
	sceneManager              *scene.SceneManager
	camera                    *rl.Camera2D
	input                     *input.InputSystem
	bus                       fluxo.Bus
}

func NewApplication(width, height int32) *Application {
	rl.SetTargetFPS(60)

	cam := rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 1)

	bus := fluxo.NewEventBus()
	renderer := render.NewRaylibRender(&cam)
	sceneManager := scene.NewSceneManager(renderer, bus)
	input := input.NewInputSystem()
	return &Application{
		windowWidth:  width,
		windowHeight: height,
		renderer:     renderer,
		sceneManager: sceneManager,
		input:        input,
		camera:       &cam,
		bus:          bus,
	}
}

func (app *Application) Run() {
	rl.InitWindow(app.windowWidth, app.windowHeight, "SWIND")
	app.sceneManager.SwitchScene("menu")
	var dt float32

	defer func() {
		for {
			if err := recover(); err != nil { //catch
				// TODO write to log.txt
				fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
				os.Exit(1)
			}
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		dt = rl.GetFrameTime()

		app.input.Update()
		app.sceneManager.Update(dt)

		app.sceneManager.Draw()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
