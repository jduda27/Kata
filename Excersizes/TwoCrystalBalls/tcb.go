package tcb

import (
	"math"
)

func LowestBreakingPoint(building []bool) (result int) {
	lastSafeFloor := 0
	crystalBalls := 2
	for i := 0; i < len(building); {

		if building[i] && crystalBalls < 2 {
			return i + 1
		}

		if building[i] {
			crystalBalls -= 1
			if lastSafeFloor == 0 {
				return i + 1
			} else {
				i = lastSafeFloor
			}
		}

		lastSafeFloor = i
		if crystalBalls == 1 {
			i += 1
		} else {
			i = i + int(math.Sqrt(float64(len(building))))
		}
	}
	return 0
}
