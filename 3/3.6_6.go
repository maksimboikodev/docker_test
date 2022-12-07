/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            // ...
        }
    ]
}
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
{
    "Average": 14.1
}
Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).
Если у вас возникли проблемы с чтением / записью данных, то этот комментарий для вас: в уроках об интерфейсах и работе с файлами мы рассказывали, что стандартный ввод / вывод - это файлы, к которым можно обратиться через os.Stdin и os.Stdout соответственно, они удовлетворяют интерфейсам io.Reader и io.Writer, из них можно читать, в них можно писать.
Один из способов чтения, использовать ioutil.ReadAll:

data, err := ioutil.ReadAll(os.Stdin)
if err != nil {
    // ...
}

// data - тип []byte
Задачу можно выполнить и другими способами, в частности использовать bufio. Буквально в следующем шаге (через один, на самом деле) будет рассказано о еще одном способе чтения / записи, можете забежать немного вперед, а затем вернуться к задаче.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Group - группа студентов
type Group struct {
	// ID       int
	// Number   string
	// Year     int
	Students []struct {
		// LastName   string
		// FirstName  string
		// MiddleName string
		// Birthday   string
		// Address    string
		// Phone      string
		Rating []int
	}
}

type averRating struct {
	Average float64
}

func main() {
	// file, _ := os.Open("./text.json")
	input, _ := ioutil.ReadAll(os.Stdin)
	// fmt.Println(input)

	newGroup := Group{}
	err := json.Unmarshal(input, &newGroup)
	if err != nil {
		fmt.Println(err)
	}

	students := 0
	marks := 0
	for i := 0; i < len(newGroup.Students); i++ {
		// fmt.Println(newGroup.Students[i])
		students++
		for m := 0; m < len(newGroup.Students[i].Rating); m++ {
			marks++
			// fmt.Println(newGroup.Students[i].Rating[m])
		}
	}
	newAvg := averRating{float64(marks) / float64(students)}
	avg, _ := json.MarshalIndent(newAvg, "", "    ")
	os.Stdout.Write(avg)
	// fmt.Printf("Students: %v\n Marks: %v", students, marks)
}
