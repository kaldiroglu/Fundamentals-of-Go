package main

import "fmt"

func main() {
	//s := []int{3, 4}
	//for i, b := range s {
	//	print(i)
	//	print(b)
	//	s = append(s, 2)
	//}
	//fmt.Println(s)

	//var a [5]int = [...]int{1, 3, 5, 7, 9}
	//var s []int = a[1:3]
	//fmt.Println(s)
	//fmt.Printf("Length: %d and Capacity: %d\n", len(s), cap(s))
	//
	//a[0] = 10
	//a[1] = 20
	//a[2] = 30
	//a[3] = 40
	//a[4] = 50
	//fmt.Println(s)
	//
	//// take another slice as whole array
	//fmt.Println("Taking Slice As Whole Array")
	//s = a[:]
	//fmt.Println(s)
	//fmt.Printf("Length: %d and Capacity: %d\n", len(s), cap(s))
	//
	//fmt.Println("\nModifying Slice")
	//for i := 0; i < len(s); i++ {
	//	s[i] = 100 + i
	//}
	//fmt.Println(s)
	//fmt.Println(a)
	//
	//modifyArray(a)
	//fmt.Println(a)
	//
	//fmt.Println("\n**************")
	//modifySlice1(s)
	//fmt.Println(s)
	//fmt.Println(a)
	//
	//fmt.Println("\nCreating another slice for an array")
	//s1 := a[2:4]
	//fmt.Println(s1)
	//fmt.Printf("\nLength: %d and Capacity: %d\n", len(s1), cap(s1))
	//modifySlice2(s1)
	//fmt.Println(s1)
	//fmt.Println(s)
	//fmt.Println(a)
	//
	//fmt.Println("\nCreating a slice")
	//s2 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//fmt.Println(s2)
	//fmt.Printf("%c\n", s2)
	//fmt.Printf("Length: %d and Capacity: %d\n", len(s2), cap(s2))

}
func modifyArray(a [5]int) {
	fmt.Println("Modifying Array")
	for i := 0; i < len(a); i++ {
		a[i] = -i
	}
}

func modifySlice1(a []int) {
	fmt.Println("Modifying Slice")
	for i := 0; i < len(a); i++ {
		a[i] = -i
	}
}

func modifySlice2(a []int) {
	fmt.Println("Modifying Slice")
	for i := 0; i < len(a); i++ {
		a[i] = 100
	}
}
