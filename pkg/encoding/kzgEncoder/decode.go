package kzgEncoder

import (
	rs "github.com/0glabs/0g-data-avail/pkg/encoding/encoder"
)

func (g *KzgEncoder) Decode(frames []Frame, indices []uint64, maxInputSize uint64) ([]byte, error) {
	rsFrames := make([]rs.Frame, len(frames))
	for ind, frame := range frames {
		rsFrames[ind] = rs.Frame{Coeffs: frame.Coeffs}
	}

	return g.Encoder.Decode(rsFrames, indices, maxInputSize)
}
