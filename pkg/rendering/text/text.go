package text

import (
	"strings"

	"github.com/GodsBoss/exit-condition-zero/pkg/rendering"
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
	chars := make(rendering.Renderables, 0)
	width := 0
	height := len(lines)

	for row := range lines {
		if len(lines[row]) > width {
			width = len(lines[row])
		}
		for col := range lines[row] {
			chars = append(
				chars,
				txt.spriteMap.Produce(
					"char_"+string(lines[row][col]),
					txt.x+col*6,
					txt.y+row*6,
					txt.scale,
					0,
				),
			)
		}
	}

	backgroundTiles := make(rendering.Renderables, 0)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			backgroundTiles = append(
				backgroundTiles,
				txt.spriteMap.Produce(
					"text_border_center",
					txt.x+col*6,
					txt.y+row*6,
					txt.scale,
					0,
				),
			)
		}
	}
	for row := 0; row < height; row++ {
		backgroundTiles = append(
			backgroundTiles,
			txt.spriteMap.Produce(
				"text_border_left",
				txt.x-3,
				txt.y+row*6,
				txt.scale,
				0,
			),
			txt.spriteMap.Produce(
				"text_border_right",
				txt.x+width*6,
				txt.y+row*6,
				txt.scale,
				0,
			),
		)
	}
	for col := 0; col < width; col++ {
		backgroundTiles = append(
			backgroundTiles,
			txt.spriteMap.Produce(
				"text_border_top",
				txt.x+col*6,
				txt.y-3,
				txt.scale,
				0,
			),
			txt.spriteMap.Produce(
				"text_border_bottom",
				txt.x+col*6,
				txt.y+height*6,
				txt.scale,
				0,
			),
		)
	}

	backgroundTiles = append(
		backgroundTiles,
		txt.spriteMap.Produce(
			"text_border_left_top",
			txt.x-3,
			txt.y-3,
			txt.scale,
			0,
		),
		txt.spriteMap.Produce(
			"text_border_left_bottom",
			txt.x-3,
			txt.y+height*6,
			txt.scale,
			0,
		),
		txt.spriteMap.Produce(
			"text_border_right_top",
			txt.x+width*6,
			txt.y-3,
			txt.scale,
			0,
		),
		txt.spriteMap.Produce(
			"text_border_right_bottom",
			txt.x+width*6,
			txt.y+height*6,
			txt.scale,
			0,
		),
	)

	backgroundTiles.Render(output)
	chars.Render(output)
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
