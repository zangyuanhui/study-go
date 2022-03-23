/*
范型相关代码，go版本必须>1.18才可用，提高代码复用性、灵活性
*/
package main

import "fmt"

// SumInts 这将被godoc生成文档
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

//func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}

//type Number interface {
//	int64 | float64
//}

//func SumNumbers[K comparable, V Number](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}

func main() {
	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	floats := map[string]float64{
		"first":  3.3,
		"second": 4.4,
	}

	fmt.Printf("Non-Generic Sums:%v and %v", SumInts(ints), SumFloats(floats))
	//fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	//fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}
