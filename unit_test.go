package main

import (
  "fmt"
  "testing"
)

func TestCalculateTransaction1 (t *testing.T){ // test transaction with $120.00
	var test int64 = 120
	points := calculateTransaction(test)
	fmt.Printf("in test, rewarded client %d  points \n", points )
	if points != 90 {
		t.Errorf("calculateTransaction(120.00) = %d; want 90", points) 
	}
}

/* func TestFormatCheck1 (t *testing.T){
	bool := formatCheck(120.00)
	if bool != true {
		t.Errorf("formatCheck(120.00) = %t; want true", bool) 
	}
} */