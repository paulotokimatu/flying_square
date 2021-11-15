package obstacle

import (
	"fmt"
	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/utils"
	"math"
	"math/rand"
)

type GeneratorEnt struct {
	flat_game.IEntity
	count                int
	lastSpawn            float32
	levelName            string
	secondsBetweenSpawns float32
}

func NewGenerator(config *entity.Config, secondsBetweenSpawns float32, levelName string) *GeneratorEnt {
	base := entity.NewBaseEntity(config)

	return &GeneratorEnt{
		IEntity:              base,
		secondsBetweenSpawns: secondsBetweenSpawns,
		lastSpawn:            0,
		count:                0,
		levelName:            levelName,
	}
}

func (ent *GeneratorEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	if delta+ent.lastSpawn > ent.secondsBetweenSpawns {
		scene := game.CurrentScene()

		square := scene.ChildByName("square")

		obstacleUp, obstacleDown := generateObstacle(game, ent.count, ent.levelName)

		scene.AddChild(obstacleUp)
		scene.AddChild(obstacleDown)

		scene.AddCollision(square, obstacleUp)
		scene.AddCollision(square, obstacleDown)

		ent.lastSpawn = 0
		ent.count += 1
	} else {
		ent.lastSpawn += delta
	}
}

func generateObstacle(game flat_game.IGame, count int, levelName string) (*ObstacleEnt, *ObstacleEnt) {
	gameConfig := game.Config()

	sizeMiddleSpace := (rand.Float32() * 100) + 75

	texture := game.TextureByName("obstacle")

	sizeUp, sizeDown := calcPartObstacleSize(gameConfig.Size.Y, sizeMiddleSpace)

	var positionX float32

	if levelName == "right" {
		positionX = -80
	} else {
		positionX = gameConfig.Size.X
	}

	velocityX := float32(-200)
	if levelName == "right" {
		velocityX = velocityX * -1
	}

	obstacleUp := generateObstacleParts(
		"obstacle-up"+fmt.Sprint(count),
		texture,
		positionX,
		0,
		sizeUp,
		velocityX,
	)

	obstacleDown := generateObstacleParts(
		"obstacle-down"+fmt.Sprint(count),
		texture,
		positionX,
		gameConfig.Size.Y-float32(math.Floor(float64(sizeDown))),
		sizeDown,
		velocityX,
	)

	return obstacleUp, obstacleDown
}

func calcPartObstacleSize(gameHeight float32, sizeMiddleSpace float32) (float32, float32) {
	obstacleTotalSize := calcTotalObstacleSize(gameHeight, sizeMiddleSpace)

	sizeUpPercentage := rand.Float64()

	if sizeUpPercentage < 0.1 {
		sizeUpPercentage = 0.1
	} else if sizeUpPercentage > 0.9 {
		sizeUpPercentage = 0.9
	}

	sizeDownPercentage := 1 - sizeUpPercentage

	return obstacleTotalSize * float32(sizeUpPercentage), obstacleTotalSize * float32(sizeDownPercentage)
}

func calcTotalObstacleSize(gameHeight float32, sizeMiddleSpace float32) float32 {
	minSize := gameHeight / 2

	partSize := (rand.Float32() * gameHeight)

	if partSize < minSize {
		return minSize
	} else if partSize > gameHeight-sizeMiddleSpace {
		return partSize - sizeMiddleSpace
	}

	return partSize
}

func generateObstacleParts(name string, texture flat_game.ITexture, positionX float32, positionY float32, sizeY float32, velocityX float32) *ObstacleEnt {
	config := entity.Config{
		Name:     name,
		Position: utils.Vec2{X: positionX, Y: positionY},
		Size:     utils.Vec2{X: 80, Y: sizeY},
	}

	obstacle := NewObstacleEnt(&config, velocityX)

	spriteExt := entity.NewSpriteEnt(&entity.Config{Name: "sprite"}, texture, true)
	obstacle.AddChild(spriteExt)

	return obstacle
}
