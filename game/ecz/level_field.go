package ecz

// field is a level's field when playing.
type field interface {
	// Resets this field to its initial state as found in the level's data (for
	// pre-existing fields) or configured by the player (for fields set by the player).
	Reset()

	// ExtractOutputPulses takes all the output pulses this field has to offer.
	// This may change the field, e.g. "exhaust" pulses.
	ExtractOutputPulses() []direction

	// ImmediateHit takes a beam from a direction. Returned directions are directly
	// converted to new pulses. The field should not change state.
	// The first return parameter determines wether this field was "hit" in a
	// meaningful sense. The returned directions are handled wether the field was
	// hit or not, so if a pulse is to be continued, it should be returned.
	ImmediateHit(direction) (bool, []direction)

	// Receive is called after all pulses have ended somewhere and a pulse hit
	// this field. ImmediateHit() must have returned true for Receive() to be called.
	Receive([]direction)

	// IsDeletable determines wether the player can delete this field.
	IsDeletable() bool

	// IsMovable determines wether the player can move this field.
	IsMovable() bool
}

type direction string

const (
	dirUp    direction = "up"
	dirRight direction = "right"
	dirDown  direction = "down"
	dirLeft  direction = "left"
)

// fieldWithVictoryCondition is an optional interface fields can implement.
// Every field will be questioned for victory when running. If any field
// disallows a victory, the level is not over.
type fieldWithVictoryCondition interface {
	// AllowsVictory returns true if this field allows for the player to be
	// victorious in this level.
	AllowsVictory() bool
}
