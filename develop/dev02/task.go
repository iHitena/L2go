package main

import (
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.


site

Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

*/

func main() {

	stringHandler("a4bc2d5e")
	stringHandler("abcd")
	stringHandler("45")
	stringHandler("")

	stringHandler(`qwe\4\5`)
	stringHandler(`qwe\45`)
	stringHandler(`qwe\\5`)

}

func stringHandler(str string) {
	// на случай пустой строки
	if str == "" {
		fmt.Println(str)
		return
	}

	//создаем переменные для результата - вывести результат
	/// предыдущего символа - будем его копировать
	//количество копий -
	// флаг на случай вводов впециальных символов
	result := ""
	preSymbol := ""
	countSymbol := ""
	flagSymbol := false

	// считываем руну ссымвола в строке с последующей его обработкой
	for _, runeStr := range str {
		// если руна оказалась специальным символом, то ставим флаг и переходим к след символу
		if runeStr == 92 && !flagSymbol {
			flagSymbol = true
			continue
		}
		// обработаем несколько случаев.
		switch {

		// 1) если у нас активный флаг и пустует предыдущий символ, который мы должны размножить. тогда запоминаем символ
		case flagSymbol && preSymbol == "":
			preSymbol = string(runeStr)

		case !flagSymbol && 48 <= runeStr && runeStr <= 57:

			//2) специального символа не было, значит руна ввиде цифры должна размножить предыдущий символ
			countSymbol = string(runeStr)
			// функция множит символ и возвращает строку, где каждый элемент это один и тот же символ
			result += addSymbolToStr(preSymbol, countSymbol)
			// обнуляем переменные, т.к. символ был размножен
			preSymbol, countSymbol, flagSymbol = "", "", false

		default:
			// 3) остались случаи когда у нас активный флаг - значит в предыдущий переменную предыдущего символа можно поместить что угодно.
			// или если руна пренадлежит обычному символу - тут флаг не учитывается

			result += preSymbol
			preSymbol = string(runeStr)
			flagSymbol = false
		}

	}
	//на случай не корректный данных
	if result != "" {
		fmt.Println(result + preSymbol)
	} else {
		fmt.Println("некорректная строка")
	}

}

// функция множит символ и возвращает строку, где каждый элемент это один и тот же символ
func addSymbolToStr(symbol string, countSting string) string {
	result := ""
	countInt, _ := strconv.Atoi(countSting)
	for i := 0; i < countInt; i++ {
		result += symbol
	}
	return result
}
