package tcb

import (
	"fmt"
	"math"
)

func LowestBreakingPoint(building []bool) (result int) {
	lastSafeFloor := 0
	crystalBalls := 2
	for i := 0; i < len(building); {

		if building[i] && crystalBalls < 2 {
			fmt.Println(i)
			return i
		}

		if building[i] {
			crystalBalls -= 1
			if lastSafeFloor == 0 {
				fmt.Println(i)
				return i
			} else {
				i = lastSafeFloor
			}
		}

		lastSafeFloor = i
		if crystalBalls == 1 {
			i += 1
		} else {
			i = int(math.Sqrt(float64(len(building))))
		}
	}
	return
}
