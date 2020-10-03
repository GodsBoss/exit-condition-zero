package ecz

import (
	"github.com/GodsBoss/exit-condition-zero/pkg/game"
	"github.com/GodsBoss/exit-condition-zero/pkg/rect"
	"github.com/GodsBoss/exit-condition-zero/pkg/rendering/sprite"
	"github.com/GodsBoss/exit-condition-zero/pkg/vector/int2d"

	"github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
	spriteMap sprite.Map
	levels    *levels

	gridCursor           *int2d.Vector
	isDeleteMode         bool
	isMoveMode           bool
	fieldSelectedForMove *int2d.Vector
	selectableFields     []int2d.Vector

	running             bool
	msUntilNextBeamStep int
	beams               map[beamIndex]struct{}
	pulses              []*pulse
	acceptedPulses      map[int2d.Vector]map[direction]struct{}
	firstHalf           bool
	gameOver            bool
	fields              map[int2d.Vector]field
}

func newPlaying(spriteMap sprite.Map, levels *levels) game.State {
	return &playing{
		spriteMap: spriteMap,
		levels:    levels,
	}
}

func (p *playing) Init() {
	p.running = false
	p.gameOver = false
	p.initRunningValues()
	p.fields = make(map[int2d.Vector]field)
	for x := 0; x < 11; x++ {
		for y := 0; y < 11; y++ {
			p.fields[int2d.FromXY(x, y)] = &emptyField{
				spriteMap: p.spriteMap,
			}
		}
	}

	lvlFields := p.levels.levels[p.levels.selectedLevel].getFields(p.spriteMap)
	for v := range lvlFields {
		p.fields[v] = lvlFields[v]
	}
}

func (p *playing) Tick(ms int) *game.Transition {
	if p.gameOver {
		return &game.Transition{
			NextState: "game_over",
		}
	}
	if p.running {
		p.msUntilNextBeamStep -= ms
		if p.msUntilNextBeamStep <= 0 {
			p.msUntilNextBeamStep = msPerBeamStep
			p.beamStep()
		}
	}
	return nil
}

func (p *playing) ReceiveKeyEvent(event interaction.KeyEvent) *game.Transition {
	return nil
}

func (p *playing) ReceiveMouseEvent(event interaction.MouseEvent) *game.Transition {
	if event.Type == interaction.MouseUp && event.PrimaryButton() {

		// Exit button
		if rect.FromPositionAndSize(295, 215, 20, 20).Inside(event.X, event.Y) {
			return &game.Transition{
				NextState: "title",
			}
		}

		// Start/Stop button
		if rect.FromPositionAndSize(245, 215, 20, 20).Inside(event.X, event.Y) {
			p.toggleRun()
		}

		// Delete button
		if rect.FromPositionAndSize(245, 5, 20, 20).Inside(event.X, event.Y) && !p.running {
			p.toggleDeleteMode()
		}

		// Move button
		if rect.FromPositionAndSize(270, 5, 20, 20).Inside(event.X, event.Y) && !p.running {
			p.toggleMoveMode()
		}

		// Something on the grid.
		if p.gridCursor != nil {
			if p.isDeleteMode {
				p.attemptToDelete()
				return nil
			}
			if p.isMoveMode {
				p.attemptToMove()
				return nil
			}
		}
	}

	if event.Type == interaction.MouseMove {
		p.gridCursor = nil
		gridPos := pointerPositionToGrid(int2d.FromXY(event.X, event.Y))
		if gridPos.X() >= 0 && gridPos.X() < 11 && gridPos.Y() >= 0 && gridPos.Y() < 11 {
			p.gridCursor = &gridPos
		}
	}

	return nil
}

func (p *playing) toggleDeleteMode() {
	p.clearMoveMode()
	p.fieldSelectedForMove = nil
	if p.isDeleteMode {
		p.isDeleteMode = false
		return
	}
	p.isDeleteMode = true
	p.findSelectableFields(
		func(f field) bool {
			return f.IsDeletable()
		},
	)
}

func (p *playing) attemptToDelete() {
	v := *p.gridCursor
	if p.fields[v].IsDeletable() {
		p.fields[v] = &emptyField{
			spriteMap: p.spriteMap,
			free:      true,
		}
	}
	p.isDeleteMode = false
}

func (p *playing) clearMoveMode() {
	p.isMoveMode = false
	p.fieldSelectedForMove = nil
}

func (p *playing) toggleMoveMode() {
	p.isDeleteMode = false
	if p.isMoveMode {
		p.clearMoveMode()
		return
	}
	p.isMoveMode = true
	p.findSelectableFields(
		func(f field) bool {
			return f.IsMovable()
		},
	)
}

func (p *playing) attemptToMove() {
	v := *p.gridCursor

	// Select field.
	if p.fieldSelectedForMove == nil {
		if p.fields[v].IsMovable() {
			p.fieldSelectedForMove = &v
			p.findSelectableFields(isValidMoveDestination)
			return
		}
		p.clearMoveMode()
		return
	}

	// Check wether fields to move are the same. In that case, do not move anything.
	if v == *p.fieldSelectedForMove {
		p.clearMoveMode()
		return
	}

	destField := p.fields[v]

	if isValidMoveDestination(destField) {
		p.fields[v] = p.fields[*p.fieldSelectedForMove]
		if isFieldFree(destField) || destField.IsMovable() {
			p.fields[*p.fieldSelectedForMove] = destField
		} else {
			p.fields[*p.fieldSelectedForMove] = &emptyField{
				spriteMap: p.spriteMap,
				free:      true,
			}
		}
	}

	p.clearMoveMode()
}

func isValidMoveDestination(f field) bool {
	return isFieldFree(f) || f.IsMovable() || f.IsDeletable()
}

func (p *playing) clearSelectableFields() {
	p.selectableFields = make([]int2d.Vector, 0)
}

func (p *playing) findSelectableFields(criteria func(field) bool) {
	p.clearSelectableFields()
	for v := range p.fields {
		if criteria(p.fields[v]) {
			p.selectableFields = append(p.selectableFields, v)
		}
	}
}

func (p *playing) toggleRun() {
	if p.running {
		p.stopRunning()
	} else {
		p.startRunning()
	}
}

func (p *playing) startRunning() {
	p.running = true
	p.initRunningValues()
	p.extractPulses()
}

func (p *playing) initRunningValues() {
	p.firstHalf = true
	p.msUntilNextBeamStep = msPerBeamStep
	p.beams = make(map[beamIndex]struct{})
	p.pulses = make([]*pulse, 0)
	p.acceptedPulses = make(map[int2d.Vector]map[direction]struct{})
}

func (p *playing) extractPulses() {
	for v := range p.fields {
		dirs := p.fields[v].ExtractOutputPulses()
		for i := range dirs {
			p.pulses = append(
				p.pulses,
				&pulse{
					pos: v,
					dir: dirs[i],
				},
			)
		}
	}
}

func (p *playing) beamStep() {
	if len(p.pulses) == 0 {
		p.pulsesExhausted()
		return
	}

	// Remove pulses duplicating already existing beams. A beam is uniquely
	// identified by position and direction.
	// This is only needed in the first half of beam creation.
	if p.firstHalf {
		leftOverPulses := make([]*pulse, 0)
		for i := range p.pulses {
			pulse := p.pulses[i]
			if _, ok := p.beams[beamIndex{v: pulse.pos, d: pulse.dir, firstHalf: true}]; !ok {
				leftOverPulses = append(leftOverPulses, pulse)
			}
		}
		p.pulses = leftOverPulses
	}

	// Create beams.
	for i := range p.pulses {
		p.beams[beamIndex{v: p.pulses[i].pos, d: p.pulses[i].dir, firstHalf: p.firstHalf}] = struct{}{}
	}

	// The second part of beam creation is the interesting one. Fields may be hit
	// immediately or will be remembered for later accepting beams.
	if !p.firstHalf {
		nextPulses := make([]*pulse, 0)

		for i := range p.pulses {
			puls := p.pulses[i]
			nextPos := realGridPosition(int2d.Add(puls.pos, directionVectors[puls.dir]))

			hit, nextDirs := p.fields[nextPos].ImmediateHit(puls.dir)

			// Mark this field as having accepted a pulse.
			if hit {
				if p.acceptedPulses[nextPos] == nil {
					p.acceptedPulses[nextPos] = make(map[direction]struct{})
				}
				p.acceptedPulses[nextPos][puls.dir] = struct{}{}
			}

			// Create new pulses, according to new directions from hit.
			for i := range nextDirs {
				nextPulses = append(
					nextPulses,
					&pulse{
						pos: nextPos,
						dir: nextDirs[i],
					},
				)
			}
		}

		p.pulses = nextPulses
	}

	p.firstHalf = !p.firstHalf
}

func (p *playing) pulsesExhausted() {
	for v := range p.acceptedPulses {
		dirs := make([]direction, 0)
		for dir := range p.acceptedPulses[v] {
			dirs = append(dirs, dir)
		}
		p.fields[v].Receive(dirs)
	}

	if p.hasWon() {
		p.running = false
		p.gameOver = true
		return
	}

	p.initRunningValues()
	p.extractPulses()
}

func (p *playing) hasWon() bool {
	for v := range p.fields {
		if victoryCondition, ok := p.fields[v].(fieldWithVictoryCondition); ok {
			if !victoryCondition.AllowsVictory() {
				return false
			}
		}
	}
	return true
}

func (p *playing) stopRunning() {
	p.running = false
	p.initRunningValues()
}

func (p *playing) Renderables(scale int) []game.Renderable {
	r := []game.Renderable{
		p.spriteMap.Produce("bg_playing", 0, 0, scale, 0),
		p.spriteMap.Produce("playing_button_reset", 270, 215, scale, 0),
		p.spriteMap.Produce("playing_button_exit", 295, 215, scale, 0),

		p.spriteMap.Produce("p_button_delete", 245, 5, scale, 0),
		p.spriteMap.Produce("p_button_move", 270, 5, scale, 0),
		p.spriteMap.Produce("p_button_configure", 295, 5, scale, 0),
	}

	if p.running {
		r = append(r, p.spriteMap.Produce("playing_button_stop", 245, 215, scale, 0))
	} else {
		r = append(r, p.spriteMap.Produce("playing_button_run", 245, 215, scale, 0))
	}

	for v := range p.fields {
		r = append(
			r,
			p.fields[v].Renderable(
				fieldsOffsetX+v.X()*fieldsWidth,
				fieldsOffsetY+v.Y()*fieldsHeight,
				scale,
			),
		)
	}

	for bi := range p.beams {
		pos := bi.v
		if !bi.firstHalf {
			pos = realGridPosition(int2d.Add(pos, directionVectors[bi.d]))
		}
		r = append(
			r,
			p.spriteMap.Produce(
				beamSpriteIDs[bi.firstHalf][bi.d],
				pos.X()*fieldsWidth+fieldsOffsetX-1,
				pos.Y()*fieldsHeight+fieldsOffsetY-1,
				scale,
				0,
			),
		)
	}

	if p.gridCursor != nil {
		r = append(
			r,
			p.spriteMap.Produce(
				"p_cursor",
				p.gridCursor.X()*fieldsWidth+fieldsOffsetX,
				p.gridCursor.Y()*fieldsHeight+fieldsOffsetY,
				scale,
				0,
			),
		)
	}

	if p.isDeleteMode || p.isMoveMode {
		for i := range p.selectableFields {
			r = append(
				r,
				p.spriteMap.Produce(
					"p_marker",
					p.selectableFields[i].X()*fieldsWidth+fieldsOffsetX,
					p.selectableFields[i].Y()*fieldsHeight+fieldsOffsetY,
					scale,
					0,
				),
			)
		}
	}

	if p.isDeleteMode {
		r = append(r, p.spriteMap.Produce("p_cursor", 245, 5, scale, 0))
	}

	if p.isMoveMode {
		r = append(r, p.spriteMap.Produce("p_cursor", 270, 5, scale, 0))
		if p.fieldSelectedForMove != nil {
			r = append(
				r,
				p.spriteMap.Produce(
					"p_cursor",
					(*p.fieldSelectedForMove).X()*fieldsWidth+fieldsOffsetX,
					(*p.fieldSelectedForMove).Y()*fieldsHeight+fieldsOffsetY,
					scale,
					0,
				),
			)
		}
	}

	return r
}

var beamSpriteIDs = map[bool]map[direction]string{
	true: {
		dirUp:    "p_beam_up_1",
		dirRight: "p_beam_right_1",
		dirDown:  "p_beam_down_1",
		dirLeft:  "p_beam_left_1",
	},
	false: {
		dirUp:    "p_beam_up_2",
		dirRight: "p_beam_right_2",
		dirDown:  "p_beam_down_2",
		dirLeft:  "p_beam_left_2",
	},
}

const fieldsOffsetX = 10
const fieldsOffsetY = 10
const fieldsWidth = 20
const fieldsHeight = 20

const msPerBeamStep = 50

type pulse struct {
	pos int2d.Vector
	dir direction
}

type beamIndex struct {
	v         int2d.Vector
	d         direction
	firstHalf bool
}

func realGridPosition(v int2d.Vector) int2d.Vector {
	x := v.X()
	y := v.Y()
	if x < 0 {
		x += 11
	}
	if y < 0 {
		y += 11
	}
	if x >= 11 {
		x -= 11
	}
	if y >= 11 {
		y -= 11
	}
	return int2d.FromXY(x, y)
}

func pointerPositionToGrid(v int2d.Vector) int2d.Vector {
	x := v.X() - fieldsOffsetX
	y := v.Y() - fieldsOffsetY
	if x < 0 || y < 0 {
		return int2d.FromXY(-1, -1)
	}
	return int2d.FromXY(
		x/fieldsWidth,
		y/fieldsHeight,
	)
}
