package learnslice

import "fmt"

// --------------- 切片的赋值拷贝 ---------------
/*
	两个切片共享底层数组，改变一个切片会影响另一个切片
 */

func PrintSlice03() {
	a1 := [...]int{1, 4, 2, 4, 6, 2, 23, 4, 5}
	s1 := a1[3:6]
	s2 := s1[1:5]   //s1 和 s2 都指向了同一个底层数组
	fmt.Println(a1, s1, s2)

	a1[4] = 10 //改变底层数组的值，会改变切片的值
	fmt.Println(a1, s1, s2)

	s1[2] = 39 //改变底层切片的值，会改变新切片和底层数组的值
	fmt.Println(a1, s1, s2)

	s2[0] = 67 //改变新切片的值，会改变底层切片和底层数组的值
	fmt.Println(a1, s1, s2)

}
