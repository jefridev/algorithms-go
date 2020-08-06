package main

import "testing"

func TestIfCountOfApplesAreOrangeAreOk(t *testing.T) {

	// Arrenge
	startpoint, endpoint := int32(7), int32(11) // house starting (s), endpoint (t)
	a, b := int32(5), int32(15)                 // apple tree location (a), orange tree location (b)
	var apples = []int32{-2, 2, 1}              // apples distances.
	var oranges = []int32{5, -6}                // oranges distances.

	//Assert
	resultedValues := countApplesAndOranges(startpoint, endpoint, a, b, apples, oranges)
	if resultedValues == nil {
		t.Error("Resulted values cannot be nil")
	} else {
		if resultedValues[0] != 1 || resultedValues[1] != 1 {
			t.Error("Returned values should 1,1 for this test.")
		}
	}
}
