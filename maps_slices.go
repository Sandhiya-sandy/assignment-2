package main

import "fmt"

func main() {
	a := map[string]int{"sandy": 92, "sai": 19, "alagappan": 50}

	fmt.Println(a["sai"])
	b := [5]int{1, 2, 3, 4, 5}
	slice_b := b[1:3]
	fmt.Println("slice_b:", slice_b)

}
