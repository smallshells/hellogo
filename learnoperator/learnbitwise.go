package learnoperator

import "fmt"

//---------------- 位运算符 ---------------
/*
	位运算符对整数在内存中的二进制位进行操作
	&  参与运算的两个数各位对应的二进位相与（两位均为 1 才为 1 ）
	|  参与运算的两个数各位对应的二进位相或（两位有一位为 1 就为 1 ）
	^  参与运算的两个数各位对应的二进位相异或（两位不一样则为 1 ）
	<< 左移n位就是乘以2的你n次方（ a << b 就是把a的各二进位全部左移b位，高位丢弃，低位补0）
	>> 左移n位就是除以2的你n次方（ a >> b 就是把a的各二进位全部右移b位）
*/

func PrintBitwise() {
	// 5 的二进制表示： 101
	// 6 的二进制表示： 110

	fmt.Println("5 & 6 = ", 5&6)
	fmt.Println("5 | 6 = ", 5|6)
	fmt.Println("5 ^ 6 = ", 5^6)
	fmt.Println("5 << 3 = ", 5<<3)
	fmt.Println("5 >> 2 = ", 5>>2)

	var m = int8(1)
	fmt.Println(m << 10)  //由于m只有8位，左移10位越界了，结果为0
}