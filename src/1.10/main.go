package main

import "fmt"

func groupTemperature(in []float32) map[int][]float32 {
	m := make(map[int][]float32, 8)

	for _, num := range in {
		norm := int(num/10) * 10
		t, _ := m[int(norm)]
		t = append(t, num)
		m[int(norm)] = t
	}

	return m
}

func main() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for norm, num := range groupTemperature(temperature) {
		fmt.Println(norm, num)
	}
}
