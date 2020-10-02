// build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type Game struct{}

func New() *Game {
	return &Game{}
}

func (g *Game) TicksPerSecond() int {
	return 40
}

func (g *Game) SetOutput(ctx2d *dom.Context2D) {}

func (g *Game) Render() {}

func (g *Game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	return availableWidth, availableHeight, 1.0, 1.0
}

func (g *Game) Tick(ms int) {}

func (g *Game) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (g *Game) ReceiveMouseEvent(event interaction.MouseEvent) {}
