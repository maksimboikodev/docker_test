/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
Входные данные
Вводится четыре числа.
Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел
Sample Input:
4 5 6 7
Sample Output:
4
*/
package main

import "fmt"

func minimumFromFour() int {
	var n, k int
	k = 99999
	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		// m=n
		if n <= k {
			k = n
		}
	}
	return k
}
