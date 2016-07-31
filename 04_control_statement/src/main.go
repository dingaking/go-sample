package main

import "fmt"
import "io/ioutil"

/**********************************************************************
 * 1. if statement {}
 *   1.1 if와 같은 라인에 "{" 필요
 *   1.2 "{" 생략 불가.
 *   1.3 if 조건문에 함수 사용
 * 2. for 시작식; 종료조건; 변화식 {}
 *   2.1 single variable
 *   2.2 multiple variable
 * 3. switch variable { case value: code }
 *   3.1 break
 *   3.2 fallthrough
 *   3.3 multiple
 *   3.4 statement
 **********************************************************************/

func switch_fallthrough_func(i int) {

	switch i {
	case 0:
		fmt.Println("switch fallthrough: ", 0)
		fallthrough
	case 1:
		fmt.Println("switch fallthrough: ", 1)
		fallthrough
	case 2:
		fmt.Println("switch fallthrough: ", 2)
	case 3:
		fmt.Println("switch fallthrough: ", 3)
		fallthrough
	default:
		fmt.Println("switch fallthrough: ", "default")
	}
}

func switch_func(i int) {

	switch i {
	case 0:
		fmt.Println("switch : ", i)
	case 1:
		fmt.Println("switch : ", i)
	case 2:
		fmt.Println("switch : ", i)
	case 3:
		fmt.Println("switch : ", i)
	default:
		fmt.Println("switch : ", "default")
	}
}

func main() {

	// 1. if statement {}
	if 1 > 5 {
		fmt.Println("1은 5보다 크다.")
	} else if 1 >= 5 {
		fmt.Println("1은 5보다 크거나 같다.")
	} else {
		fmt.Println("1은 5보다 작거나 같다.")
	}

	/*
	// 1.1 if와 같은 라인에 "{" 필요
	if 1 > 5
	{
		fmt.Println("1은 5보다 크다.")
	}*/

	/*
	// 1.2 "{" 생략 불가.
	if 1 > 5
		fmt.Println("1은 5보다 크다.")
	*/

	// 1.3 if 조건문에 함수 사용
	if b, err := ioutil.ReadFile("./test.txt"); err == nil {
		fmt.Printf("%s", b)
	}

	// 2.1 single variable
	for i:=0; i<15; i++ {
		if i == 2 {
			continue
		}
		if i == 10 {
			break
		}
		fmt.Println("for statement : ", i)
	}

	// 2.2 multiple variable
	for i,j := 0, 10; i<10 && j>0; i, j = i + 1, j - 1 {
		fmt.Println("multiple i, j : ", i, ", ", j)
	}

	// 3. switch variable { case value: code }
	switch_func(0)
	switch_func(3)
	switch_func(5)

	// 3.1 break
	i := 0
	switch i {
	case 0:
		if i==0 {
			fmt.Println("switch break")
			break
		}
		fmt.Println("switch : ", i)
	case 1:
		fmt.Println("switch : ", i)
	case 2:
		fmt.Println("switch : ", i)
	case 3:
		fmt.Println("switch : ", i)
	default:
		fmt.Println("switch : ", "default")
	}

	// 3.2 fallthrough
	switch_fallthrough_func(1)

	// 3.3 multiple
	i = 2
	switch i {
	case 0,1:
		fmt.Println("switch multiple : ", i)
	case 2,3:
		fmt.Println("switch multiple : ", i)
	default:
		fmt.Println("switch multiple : ", "default")
	}

	// 3.4 statement
	i = 2
	switch {
	case 0 <= i && i < 2:
		fmt.Println("switch statement(0 <= i && i < 2) : ", i)
	case 2 <= i && i < 4:
		fmt.Println("switch statement(2 <= i && i < 4) : ", i)
	default:
		fmt.Println("switch statement(default) : ", "default")
	}

}

