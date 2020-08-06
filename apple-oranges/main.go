package main

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) []int32 {
	applesCount := int32(0)
	for i := 0; i < len(apples); i++ {
		currentLocation := a + apples[i]
		if currentLocation >= s && currentLocation <= t {
			applesCount++
		}
	}
	orangeCount := int32(0)
	for i := 0; i < len(oranges); i++ {
		currentLocation := b + oranges[i]
		if currentLocation >= s && currentLocation <= t {
			orangeCount++
		}
	}
	return []int32{applesCount, orangeCount}
}
