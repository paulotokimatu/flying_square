package main

import (
	"flat_game"
	"flat_game/game"
	"flying_square/menu"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	config := flat_game.NewConfigFromJson("config.json")

	flyingSquare := game.NewGame(config)

	start(flyingSquare)

	flyingSquare.Run()
}

func start(flyingSquare flat_game.IGame) {
	menuScene, err := menu.NewMenuScene(flyingSquare)
	if err != nil {
		panic(err)
	}

	flyingSquare.SetScene(menuScene, false)
}
