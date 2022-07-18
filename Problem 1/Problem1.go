package main

import (
	"fmt"
	"sort"
)

func main() {
	angka := [...]int{
		1,
		4,
		5,
		6,
		8,
		2,
	}

	fmt.Println("Random array", "\n")
	for i := 0; i <= 5; i++ {
		fmt.Print(angka[i])

		for j := 0; j < angka[i]; j++ {
			fmt.Print(" |")
		}

		fmt.Println(" ")
	}

	fmt.Println("==================================", "\n", "Sorted array", "\n")

	// digunakan untuk sorting ascending slicearray 
	ascendingSlice := angka[:]
	sort.Sort(sort.IntSlice(ascendingSlice))

	for i := 0; i <= 5; i++ {
		fmt.Print(angka[i])

		for j := 0; j < angka[i]; j++ {
			fmt.Print(" |")
		}

		fmt.Println(" ")
	}

	fmt.Println("==================================", "\n", "Reverse Sorted array", "\n")

	// digunakan untuk sorting decending slicearray
	decendingSlice := angka[:]
	sort.Sort(sort.Reverse(sort.IntSlice(decendingSlice)))

	for i := 0; i <= 5; i++ {
		fmt.Print(angka[i])

		for j := 0; j < angka[i]; j++ {
			fmt.Print(" |")
		}

		fmt.Println(" ")
	}

}
