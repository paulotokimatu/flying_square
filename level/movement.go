package level

import (
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/input"
)

type MovementEnt struct {
	flat_game.IEntity
	gravity       float32
	jumpIntensity float32
	velocityY     float32
}

func NewMovementEnt(config *entity.Config, gravity float32, velocityY float32, jumpIntensity float32) *MovementEnt {
	entity := entity.NewBaseEntity(config)

	return &MovementEnt{
		IEntity:       entity,
		gravity:       gravity,
		jumpIntensity: jumpIntensity,
		velocityY:     velocityY,
	}
}

func (ent *MovementEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	newPosition := parent.Position()

	ent.velocityY += ent.gravity * delta

	newPosition.SetY(newPosition.Y + ent.velocityY)

	parent.SetPosition(newPosition)
}

func (ent *MovementEnt) OnKeyEvent(key input.Key, event input.KeyEvent) {
	if key == input.KeySpace && event == input.EventKeyReleased {
		ent.velocityY += ent.jumpIntensity
	}
}
