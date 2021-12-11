package main

import "fmt"

func main() {

	var int_array [5]int;
	var str_array [5]string;

	
	int_array[0] = 10
	int_array[1] = 20
	int_array[2] = 30
	int_array[3] = 40
	int_array[4] = 50

	fmt.Println("This is the integer Array: ", int_array)

	str_array[0] = "first"
	str_array[1] = "second"
	str_array[2] = "third"
	str_array[3] = "fourth"

	fmt.Println("This is the string array: ", str_array)

} 
