package algorithm

import (
	"testing"
)

func TestProbability_GetRand(t *testing.T) {
		proArr := make(map[string]int)
		proArr["1"] = 100
		proArr["2"] = 200
		proArr["3"] = 300
		proArr["4"] = 400

		am := NewProbability(proArr)

		// test ...
		sts := make(map[string]int)
		for i := 0; i < 10000000; i++ {
			sts[am.GetRand()]++
		}

		t.Log(sts)
}