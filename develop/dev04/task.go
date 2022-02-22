package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.


site


Поиск анаграмм по словарю

Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func main() {

	arrayString := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	b := FindAnagrams(&arrayString)

	for key, value := range *b {
		fmt.Println(key, " ", *value)
	}

}

// FindAnagrams принимает массив строк и создает мапу с анаграммами
func FindAnagrams(arrayString *[]string) *map[string]*[]string {
	result := make(map[string]*[]string)

	// обрабатывае слова по одиночке
	for _, word := range *arrayString {
		// если слово из 1 символа - пропускаем его
		if len(word) <= 1 {
			continue
		}
		// слова регистранезависимы - приводим в одному регистру
		lowerWord := strings.ToLower(word)
		// создаем ключ для слова. ключем для анаграммы явлюятся руны упорядоченные по возрастанию
		sotrRuneWord := sortRune([]rune(lowerWord))
		keyWords := string(sotrRuneWord)

		// если ранее слова не добавлялись, создае новый сслайс и ссылаемся на него
		if result[keyWords] != nil {
			*result[keyWords] = addToAnamgram(*result[keyWords], lowerWord)
		} else {
			slice := []string{lowerWord}
			result[keyWords] = &slice
		}

	}

	result = sortAnagrams(result)

	return &result
}

// если слова уже были ранее добавлены в срез, то вернем старые, иноче вернем новый с добавленным новым словом
func addToAnamgram(anagram []string, word string) []string {

	for _, wordInAnagram := range anagram {
		if strings.Contains(wordInAnagram, word) {
			return anagram
		}
	}

	return append(anagram, word)
}

//сортируем анаграмму
func sortAnagrams(mapString map[string]*[]string) map[string]*[]string {

	for key, anagram := range mapString {

		slice := *anagram
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		mapString[key] = &slice
	}

	return mapString
}

// сортируем руны для созданию ключа карты
func sortRune(runeStr []rune) []rune {
	sort.Slice(runeStr, func(i, j int) bool {
		return runeStr[i] < runeStr[j]
	})
	return runeStr
}
