package gen

import "github.com/L7-MCPE/lav7/types"

type SampleGenerator struct{}

func (s SampleGenerator) Gen(x, z int32) (chunk *types.Chunk) {
	chunk = new(types.Chunk)
	chunk.Init()
	chunk.Mutex().Lock()
	defer chunk.Mutex().Unlock()
	for x := byte(0); x < 16; x++ {
		for z := byte(0); z < 16; z++ {
			for y := byte(0); y < 60; y++ {
				chunk.SetBlock(x, y, z, 3)
			}
			chunk.SetBlock(x, 60, z, 2)
			chunk.SetBiomeColor(x, z, x*16, x*z, z*16)
		}
	}
	chunk.PopulateHeight()
	return
}
