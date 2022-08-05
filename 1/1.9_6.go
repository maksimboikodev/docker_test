/*По данному трехзначному числу определите, все ли его цифры различны.
Формат входных данных
На вход подается одно натуральное трехзначное число.
Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO

*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	vol1 := a / 100
	vol2 := a / 10 % 10
	vol3 := a % 10
	if vol1 == vol2 || vol1 == vol3 || vol2 == vol3 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
