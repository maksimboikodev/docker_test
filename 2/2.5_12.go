/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов,
он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok",
иначе вывести "Wrong password"
Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok
*/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := false
	var str string
	fmt.Scan(&str)
	str2 := []rune(str)
	if len(str) < 5 {
		a = false
	} else {
		a = true
		for _, i := range str2 {
			if !unicode.Is(unicode.Latin, i) && !unicode.Is(unicode.Digit, i) {
				a = false
			}
		}
	}
	if a {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}
