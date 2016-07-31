package main

import "fmt"

/**********************************************************************
 * 1. var exist
 * 2. var non exist
 * 3. initialization
 * 4. multidimensional
 * 5. iteration
 *   5.1 for-len
 *   5.2 for-range
 **********************************************************************/

func main() {

	// 1. var exist
	var a [5]string;
	a[1] = "z"
	fmt.Println(a)
	fmt.Println(a[1])

	// 2. var non exist
	b := [5]string{"a","b","c","d","e"}
	fmt.Println(b)
	fmt.Println(b[1])

	// 3. initialization
	c := [5]string {
		"z",
		"y",
		"x",
		"w",
		"v",
	}
	fmt.Println(c)
	fmt.Println(c[1])

	// 4. multidimensional
	d := [2][2]int {
		{1,2},
		{3,4},
	}
	fmt.Println(d)
	fmt.Println(d[1][0])

	// 5.1 for-len
	e := [5]int {1,2,3,4,5}
	for i := 0; i < len(e); i++ {
		fmt.Println(e[i])
	}

	// 5.2 for-range
	f := [5]int{6,7,8,9,10}
	for i := range f {
		fmt.Println(i, f[i])
	}
	for i, val := range f {
		fmt.Println(i, val)
	}
	for _, val := range f {
		fmt.Println(val)
	}
	
}
