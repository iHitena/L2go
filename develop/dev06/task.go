package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.



site

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN,
разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем


*/

func main() {

	fieldsKey := flag.String("f", "", "key: fields")
	delimiterKey := flag.String("d", "\t", "key: delimiter")
	separatedKey := flag.Bool("s", false, "key: separated")

	//некое хранилище строк
	sliceString := []string{}

	// сканируем stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		scanStr := scanner.Text()
		// обрабатываем ключи строке из сканера
		flag.CommandLine.Parse([]string{scanStr})

		// если ключей нет, то добавляем строку в хранилище
		if !strings.Contains(scanStr, "-f") && !strings.Contains(scanStr, "-d") && !strings.Contains(scanStr, "-s") {
			sliceString = append(sliceString, scanStr)
			fmt.Println("Added!")
		}

		// если ключ обработки не пуст, то работает согласно инструкции
		if *fieldsKey != "" {

			// считывает со строки значение, которые необходимо вывести
			numericFiled := numericsHandler(*fieldsKey)

			// выводи значения учитывая ключи
			printStrings(sliceString, numericFiled, *delimiterKey, *separatedKey)

			// обнуляет ключ вывода
			*fieldsKey = ""
		}
	}

}

// обрабатывает строку и возвращает набор значений
func numericsHandler(numbersStr string) []int {
	var result []int

	numberSliceStr := strings.Split(numbersStr, ",")
	// добавляет в массив числа, разделенные запятом - 0,3 -> 0,3
	for _, numberStr := range numberSliceStr {

		number, err := strconv.Atoi(numberStr)

		if err == nil {
			result = append(result, number)
		} else {
			// возвращает slice с добавленными новыми интервалами  0-2 -> 0,1,2
			result = intervalHandler(numberStr, result)
		}

	}

	return result
}

// обрабатывает значения если были интервалы чисел  0-2 -> 0,1,2
func intervalHandler(interval string, result []int) []int {

	sliceIntervalStr := strings.Split(interval, "-")

	for i := 0; i < len(sliceIntervalStr)/2; i += 2 {

		// приводит старt и конец интервала к числам
		startInterval, errStart := strconv.Atoi(sliceIntervalStr[i])
		endInterval, errEnd := strconv.Atoi(sliceIntervalStr[i+1])

		if errStart != nil || errEnd != nil {
			continue
		}
		// добавляет цифры интервала к срезу
		for i := startInterval; i <= endInterval; i++ {
			result = append(result, i)
		}
	}

	return result
}

// выводит строку учитвая ключи
func printStrings(strs []string, numericField []int, delimiterKey string, separaredKey bool) {

	// столбики
	fmt.Println(strs, numericField, []rune(delimiterKey), separaredKey)

	// построчный вывод
	for _, str := range strs {
		// пропускаем итерацию если ключ требующий наличие строки с раздилителем активен, а  подстроки с ключом нет
		if separaredKey && !strings.Contains(str, delimiterKey) {
			continue
		}

		//разделяем строку в слайс специальным ключом
		sliceStr := strings.Split(str, delimiterKey)

		//выводим строку без ключа
		for _, numberField := range numericField {
			if numberField < len(sliceStr) {
				fmt.Print(sliceStr[numberField], " ")
			}
		}

		fmt.Print("\n")
	}
}
