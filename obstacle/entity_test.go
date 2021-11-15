package obstacle_test

import (
	"flying_square/obstacle"
	"testing"

	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/game"
	"github.com/paulotokimatu/flat_game/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ObstacleEntMockGame struct {
	game.Game
	mock.Mock
}

func (game *ObstacleEntMockGame) SceneByName(name string) flat_game.IScene {
	return nil
}

func (game *ObstacleEntMockGame) SetScene(scene flat_game.IScene, deletePreviousScene bool) {
	game.Called(nil, true)
}

func TestShouldChangeSceneOnCollision(t *testing.T) {
	obstacle := obstacle.NewObstacleEnt(&entity.Config{Name: "foo"}, -200)
	externalEnt := entity.NewBaseEntity(&entity.Config{Name: "foo"})

	game := ObstacleEntMockGame{}

	game.On("SetScene", nil, true).Return(nil)

	obstacle.OnCollision(&game, externalEnt)

	game.AssertCalled(t, "SetScene", nil, true)
}

func TestTickShouldMoveObstacle(t *testing.T) {
	obstacle := obstacle.NewObstacleEnt(&entity.Config{
		Name:     "foo",
		Position: utils.Vec2{X: 500, Y: 500},
	}, -100)

	assert.Equal(t, float32(500), obstacle.Position().X)

	obstacle.Tick(nil, nil, 1)

	assert.Equal(t, float32(400), obstacle.Position().X)
}

func TestTickShouldClampPosition(t *testing.T) {
	obstacle := obstacle.NewObstacleEnt(&entity.Config{
		Name:     "foo",
		Position: utils.Vec2{X: 50, Y: 500},
	}, -100)

	assert.Equal(t, float32(50), obstacle.Position().X)

	obstacle.Tick(nil, nil, 1)

	assert.Equal(t, float32(0), obstacle.Position().X)
	assert.True(t, obstacle.IsPendingRemoval())
}
