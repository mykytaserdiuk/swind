package main

import "github.com/nikitaserdiuk9/swind/pkg/app"

func main() {
	app := app.NewApplication(1280, 720)
	app.Run()
}
