package main

import "fmt"

func SumInts(ints map[string]int64) int64 {
	var s int64
	for _, value := range ints {
		s += value
	}
	return s
}

func SumFloats(fls map[string]float64) float64 {
	var s float64
	for _, value := range fls {
		s += value
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{"first": 30.00, "second": 50.00}
	fls := map[string]float64{"first": 25.25, "second": 35.35}
	//fmt.Printf("Non-generic sum: %s \n", SumInts(ints))
	//fmt.Printf("Non-generic sum: %s \n", SumFloats(fls))

	fmt.Printf("Generics Sum Ints: %s\n", SumIntsOrFloats[string, int64](ints))
	fmt.Printf("Generics Sum Floats: %s\n", SumIntsOrFloats[string, float64](fls))
}
