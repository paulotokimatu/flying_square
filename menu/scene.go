package menu

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/utils"
	"flying_square/level"
)

func NewMenuScene(flyingSquare flat_game.IGame) (flat_game.IScene, error) {
	font, err := flyingSquare.AddFont("luxis", "assets/luxisr.ttf", 32, 127, 24)
	if err != nil {
		return nil, err
	}

	menuScene := entity.NewScene(&entity.Config{Name: "menu"})

	optionsLabel := entity.NewLabelEnt(
		&entity.Config{Name: "text", Position: utils.Vec2{X: 20, Y: 40}},
		font,
		"Select the game mode by pressing LEFT or RIGHT",
		&utils.Vec3{X: 1.0, Y: 1.0, Z: 1.0},
	)

	selectionConfig := &entity.Config{
		Name:     "selection",
		Children: []flat_game.IEntity{optionsLabel},
	}
	selectionEnt := NewSelectionEnt(selectionConfig, selectLevel)

	menuScene.AddChild(selectionEnt)

	return menuScene, nil
}

func selectLevel(game flat_game.IGame, levelName string) {
	selectedLevel, err := level.NewLevel(game, levelName)

	if err != nil {
		panic("error creating level")
	}

	game.SetScene(selectedLevel, false)
}
