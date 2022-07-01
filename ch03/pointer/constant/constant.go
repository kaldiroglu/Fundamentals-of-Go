package main

import "fmt"

func main() {

	const const1 int = 7
	const const2 float64 = 0.18

	//const1++
	//const1 = 7

	const unTypedConst1 = 3
	const unTypedConst2 = 3.14

	var int1 int = unTypedConst1
	var int2 int32 = unTypedConst1
	var int3 int64 = unTypedConst1
	var int4 uint32 = unTypedConst1

	fmt.Printf("%d %d %d %d\n", int1, int2, int3, int4)

	var fp1 float32 = unTypedConst1
	var fp2 float32 = unTypedConst1
	fmt.Printf("%f %f\n", fp1, fp2)

	//int1 = int(unTypedConst2)
	//var int5 int64 = int64(unTypedConst2)
	//fmt.Printf("%f %f\n", int1, int5)
}
