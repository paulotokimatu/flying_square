package obstacle

import (
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
)

type ObstacleEnt struct {
	flat_game.IEntity
	velocityX float32
}

func NewObstacleEnt(config *entity.Config, velocityX float32) *ObstacleEnt {
	entity := entity.NewBaseEntity(config)

	return &ObstacleEnt{entity, velocityX}
}

func (ent *ObstacleEnt) OnCollision(game flat_game.IGame, externalEntity flat_game.IEntity) {
	menu := game.SceneByName("menu")

	game.SetScene(menu, true)
}

func (ent *ObstacleEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	newPosition := ent.Position()

	newPosition.SetX(newPosition.X + (ent.velocityX * delta))

	if newPosition.X+ent.Size().X < 0 {
		newPosition.SetX(0)

		ent.SetPendingRemoval(true)
	} else {
		ent.SetPosition(newPosition)
	}
}
