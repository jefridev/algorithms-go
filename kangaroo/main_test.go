package main

import "testing"

func TestIfTheyMeetAtOnePoint(t *testing.T) {
	const expected = "YES"
	value := kangaroo(4523, 8092, 9419, 8076)
	if value != expected {
		t.Error("It should be return YES, because they meet at one point.")
	}
}

/*
	It should return a value.
	X1,V1 - X2,V2.
	4523,
*/
