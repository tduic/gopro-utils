package telemetry

import (
	"encoding/binary"
	"errors"
	"math"
)

// Microsecond Timestamp
type STMP struct {
	Stmp uint32
}

func (stmp *STMP) Parse(bytes []byte) error {
	if 4 != len(bytes) {
		return errors.New("Invalid length STMP packet")
	}

	bits := binary.BigEndian.Uint32(bytes[0:4])

	stmp.Stmp = uint32(math.Float32frombits(bits))

	return nil
}
