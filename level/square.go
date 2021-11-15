package level

import (
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
)

type SquareEnt struct {
	flat_game.IEntity
}

func NewSquareEnt(config *entity.Config, squareTexture flat_game.ITexture) *SquareEnt {
	spriteEnt := entity.NewSpriteEnt(&entity.Config{Name: "sprite"}, squareTexture, true)

	movementEnt := NewMovementEnt(&entity.Config{
		Name: "movement",
	}, 12, 0, -8)

	config.Children = []flat_game.IEntity{
		spriteEnt,
		movementEnt,
	}

	baseEnt := entity.NewBaseEntity(config)

	return &SquareEnt{baseEnt}
}

func (ent *SquareEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	config := game.Config()

	if ent.Position().Y >= config.Size.Y-ent.Size().Y {
		menu := game.SceneByName("menu")

		game.SetScene(menu, true)
	}
}
