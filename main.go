package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//第一节课
//lv1:实现一个函数
//● 输入一个字符串，输出这个字符串的倒叙
//● 例如：输入"hello"，输出"olleh"
//func main() {
//	var str1 = [5]string{"H", "e", "l", "l", "o"}
//	fmt.Println(str1[4], str1[3], str1[2], str1[1], str1[0])
//}
//lv2:自行实现一个计算器，要求实现以下功能
//● 指令格式为a cmd b，（a，b为待运算数，只用考虑整数情况，cmd为运算符，只考虑加法减法乘法，例如1 + 2, 5 * 9, ...)
//● 可以只运行一次程序多次执行计算指令
//● 保存每次的运算结果并输出
//● 变量名符合命名规范，代码可读性良好

//func Add(number1, number2 int) int {
//	return number1 + number2
//}
//func minus(number1, number2 int) int {
//	return number1 - number2
//}
//func mcl(number1, number2 int) int {
//	return number1 * number2
//}
//func main() {
//	var number1, number2 int
//	var ans int
//	var cmd string
//	var history []int
//	history = make([]int, 0)
//	for {
//		fmt.Scanln(&number1, &cmd, &number2)
//		switch cmd {
//		case "+":
//			ans = Add(number1, number2)
//		case "-":
//			ans = minus(number1, number2)
//		case "*":
//			ans = mcl(number1, number2)
//		}
//		fmt.Println(ans)
//		history = append(history, ans)
//		for _, i := range history {
//			fmt.Print(i, " ")
//		}
//		fmt.Println()
//	}
//}

//lv3:无论学习什么知识都需要良好的信息搜集能力和自学能力，请你通过搜索引擎搜集相关资料，写程序实现以下功能
//● 声明一个切片s，调用go语言提供的随机数函数，生成100个随机数并append进s，
//● 调用go语言提供的排序函数，对切片s进行从大到小的排序
//● 输出有序之后的s切片
func rank(num int) []int {
	//冒泡排序
	var a []int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				b := a[j]
				a[j] = a[j+1]
				a[j+1] = b
			}
		}
	}
	return a
}

func main() {
	var s []int
	for i := 1; i <= 100; i++ {
		rand.Seed(time.Now().UnixNano())
		r1 := rand.Intn(100)
		s = append(s, r1)
	}
	sort.Ints(s)
	fmt.Println(s)
}
