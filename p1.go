package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(x, y int) int {
	k := 0
	k = x + y
	return k
}

func make_map(x int) {
	m := make(map[string]int)
	for i := 0; i <= x; i++ {
		z := strconv.Itoa(i)
		m[z] = i
	}
	fmt.Println(m)
}

func make_aray(x int) {

}

func make_slice(x int) {
	slice := make([]int, x+1)
	for i := 0; i <= x; i++ {
		slice[i] = i
	}
	fmt.Println(slice)

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
	make_slice(z)
	make_map(z)
}
