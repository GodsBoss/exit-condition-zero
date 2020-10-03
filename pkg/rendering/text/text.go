package text

import (
	"strings"

	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/gggg/pkg/dom"
)

type Text struct {
	spriteMap sprite.Map
	contents  string
	x         int
	y         int
	scale     int
}

func New(
	spriteMap sprite.Map,
	contents string,
	x int,
	y int,
	scale int,
) *Text {
	return &Text{
		spriteMap: spriteMap,
		contents:  strings.Map(allowChar, strings.ToLower(contents)),
		x:         x,
		y:         y,
		scale:     scale,
	}
}

func (txt *Text) Render(output *dom.Context2D) {
	lines := strings.Split(txt.contents, "\n")
	for row := range lines {
		for col := range lines[row] {
			txt.spriteMap.Produce(
				"char_"+string(lines[row][col]),
				txt.x+col*6,
				txt.y+row*6,
				txt.scale,
				0,
			).Render(output)
		}
	}
}

const allowedChars = "\nabcdefghijklmnopqrstuvwxyz 01234567890.,:;!?()[]-=><_+'\"%|~#/\\"

var allowedCharsMap = func() map[rune]struct{} {
	m := map[rune]struct{}{}
	for _, r := range []rune(allowedChars) {
		m[r] = struct{}{}
	}
	return m
}()

func allowChar(r rune) rune {
	if _, ok := allowedCharsMap[r]; ok {
		return r
	}
	return -1
}
