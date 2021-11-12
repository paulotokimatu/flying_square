package level

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/input"
	"flat_game/utils"
	"flying_square/obstacle"
)

func NewLevel(flyingSquare flat_game.IGame, levelName string) (flat_game.IScene, error) {
	config := flyingSquare.Config()
	newLevel := entity.NewScene(&entity.Config{Name: "level"})

	squareTexture, err := flyingSquare.AddTexture("square", "assets/square.png")
	if err != nil {
		return nil, err
	}
	_, err = flyingSquare.AddTexture("obstacle", "assets/obstacle.png")
	if err != nil {
		return nil, err
	}
	bgTexture, err := flyingSquare.AddTexture("bg", "assets/bg.png")
	if err != nil {
		return nil, err
	}

	bgEnt := entity.NewSpriteEnt(&entity.Config{
		Name: "bg",
		Size: utils.Vec2{X: config.Size.X, Y: config.Size.Y},
	}, bgTexture, false)

	var squarePositionX float32

	if levelName == "right" {
		squarePositionX = config.Size.X - 130
	} else {
		squarePositionX = 50
	}

	squareEnt := NewSquareEnt(&entity.Config{
		Name:     "square",
		Position: utils.Vec2{X: squarePositionX, Y: 100},
		Size:     utils.Vec2{X: 80, Y: 80},
	}, squareTexture)
	squareMovementEnt, _ := squareEnt.ChildByName("movement").(input.IKeyEventListener)
	newLevel.AddKeyEventListener(squareMovementEnt)

	obstacleGeneratorEnt := obstacle.NewGenerator(&entity.Config{
		Name:     "obstacle_generator",
		Position: utils.Vec2{X: 0, Y: 0},
		Size:     utils.Vec2{X: 0, Y: 0},
	}, 2, levelName)

	newLevel.AddChild(bgEnt)
	newLevel.AddChild(obstacleGeneratorEnt)
	newLevel.AddChild(squareEnt)

	return newLevel, nil
}
