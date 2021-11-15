package menu_test

import (
	"flying_square/menu"
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/game"
	"github.com/paulotokimatu/flat_game/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SelectionEntMockGame struct {
	game.Game
	graphics flat_game.IGraphics
}

func (game SelectionEntMockGame) Graphics() flat_game.IGraphics {
	return game.graphics
}

type SelectionEntMockGraphics struct {
	flat_game.IGraphics
}

func (graphics SelectionEntMockGraphics) IsKeyPressed(key input.Key) bool {
	return key == input.KeyRight
}

func TestShouldTickEntities(t *testing.T) {
	mockGraphics := SelectionEntMockGraphics{}
	mockGame := &SelectionEntMockGame{
		graphics: mockGraphics,
	}

	selectionEnt := menu.NewSelectionEnt(&entity.Config{Name: "foo"}, sceneSelected)

	assert.Nil(t, mockGame.CurrentScene())

	selectionEnt.Tick(mockGame, nil, 1)

	assert.Equal(t, "right", mockGame.CurrentScene().Name())
}

func sceneSelected(game flat_game.IGame, sceneName string) {
	newScene := entity.NewScene(&entity.Config{Name: sceneName})

	game.SetScene(newScene, false)
}
