package menu

import (
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/input"
)

type selectSceneFunc func(game flat_game.IGame, sceneName string)

type SelectionEnt struct {
	flat_game.IEntity
	selectScene selectSceneFunc
}

func NewSelectionEnt(config *entity.Config, selectScene selectSceneFunc) *SelectionEnt {
	base := entity.NewBaseEntity(config)
	return &SelectionEnt{base, selectScene}
}

func (ent *SelectionEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	if game.Graphics().IsKeyPressed(input.KeyLeft) {
		ent.selectScene(game, "left")
	} else if game.Graphics().IsKeyPressed(input.KeyRight) {
		ent.selectScene(game, "right")
	}
}
