package main

import "fmt"

/**********************************************************************
 * 1. single, var 키워드 사용 & 자료형 사용(초기값 옵션)
 * 2. single, var 키워드 사용 & 자료형 생략(초기값 필수)
 * 3. single, var 키워드 생략 & 자료형 생략(초기값 필수)
 * 4. multiple, var 키워드 사용 & 자료형 사용(초기값 옵션)
 * 5. multiple, var 키워드 사용 & 자료형 생략(초기값 필수)
 * 6. multiple, var 키워드 생략 & 자료형 생략(초기값 필수) 
 * 7. mulitple, var 키워드 사용 & 자료형 사용 | 생략
 **********************************************************************/

func main() {

	// 1. single, var 키워드 사용 & 자료형 사용(초기값 옵션)
	var sKeyYesTypeYes int
	sKeyYesTypeYes = 10
	fmt.Println(sKeyYesTypeYes)

	// 2. single, var 키워드 사용 & 자료형 생략(초기값 필수)
	var sKeyYesTypeNo = 20
	fmt.Println(sKeyYesTypeNo)

	// 3. single, var 키워드 생략 & 자료형 생략(초기값 필수)
	sKeyNoTypeYes := 30
	fmt.Println(sKeyNoTypeYes)

	// 4. multiple, var 키워드 사용 & 자료형 사용(초기값 옵션)
	var mKeyYesTypeYes1, mKeyYesTypeYes2 int
	mKeyYesTypeYes1, mKeyYesTypeYes2 = 40, 41
	fmt.Println(mKeyYesTypeYes1, mKeyYesTypeYes2)

	// 5. multiple, var 키워드 사용 & 자료형 생략(초기값 필수)
	var mKeyYesTypeNo1, mKeyYesTypeNo2 = 50, 51
	fmt.Println(mKeyYesTypeNo1, mKeyYesTypeNo2)

	// 6. multiple, var 키워드 생략 & 자료형 생략(초기값 필수)
	mKeyNoTypeNo1, mKeyNoTypeNo2 := 60, 61
	fmt.Println(mKeyNoTypeNo1, mKeyNoTypeNo2)

	// 7. mulitple, var 키워드 사용 & 자료형 사용 | 생략
	var (
		a, b int = 70, 71
		x, y = 80, "팔십일"
	)
	fmt.Println(a, b, x, y)
}
