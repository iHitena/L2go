package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.'

site


Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то) Chdir
- pwd - показать путь до текущего каталога. Getwd
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)  kill(fing(p) )
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат* / libr PS


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).


*/

func main() {

	scaner := bufio.NewScanner(os.Stdin)
	// сканируем стдИН и обрабатываем команды
	for scaner.Scan() {
		commandHandler(scaner.Text())
	}

}

func commandHandler(stringCommand string) {
	// отделяем команду от её тела
	command := strings.Split(stringCommand, " ")[0]

	// выходим из программы по команде
	switch command {
	case `\quit`:
		fmt.Println("exit")
		os.Exit(0)

		// смена директории
	case "cd":
		cd := strings.Replace(stringCommand, "cd ", "", 1)
		os.Chdir(cd)
	// путь до текущей директории
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
		// вывод аргумента в стдаут
	case "echo":
		str := strings.Replace(stringCommand, "echo ", "", 1)
		fmt.Println(str)

		// убийство процессора
	case "kill":
		strProc := strings.Replace(stringCommand, "kill ", "", 1)
		pid, err := strconv.Atoi(strProc)
		if err != nil {
			fmt.Println(err)
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println(err)
		}
		proc.Kill()

		//вывод всех процессов
	case "ps":
		sliceProc, _ := ps.Processes()

		for _, proc := range sliceProc {
			fmt.Printf("Name p: %v Pid: %v\n", proc.Executable(), proc.Pid())
		}

	default:
		fmt.Println("command not recognized")
	}
}
