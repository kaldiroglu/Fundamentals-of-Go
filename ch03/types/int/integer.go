package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var anInt1 int = 10
	var anInt2 uint = 10
	fmt.Printf("%d %d\n", anInt1, anInt2)

	//fmt.Println(anInt1 == anInt2)

	fmt.Printf("%T %T\n", anInt1, anInt2)

	size1 := unsafe.Sizeof(anInt1) // size in bytes
	size2 := unsafe.Sizeof(anInt2)
	fmt.Printf("%d %d\n", size1, size2)

	var anInt3 int32 = 7
	var anInt4 int64 = 12
	size3 := unsafe.Sizeof(anInt3)
	size4 := unsafe.Sizeof(anInt4)
	fmt.Printf("%d %d\n", size3, size4)

	//anInt5 := anInt3 + anInt4
	//fmt.Printf("%d\n", anInt5)

	//anInt5 := anInt1 + anInt4
	//fmt.Printf("%d\n", anInt5)

	//	f1(anInt1)
	//	f1(anInt2)
	//	f1(anInt3)
	f1(anInt4)

	f2(anInt1)
	//	f2(anInt2)
	//	f2(anInt3)
	//	f2(anInt4)

	fmt.Println("\nConversions")
	var i int = anInt1
	//var j int = anInt2
	var j int = int(anInt2)

	fmt.Printf("%d %d\n", i, j)

	/*fmt.Println("\nOperators")
	u := 137
	v := 98
	fmt.Println(u ^ v)
	fmt.Println(^v)*/

	operations()
}

func f1(i int64) {
	fmt.Println(i)
}

func f2(i int) {
	fmt.Println(i)
}

// From Go PL book p. 54
func operations() {
	fmt.Println("\nOperators")
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
