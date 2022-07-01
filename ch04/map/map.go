package main

import "fmt"

func main() {
	//map1()
	//map2()
	//iterateMap()
	map3()
}

func map1() {
	m := make(map[int]string)
	fmt.Println(m)

	m[1] = "Adana"
	m[6] = "Ankara"
	m[10] = "Balıkesir"
	m[16] = "Bursa"
	m[30] = "Hakkari"
	m[34] = "İstanbul"
	m[34] = "İstanbul"
	m[35] = "İzmir"
	m[61] = "Trabzon"
	m[81] = "Düzce"

	fmt.Println(m)

	fmt.Println(m[1])
	fmt.Println(m[61])

	s := m[10]
	fmt.Println("City with plate # 10 is", s)

	s = m[24]
	fmt.Println("City with plate # 24 is", s)

	delete(m, 10)
	s = m[10]
	fmt.Println("City with plate # 10 is", s)

	m[81] += " Dümdüz"
	fmt.Println(m[81])
	fmt.Println(m)

	// Replace an existing value
	m[35] = "İzmirX"
	fmt.Println(m)
}

func map2() {
	m1 := map[int]string{
		1: "Adana",
		6: "Ankara",
		//6:  "XXX",
		10: "Balıkesir",
		16: "Bursa",
	}

	fmt.Println(m1)

	var m2 map[string]string
	fmt.Println(m2)

	//m2["greeting"] = "Selam" // Panic!

	m2 = map[string]string{
		"greeting": "Selam",
		"goodbye":  "Hoşçakal",
	}
	fmt.Println(m2)
	s1, ok1 := m2["goodbye"]
	if ok1 {
		fmt.Println(s1)
	} else {
		fmt.Println("No such entry exist!")
	}

	if _, ok2 := m2["goodbyex"]; !ok2 {
		fmt.Println("No such entry exist!")
	}
}

func iterateMap() {
	m := make(map[int]string)
	fmt.Println(m)

	m[1] = "Adana"
	m[6] = "Ankara"
	m[10] = "Balıkesir"
	m[16] = "Bursa"
	m[30] = "Hakkari"
	m[34] = "İstanbul"
	m[34] = "İstanbul"
	m[35] = "İzmir"
	m[61] = "Trabzon"
	m[81] = "Düzce"

	fmt.Println("\nIterating Through Keys")
	for plate := range m {
		fmt.Printf("%d\n", plate)
	}

	fmt.Println("Iterating Through Values")
	for _, city := range m {
		fmt.Printf("%s\n", city)
	}

	fmt.Println("Iterating Through Map")
	for plate, city := range m {
		fmt.Printf("%d\t%s\n", plate, city)
	}
}

func map3() {
	m := make(map[int]string)
	fmt.Println(m)

	m[1] = "Adana"
	m[6] = "Ankara"
	m[10] = "Balıkesir"
	m[16] = "Bursa"
	m[30] = "Hakkari"
	m[34] = "İstanbul"
	m[34] = "İstanbul"
	m[35] = "İzmir"
	m[61] = "Trabzon"
	m[81] = "Düzce"

	plates := make([]int, 0, len(m))
	fmt.Println(plates)
}
