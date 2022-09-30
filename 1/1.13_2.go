/*
Дано трехзначное число. Найдите сумму его цифр.
Формат входных данных
На вход дается трехзначное число.
Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:
745
Sample Output:
16
*/
package main

import "fmt"

func main() {
	var a, c int
	fmt.Scan(&a)
	for i := a; a != 0; i-- {
		b := a % 10
		a = a / 10
		c += b

	}
	fmt.Println(c)
}
