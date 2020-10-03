// build js,wasm

package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type Game struct {
	scale int
	state state

	output *dom.Context2D
}

func New() *Game {
	return &Game{
		state: &title{},
		scale: 1,
	}
}

func (g *Game) TicksPerSecond() int {
	return 40
}

func (g *Game) SetOutput(ctx2d *dom.Context2D) {
	ctx2d.DisableImageSmoothing()
	g.output = ctx2d
}

func (g *Game) Render() {}

func (g *Game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	sx := (availableWidth - 20) / uiWidth
	sy := (availableHeight - 20) / uiHeight
	if sx < 1 {
		sx = 1
	}
	if sy < 1 {
		sy = 1
	}
	g.scale = sx
	if sy < g.scale {
		g.scale = sy
	}
	return uiWidth * g.scale, uiHeight * g.scale, float64(g.scale), float64(g.scale)
}

func (g *Game) Tick(ms int) {
	g.state.Tick(ms)
}

func (g *Game) ReceiveKeyEvent(event interaction.KeyEvent) {
	g.state.ReceiveKeyEvent(event)
}

func (g *Game) ReceiveMouseEvent(event interaction.MouseEvent) {
	g.state.ReceiveMouseEvent(event)
}

const (
	uiWidth  = 320
	uiHeight = 240
)
