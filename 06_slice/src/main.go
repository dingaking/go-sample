package main

import "fmt"

/**********************************************************************
 * 1. var O, type O
 * 2. var O, type X
 * 3. var X, type X
 * 4. initialization
 * 5. append
 * 6. reference
 * 7. copy
 **********************************************************************/

func main() {

	// 1. var O, type O 
	/*
	// 초기화 또는 assign 없이 값 설정시 에러 발생
	var a []int;
	a[1] = 1
	fmt.Println(a)
	*/
	var b []int
	b = make([]int, 5)
	b[1] = 1
	fmt.Println(b)
	
	// 2. var O, type X
	/*
	// 선언시 슬라이스에 대한 명시 필요
    var c
    c = make([]int, 5)
    c[1] = 2
	fmt.Println(c)
	*/
	var d = make([]int, 5)
	d[1] = 2
	fmt.Println(d)

	// 3. var X, type X
	e := make([]int, 5)
	e[1] = 3
	fmt.Println(e)

	// 4. initialization
	f := []int{1,2,3,4,5}
	fmt.Println(f)
	g := []int {
		6,
		7,
		8,
		9,
		10,
	}
	fmt.Println(g)
	h := make([]int, 5, 10)
	fmt.Println(h)
	fmt.Println(len(h))
	fmt.Println(cap(h))
	/*
	// over index
	fmt.Println(h[5])
	*/

	// 5. append
	i := []int{1,2,3}
	i = append(i, 4, 5)
	/*
    // i에 할당하지 않으면 컴파일 에러 발생
	append(i, 4, 5)
    */
	fmt.Println(i)

	// 6. reference(difference from array)
	fmt.Println("\n6. reference(difference from array)")
	j_array_1 := [3]int {1, 2, 3}
	var j_array_2 [3]int
	fmt.Println(j_array_1)
	fmt.Println(j_array_2)
	j_array_2 = j_array_1
	fmt.Println(j_array_1)
	fmt.Println(j_array_2)
	j_array_2[2] = 9
	fmt.Println(j_array_1)
	fmt.Println(j_array_2)

	fmt.Println("")

	j_slice_1 := []int {1, 2, 3}
	var j_slice_2 []int
	fmt.Println(j_slice_1)
	fmt.Println(j_slice_2)
	j_slice_2 = j_slice_1
	fmt.Println(j_slice_1)
	fmt.Println(j_slice_2)
	j_slice_2[2] = 9
	fmt.Println(j_slice_1)
	fmt.Println(j_slice_2)
}
