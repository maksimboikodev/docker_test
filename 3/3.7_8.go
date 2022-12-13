/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).
Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:
12 мин. 13 сек.
Sample Output:
Fri May 15 19:28:18 UTC 2020
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	splitted := strings.Split(text, ".")
	minutes := strings.TrimSpace(splitted[0])[:2]
	seconds := strings.TrimSpace(splitted[1])[:2]
	converted := fmt.Sprintf("%vm%vs", minutes, seconds)
	// fmt.Printf("%v\n minutes: %v, seconds: %v, converted: %v\n", splitted, minutes, seconds, converted)
	duration, _ := time.ParseDuration(converted)

	defaultTime := time.Unix(now, 0)
	fmt.Printf("%v\n", defaultTime.UTC().Add(duration).Format(time.UnixDate))
}
