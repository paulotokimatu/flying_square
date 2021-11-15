package obstacle_test

import (
	"flying_square/obstacle"
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/game"
	"github.com/paulotokimatu/flat_game/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GeneratorEntMockGame struct {
	game.Game
}

func (game GeneratorEntMockGame) Config() flat_game.Config {
	return flat_game.Config{
		Size: utils.Vec2{X: 1000, Y: 1000},
	}
}

func TestTickShouldGenerateObstacle(t *testing.T) {
	testGame := GeneratorEntMockGame{}
	testScene := entity.NewScene(&entity.Config{Name: "scene"})

	testGame.SetScene(testScene, false)

	generator := obstacle.NewGenerator(&entity.Config{Name: "foo"}, 11, "left")

	generator.Tick(&testGame, nil, 6)

	assert.Equal(t, 0, len(testScene.ChildrenToAdd()))

	generator.Tick(&testGame, nil, 6)

	assert.Equal(t, 2, len(testScene.ChildrenToAdd()))

	testScene.CommitChild(testScene.ChildrenToAdd()[0])
	testScene.CommitChild(testScene.ChildrenToAdd()[1])

	obstacleUp := testScene.ChildByName("obstacle-up0")
	obstacleDown := testScene.ChildByName("obstacle-down0")

	assert.True(t, obstacleUp.Size().Y+obstacleDown.Size().Y < testGame.Config().Size.Y)
}
