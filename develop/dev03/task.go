package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.


site

Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
 на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

*/

func main() {

	k := flag.String("k", "1234", "specifying a column to sort")
	n := flag.Bool("n", false, "sort by numeric value")
	r := flag.Bool("r", false, "sort in reverse order")
	u := flag.Bool("u", false, "do not output duplicate lines")
	flag.Parse()

	// читаем файл и данные преобразуем в сроковый слайс
	dataFile := readingFile("in.txt")

	//если флаг дублей активен, убираем из слайка дубли
	if *u {
		dataFile = deletDublicate(dataFile)
	}

	switch {

	case *k != "1234": // кейс обрабатывает случай, если колонка была указана. так же внутри кейса проверяются и другие ключи
		dataFile = specifyingColumnSort(dataFile, *k, !*r)

	case *r: // кейс обрабатывает случай если нужна только сортировка в обратном порядке
		dataFile = sortStrings(dataFile, !*r)

	case *n: // стандартная сортировка

		dataFile = sortStrings(dataFile, *n)
	}

	recordFile("out.txt", dataFile)

}

func specifyingColumnSort(sliceStrings []string, column string, normalSort bool) []string {

	for i := 0; i < len(sliceStrings); i++ {
		// функция передает строки в специальный обработчик , внутри которого происходит обработка колонок
		sliceStrings[i] = stringHeandler(sliceStrings[i], column, normalSort)
	}

	return sliceStrings
}

// возвращает строку где первое слово - слово солонки, после идут отсортированные колонки
func stringHeandler(str string, column string, normalSort bool) string {
	sliceResult := []string{""}
	sliceStr := strings.Split(str, " ")

	for _, word := range sliceStr {
		// колонки формуруются на основании переданного слова (column). каждая колонка - отдельный элемент слайса
		if word == column {
			sliceResult = append(sliceResult, "")
			continue
		}

		sliceResult[len(sliceResult)-1] += word + " "

	}

	str = ""
	//  если елемент всего 1 - значит доп колонки не создавались - можно вернуть страку обработав её стандартной сортировкой
	if len(sliceResult) == 1 {
		return sortString(sliceResult[0], normalSort)
	}
	// каждый элемент среза - колонка. каждая колонка обрабатывается сортировкой отдельно. после формируется строка
	for i := 0; i < len(sliceResult); i++ {
		str = strings.TrimSpace(str)
		str += " " + column + " " + sortString(sliceResult[i], normalSort)

	}

	return strings.TrimSpace(str)
}

//сотрирует строки
func sortStrings(sliceStrings []string, normalSort bool) []string {
	for i := 0; i < len(sliceStrings); i++ {
		sliceStrings[i] = sortString(sliceStrings[i], normalSort)
	}
	return sliceStrings
}

// сортирует строку. возрастает если  булевого значения верно, в противном случае убывает
func sortString(str string, normalSort bool) string {
	str = strings.TrimSpace(str)
	sliceStr := strings.Split(str, " ")
	// сотрировка зависит от флага . она может быть убывающей или возрастающей
	if normalSort {
		sort.Slice(sliceStr, func(i, j int) bool {
			return sliceStr[i] < sliceStr[j]
		})
	} else {
		sort.Slice(sliceStr, func(i, j int) bool {
			return sliceStr[i] > sliceStr[j]
		})
	}

	str = ""
	for _, sStr := range sliceStr {
		str += sStr + " "
	}

	return strings.TrimSpace(str)
}

// удаляет дублирующие строки, если таковые есть. (создает новый слайс пропуская не нужные элементы)
func deletDublicate(sliceStrings []string) []string {

	for i, str := range sliceStrings {
		for j := i + 1; j < len(sliceStrings); j++ {
			if str == sliceStrings[j] {
				sliceStrings = append(sliceStrings[:i], sliceStrings[j:]...)
			}
		}
	}
	return sliceStrings
}

// считывает построчно файл и возвращает слайс
func readingFile(file string) []string {
	var result []string

	inFile, _ := os.Open(file)
	defer inFile.Close()

	fileScanner := bufio.NewScanner(inFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		result = append(result, str)
	}
	return result
}

// функция раписывает данные в файл
func recordFile(file string, array []string) {
	outFile, _ := os.Create(file)
	defer outFile.Close()

	for i := 0; i < len(array)-1; i++ {
		outFile.WriteString(array[i] + "\n")
	}
	outFile.WriteString(array[len(array)-1])
}
