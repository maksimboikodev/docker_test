/*
Из натурального числа удалить заданную цифру.
Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.

Sample Input:
38012732
3
Sample Output:
801272
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, h int
	fmt.Scan(&a, &b)
	d := a
	for i := a; a != 0; i-- {
		a = a / 10
		c++
	}
	for i := c; i != 0; i-- {
		k := math.Pow(10, float64(i-1))
		h = d / int(k)
		d = d % int(k)
		if h == b {
			continue
		}
		fmt.Print(h)
	}
}
