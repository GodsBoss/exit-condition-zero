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
	"level_select_level_selected": {
		X: 400,
		Y: 96,
		W: 48,
		H: 48,
	},
	"output_up": {
		X: 0,
		Y: 20,
		W: 20,
		H: 20,
	},
	"output_right": {
		X: 20,
		Y: 20,
		W: 20,
		H: 20,
	},
	"output_down": {
		X: 40,
		Y: 20,
		W: 20,
		H: 20,
	},
	"output_left": {
		X: 60,
		Y: 20,
		W: 20,
		H: 20,
	},
	"playing_button_run": {
		X: 300,
		Y: 0,
		W: 20,
		H: 20,
	},
	"playing_button_reset": {
		X: 320,
		Y: 0,
		W: 20,
		H: 20,
	},
	"playing_button_exit": {
		X: 340,
		Y: 0,
		W: 20,
		H: 20,
	},
}
