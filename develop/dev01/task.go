package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.

https://pkg.go.dev/github.com/beevik/ntp


site :

Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием
этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

*/

func main() {

	for {
		// сервер возвращает время или ошибку, если время что-то случилось
		connTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

		fmt.Println(connTime.Clock())

		if err != nil {
			// если возникла ошибка печатаем её в стдеррор и выходим из программы с кодом 1
			r := strings.NewReader(err.Error())
			io.Copy(os.Stderr, r)
			os.Exit(1)
		}

	}

}