// build js,wasm

package sprite

import (
	"errors"

	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/gggg/pkg/dom"
)

type Map interface {
	Produce(id string, x, y int, scale int, frame int) game.Renderable
}

func NewMap(source *dom.Image, sourceMap SourceMap) Map {
	return &spriteMap{
		source:    source,
		sourceMap: sourceMap,
	}
}

type spriteMap struct {
	source    *dom.Image
	sourceMap SourceMap
}

func (m *spriteMap) Produce(id string, x, y int, scale int, frame int) game.Renderable {
	src, err := m.sourceMap.Get(id)
	if err != nil {
		return rendering.Null{}
	}
	return sprite{
		source: m.source,

		sourceX:      src.X + frame*src.W,
		sourceY:      src.Y,
		sourceWidth:  src.W,
		sourceHeight: src.H,

		destX:      x * scale,
		destY:      y * scale,
		destWidth:  src.W * scale,
		destHeight: src.H * scale,
	}
}

type SourceMap map[string]Source

func (m SourceMap) Get(id string) (Source, error) {
	s, ok := m[id]
	if !ok {
		return Source{}, errors.New("sprite " + id + " not found")
	}
	s.ID = id
	return s
}

type Source struct {
	ID string

	// X is the horizontal position in the source image.
	X int

	// Y is the vertical position in the source image.
	Y int

	// W is the width in the source image.
	W int

	// H is the height in the source image.
	H int

	// Frames is the number of frames for animations. Zero frames means no animation.
	Frames int
}

type sprite struct {
	source *dom.Image

	sourceX      int
	sourceY      int
	sourceWidth  int
	sourceHeight int

	destX      int
	destY      int
	destWidth  int
	destHeight int
}

func (s sprite) Render(output *dom.Context2D) {
	output.DrawImage(
		s.source,
		s.sourceX,
		s.sourceY,
		s.sourceWidth,
		s.sourceHeight,
		s.destX,
		s.destY,
		s.destWidth,
		s.destHeight,
	)
}
