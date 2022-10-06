/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)
Результат: 2, 1
Sample Input:
Sample Output:
*/
package main

func sumInt(a ...int) (int, int) {
	total := 0
	c := 0
	for _, v := range a {
		c++
		total += v
	}
	return c, total
}
