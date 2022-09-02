package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func add(x, y int) int {
	k := 0
	k = x + y
	return k
}

func build_arr(x int) {
	k := x
	var arr [k]int
	for i := 0; i <= x; i++ {
		arr[i] = rand.Int()
	}
	sort.Ints(arr)
	return arr
}

func build_slice(x int) {
	slice := make([]int, x)
	for i := 0; i <= x; i++ {
		slice[i] = rand.Int()
	}
	sorted = sort.Ints(slice)
	return sorted
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
	k := build_arr(z)
	s := build_slice(z)
	fmt.Println("Sorted array of random numbers:", k)
	fmt.Println("Sorted slice of random numbers:", s)

}
