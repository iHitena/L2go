package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.



site


Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

func main() {
	go startServer()
	// даем время серверу запуститься
	time.Sleep(1 * time.Second)

	timeOut := flag.Int("timeout", 10, "Time out flag")
	flag.Parse()
	// коннектимся к серверу с некоторым таймаутом
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8081", time.Duration(*timeOut)*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	// запускаем горутину, которая будет считывать данные из стдин и отрпавлять на сервер.
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			// если потом закрыть, то закроется и соединение - cntl + D
			if err == io.EOF {
				conn.Close()
			}
			fmt.Fprint(conn, text+"\n")
		}
	}()

	for {
		mes, err := bufio.NewReader(conn).ReadString('\n')
		// если соединение закрыто, то выйдет ошибка
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Client: " + mes)
	}

}

func startServer() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8081")

	conn, _ := ln.Accept()

	for {
		// читает сообщеия и возвращает новые
		mes, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Mes server: ", mes)

		mes = "new " + mes

		conn.Write([]byte(mes + "\n"))
	}
}
