/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:
24h0m0s
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()

	splited := strings.Split(text, ",")
	firstDate := splited[0]
	secondDate := splited[1]
	// 13.03.2018 14:00:15
	layout := "02.01.2006 15:04:05"
	parsedFirst, _ := time.Parse(layout, firstDate)
	parsedSecond, _ := time.Parse(layout, secondDate)
	difference := parsedFirst.Sub(parsedSecond)
	if difference.Hours() < 0 {
		difference = parsedSecond.Sub(parsedFirst)
		fmt.Printf("%v", difference)
	} else {
		fmt.Printf("%v", difference)
	}
}
