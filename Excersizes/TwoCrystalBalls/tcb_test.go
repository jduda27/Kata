package tcb

import (
	"fmt"
	"testing"
)

func TestTCB(t *testing.T) {
	fmt.Println("Testing tcb calculation")
	var building []bool = []bool{false, false, false, false, true, true, true}
	result := LowestBreakingPoint(building)
	fmt.Println(result)
}
