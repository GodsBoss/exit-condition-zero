// build js,wasm

package game

import (
	"sync"

	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type Game struct {
	scale int

	currentStateID string
	states         map[string]State

	output *dom.Context2D

	lock sync.Mutex
}

type SpriteMap interface {
	Produce(id string, x, y int, scale int, frame int) Renderable
}

func New(
	initialStateID string,
	states map[string]State,
) *Game {
	states[initialStateID].Init()
	return &Game{
		states:         states,
		currentStateID: initialStateID,
		scale:          1,
	}
}

func (g *Game) currentState() State {
	return g.states[g.currentStateID]
}

func (g *Game) TicksPerSecond() int {
	return 40
}

func (g *Game) SetOutput(ctx2d *dom.Context2D) {
	g.output = ctx2d
}

func (g *Game) Render() {
	g.output.DisableImageSmoothing()
	g.output.ClearRect(0, 0, g.scale*uiWidth, g.scale*uiHeight)
	g.lock.Lock()
	renderables := g.currentState().Renderables(g.scale)
	g.lock.Unlock()
	for i := range renderables {
		renderables[i].Render(g.output)
	}
}

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
	g.lock.Lock()
	g.transition(g.currentState().Tick(ms))
	g.lock.Unlock()
}

func (g *Game) ReceiveKeyEvent(event interaction.KeyEvent) {
	g.lock.Lock()
	g.transition(g.currentState().ReceiveKeyEvent(event))
	g.lock.Unlock()
}

func (g *Game) ReceiveMouseEvent(event interaction.MouseEvent) {
	g.lock.Lock()
	g.transition(g.currentState().ReceiveMouseEvent(event))
	g.lock.Unlock()
}

func (g *Game) transition(trans *Transition) {
	if trans == nil {
		return
	}
	g.currentStateID = trans.NextState
	g.currentState().Init()
}

const (
	uiWidth  = 320
	uiHeight = 240
)
