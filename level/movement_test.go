package level_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/game"
	"flat_game/input"
	"flat_game/utils"
	"flying_square/level"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MovementEntMockGame struct {
	game.Game
}

func (game *MovementEntMockGame) Config() flat_game.Config {
	return flat_game.Config{
		Size: utils.Vec2{X: 100, Y: 100},
	}
}

func TestTickShouldMoveDueToGravity(t *testing.T) {
	squareMovement := level.NewMovementEnt(&entity.Config{
		Name:     "foo",
		Position: utils.Vec2{X: 0, Y: 0},
	}, 10, 0, 0)

	squareEnt := entity.NewBaseEntity(&entity.Config{
		Name:     "foo",
		Position: utils.Vec2{X: 0, Y: 0},
		Children: []flat_game.IEntity{squareMovement},
	})

	squareMovement.Tick(&MovementEntMockGame{}, squareEnt, 1)
	assert.Equal(t, float32(10), squareEnt.Position().Y)
}

func TestShouldMoveYOnSpaceRelease(t *testing.T) {
	squareMovement := level.NewMovementEnt(&entity.Config{
		Name:     "foo",
		Position: utils.Vec2{X: 0, Y: 0},
	}, 0, 0, -2)

	squareEnt := entity.NewBaseEntity(&entity.Config{
		Name:     "foo",
		Position: utils.Vec2{X: 0, Y: 0},
		Children: []flat_game.IEntity{squareMovement},
	})

	squareMovement.OnKeyEvent(input.KeyEnter, input.EventKeyReleased)
	squareMovement.Tick(&MovementEntMockGame{}, squareEnt, 1)
	assert.Equal(t, float32(0), squareEnt.Position().Y)

	squareMovement.OnKeyEvent(input.KeySpace, input.EventKeyPressed)
	squareMovement.Tick(&MovementEntMockGame{}, squareEnt, 1)
	assert.Equal(t, float32(0), squareEnt.Position().Y)

	squareMovement.OnKeyEvent(input.KeySpace, input.EventKeyReleased)
	squareMovement.Tick(&MovementEntMockGame{}, squareEnt, 1)
	assert.Equal(t, float32(-2), squareEnt.Position().Y)
}
