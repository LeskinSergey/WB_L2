package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func GetTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	} else {
		fmt.Println(time)
	}
}

//Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
//Использовать библиотеку github.com/beevik/ntp.
//Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
//
//Требования:
//Программа должна быть оформлена как go module
//Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS
func main() {
	GetTime()
}
