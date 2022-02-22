package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.


site

Реализовать утилиту wget с возможностью скачивать сайты целиком.

*/

func main() {

	u, _ := url.Parse("https://moysiytmoy.ucoz.net/")

	// создаем карту,где ключ это ссылка, а её значение будет означать статус ссылки
	// -1 - не рабочая  1 - только добавленная 2- рабочая, но не обработанная 3- обработанная
	urlPages := make(map[url.URL]int)
	urlPages[*u] = 1

	//цикл работает до тех пор, пока количество проверенных ссылок не достигнет количество добавленных
	for {

		// если обработав страницы не будет добавлено новых ссылок, то цикл прекратится
		oldCountPages := len(urlPages)

		for urlPage, statusPage := range urlPages {

			var page []byte
			// проверяем ссылку - битая небитая
			if urlPages[urlPage] == 1 {
				urlPages[urlPage], page = getPage(urlPage, statusPage)

			}
			// обрабатываем страницу
			if urlPages[urlPage] == 2 {
				// обрабатываем сыылку, попутно пытаясь найти ещё ссылок
				handlerPage(urlPages, page, u.Scheme, u.Host)
				//сохраняем страницу и меяем ей статус, дабы не сохранять несколько раз
				saveFile(urlPage, page)
				urlPages[urlPage] = 3
			}

		}
		// выход из программы, если за итерацию не было добавлено новых ссылок
		if oldCountPages == len(urlPages) {
			break
		}

	}

}

// проверка страницы
func getPage(u url.URL, statusPage int) (int, []byte) {
	//
	resp, err := http.Get(u.String())
	if err != nil {

		return (-1), nil
	}
	defer resp.Body.Close()
	// если страница вернула ошибку помечаем её как битую

	if resp.StatusCode != 200 {
		return (-1), nil
	}
	// или же возвращаем тело страницы
	body, _ := io.ReadAll(resp.Body)
	return 2, body

}

// обработка страниц и поиск новых ссылок
func handlerPage(urlPages map[url.URL]int, page []byte, scheme string, host string) {
	// пытаемся найти на страницы ссылки с помощью регулярного выражения
	re := regexp.MustCompile(`href="(.*?)"`)
	hrefPage := re.FindAllSubmatch(page, -1)

	for _, path := range hrefPage {

		var newURL url.URL
		newURL.Path = string(path[1])
		newURL.Scheme = scheme
		newURL.Host = host
		// вормируем новую ссылку и добавляем ее,если ранее её не было
		if urlPages[newURL] == 0 {
			urlPages[newURL] = 1
		}

	}

}

// сохраняем страницу
func saveFile(urlFile url.URL, body []byte) {

	path := urlFile.Path
	//формируем локальный адрес файла и его имя
	file := path[strings.LastIndex(path, `/`)+1:]
	path = strings.Replace(path, file, "", 1)
	path = strings.Replace(path, `.`, "", -1)

	if !strings.Contains(file, `.`) {
		if len(file) > 0 {
			file += `.html`
		} else {
			file = `index.html`
		}
	}
	// добавляем локальный каталор и пытаемся сформировать копию сайта
	f, err := os.OpenFile(urlFile.Host+path+file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		os.MkdirAll(urlFile.Host+path, 0777)
		f, _ = os.OpenFile(urlFile.Host+path+file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	}
	defer f.Close()
	f.Write(body)

}
