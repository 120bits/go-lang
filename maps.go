package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k1"] = 12;
	m["k2"] = 11;

	fmt.Println("map :", m)
	fmt.Println("Length of Map:", len(m))
	
	val1 := m["k1"]
	fmt.Println("K1 Value: ", val1)
	
	val2 := m["k2"]
	fmt.Println("K2 Value: ", val2)
	
	delete(m, "K2")
	fmt.Println("After deleting K2 ", m)
	
	_, prs := m["K2"]
	fmt.Println("Does K2 exists", prs)
}
