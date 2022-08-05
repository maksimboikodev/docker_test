/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
Формат входных данных
На вход дается натуральное число, не превосходящее 10000.
Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:
1234
Sample Output:
1
*/
package main

import "fmt"

func main() {
	var a, val int
	fmt.Scan(&a)
	switch {
	case a >= 10000:
		val = 4
	case a >= 1000:
		val = 3
	case a >= 100:
		val = 2
	case a >= 10:
		val = 1
	}
	for i := val; i != 0; i-- {
		a = a / 10
	}
	fmt.Println(a)
}
