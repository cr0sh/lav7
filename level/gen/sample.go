package gen

import "github.com/L7-MCPE/lav7/level"

type SampleGenerator struct{}

func (s SampleGenerator) Gen(x, z int32, chunk level.Chunk) error {
	for x := byte(0); x < 16; x++ {
		for z := byte(0); z < 16; z++ {
			for y := byte(0); y < 60; y++ {
				chunk.SetBlock(x, y, z, 3)
			}
			chunk.SetBlock(x, 61, z, 2)
		}
	}
	chunk.PopulateHeight()
	return nil
}
