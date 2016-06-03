package main

import (
	"math/rand"
	"fmt"
)

func main() {
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

	fmt.Println(sts)
}

type Probability struct {
	proArr map[string]int
	proSum int
}

func NewProbability(proArr map[string]int) *Probability {
	return &Probability{
		proArr: proArr,
	}
}

func (r *Probability) sum() {
	r.proSum = 0
	for _, v := range r.proArr {
		r.proSum += v
	}
}

func (r *Probability) GetRand() (key string) {
	r.sum()

	for k, v := range r.proArr {
		rNum := rand.Intn(r.proSum) + 1
		if rNum <= v {
			key = k
			break
		} else {
			r.proSum -= v
		}
	}
	return
}