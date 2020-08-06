package main

import "testing"

func TestCandlesShouldBeTurnOff(t *testing.T) {
	candlesHeights := []int32{4, 4, 4, 1, 2, 2, 2}
	turnOffCandles := birthdayCakeCandles(candlesHeights)
	t.Log(turnOffCandles)
	if turnOffCandles != int32(5) {
		t.Error("There should be only 3 candles turned off")
	}
}
