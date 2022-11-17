/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

Sample Input:
9119
Sample Output:
811181
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	SqrStrVal(str)
}
func SqrStrVal(str string) {
	for _, val := range str {
		intVal, _ := strconv.Atoi(string(val))
		Sqr := intVal * intVal
		fmt.Print(Sqr)
	}
}
