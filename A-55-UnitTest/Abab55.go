package main

import "math"

// Kubus struct
type Kubus struct {
	Sisi float64
}

// Volume kubus
func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

// Luas kubus
func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

// Keliling Kubus
func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}

