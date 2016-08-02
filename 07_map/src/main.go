package main

import "fmt"

/**********************************************************************
 * 1. Declaration
 * 2. Interation
 * 3. Delete
 * 4. Multidimensional
 **********************************************************************/

func main() {

	// 1. Declaration
	fmt.Println("\n1. Declaration")
	var a map[string]int = make(map[string]int)
	fmt.Println(a)
	a["key1"] = 1
	fmt.Println(a)

	/*
	// error : when use before allocation
	var b map[string]int
	b["key1"] = 1
	fmt.Println(b)
	*/

	c := map[string]int{"key1":2}
	fmt.Println(c)

	d := map[string]int{"key1":1,
		"key2":2,}
	fmt.Println(d)

	// 2. Interation
	fmt.Println("\n2. Interation")
	e := map[string]int {"key1":1, "key2":2, "key3":3}
	for k, v := range e {
		fmt.Println(k, v)
	}
	for _, v := range e {
		fmt.Println(v)
	}
	for k, _ := range e {
		fmt.Println(k)
	}

	// 3. Delete
	fmt.Println("\n3. Delete")
	l := map[string]int {"key1":1, "key2":2, "key3":3}
	fmt.Println(l)
	delete(l, "key1")
	fmt.Println(l)

	// 4. Multidimensional
	fmt.Println("\n4. Multidimensional")
	m := map[string]map[string]int {
		"key1":map[string]int {"key11":1},
		"key2":map[string]int {"key21":2},
		"key3":map[string]int {"key31":3}}
	fmt.Println(m)
	fmt.Println(m["key1"])
	fmt.Println(m["key2"]["key21"])
	
}
