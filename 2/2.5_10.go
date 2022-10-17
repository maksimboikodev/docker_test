/*
 На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot
Sample Output:
hello
*/
package main

import "fmt"

func main() {
	var str, b string
	fmt.Scan(&str)
	a := []rune(str)
	for i := range a {
		if i%2 != 0 {
			b += string(a[i])
		}
	}
	fmt.Print(string(b))
}
