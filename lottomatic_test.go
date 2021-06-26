package main

import "testing"

func TestGenerate(t *testing.T) {
	t.Parallel()

	low := 1
	mPool := 35
	mNum := 7
	sPool := 20
	sNum := 1

	for i := 0; i < 100; i++ {
		w := generate(mPool, mNum, sPool, sNum)
		lenNums := len(w.numbers)
		lenSups := len(w.supplementaries)

		if lenNums != mNum {
			t.Errorf("main pool contains %v numbers expected %v", lenNums, mNum)
		}
		if lenSups != sNum {
			t.Errorf("supplementary pool contains %v numbers expected %v", lenSups, sNum)
		}
		for _, v := range w.numbers {
			if v <= 0 {
				t.Errorf("main pool: %v is lower than %v", v, low)
			}
			if v > mPool {
				t.Errorf("main pool: %v is higher than %v", v, mPool)
			}
		}
		for _, v := range w.supplementaries {
			if v <= 0 {
				t.Errorf("supplementary pool: %v is lower than %v", v, low)
			}
			if v > sPool {
				t.Errorf("supplementary pool: %v is higher than %v", v, mPool)
			}
		}
	}
}
