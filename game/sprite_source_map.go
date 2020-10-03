package main

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
)

var spriteSources = sprite.SourceMap{
	"bg_title": {
		X: 480,
		Y: 0,
		W: 320,
		H: 240,
	},
	"bg_level_select": {
		X: 480,
		Y: 240,
		W: 320,
		H: 240,
	},
	"bg_playing": {
		X: 480,
		Y: 480,
		W: 320,
		H: 240,
	},
	"bg_game_over": {
		X: 480,
		Y: 720,
		W: 320,
		H: 240,
	},
	"level_select_level": {
		X: 400,
		Y: 0,
		W: 48,
		H: 48,
	},
	"level_select_level_hover": {
		X: 400,
		Y: 48,
		W: 48,
		H: 48,
	},
}
