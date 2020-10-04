package ecz

import (
	"math"
	"math/rand"

	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/text"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

// gameOver is the state showing that the level was finished successfully.
// A "Game over" in a failure sense does not exist.
type gameOver struct {
	spriteMap sprite.Map
	movers    []*gameOverMover
}

func newGameOver(spriteMap sprite.Map) game.State {
	return &gameOver{
		spriteMap: spriteMap,
	}
}

func (over *gameOver) Init() {
	over.movers = make([]*gameOverMover, gameOverMoverCount)
	for i := 0; i < gameOverMoverCount; i++ {
		over.movers[i] = newGameOverMover(
			int2d.FromXY(160, 120),
			100,
			float64(i)*2*math.Pi/float64(gameOverMoverCount),
			0.25*float64((1-2*(i%2)))*(0.5*rand.Float64()+0.5),
		)
	}
}

func (over *gameOver) Tick(ms int) *game.Transition {
	for i := range over.movers {
		over.movers[i].tick(ms)
	}
	return nil
}

func (over *gameOver) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (over *gameOver) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseUp && event.PrimaryButton() {
		return &game.Transition{
			NextState: "title",
		}
	}
	return nil
}

func (over *gameOver) Renderables(scale int) []game.Renderable {
	r := []game.Renderable{
		over.spriteMap.Produce("bg_game_over", 0, 0, scale, 0),
		text.New(over.spriteMap, "Congratulations!\nYou managed to break\nthe loop successfully!", 95, 100, scale),
	}
	for i := range over.movers {
		r = append(r, over.movers[i].renderable(over.spriteMap, scale))
	}
	return r
}

type gameOverMover struct {
	center         int2d.Vector
	distance       float64
	angle          float64
	anglePerSecond float64
	anim           *animation
}

func newGameOverMover(center int2d.Vector, distance float64, initialAngle float64, anglePerSecond float64) *gameOverMover {
	return &gameOverMover{
		center:         center,
		distance:       distance,
		angle:          initialAngle,
		anglePerSecond: anglePerSecond,
		anim: &animation{
			progress: rand.Float64() * 4.0,
			fps:      4,
			frames:   4,
		},
	}
}

func (mover *gameOverMover) tick(ms int) {
	mover.anim.tick(ms)

	mover.angle = mover.angle + mover.anglePerSecond*float64(ms)/1000.0
	if mover.angle > 2*math.Pi {
		mover.angle -= 2 * math.Pi
	}
}

func (mover *gameOverMover) renderable(spriteMap sprite.Map, scale int) game.Renderable {
	frame := mover.anim.frame()
	if frame == 3 {
		frame = 1
	}
	return spriteMap.Produce(
		"level_over_mover",
		mover.center.X()+int(mover.distance*math.Sin(mover.angle))-10,
		mover.center.Y()+int(mover.distance*math.Cos(mover.angle))-10,
		scale,
		frame,
	)
}

const gameOverMoverCount = 18
