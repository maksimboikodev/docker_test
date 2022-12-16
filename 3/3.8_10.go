/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)
Функция должна называться removeDuplicates()
Выводить или вводить ничего не нужно!
*/
func removeDuplicates(inputStream, outputStream chan string) {
var previousValue string
for v := range inputStream {
if previousValue != v {
outputStream <- v
previousValue = v
}
}
close(outputStream)
}
