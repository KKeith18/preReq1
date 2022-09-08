package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func add(x, y int) int {
	k := 0
	k = x + y
	return k
}

// func build_arr(x int) {
// 	k := x
// 	var arr [k]int
// 	for i := 0; i <= x; i++ {
// 		arr[i] = rand.Int()
// 	}
// 	sort.Ints(arr)
// 	return arr
//}

func build_slice(x int) {
	slice := make([]int, x)
	for i := 0; i <= x-1; i++ {
		slice[i] = rand.Int()
		//slice[i] = 1
	}
	fmt.Println(slice)
	//sorted = sort.Ints(slice)
	//return sorted
}
func main() {

	firstArg := os.Args[1] //this line copied from go documenation
	fmt.Println(firstArg)
	z, err := strconv.Atoi(firstArg)
	w := add(z, z)
	fmt.Println(w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(z)
	//k := build_arr(z)
	build_slice(z)
	//fmt.Println("Sorted array of random numbers:", k)
	//fmt.Println("Sorted slice of random numbers:", )
	fmt.Printf("z: %T\n", z)

}
