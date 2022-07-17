package main

import (
	"fmt"
	"time"
)

type Computer interface {
	TurnON()
	TurnOFF()
}
type PC struct {
	Monitor  *Monitor
	Keyboard *Keyboard
	Mouse    *Mouse
}

func (p PC) TurnON() {
	fmt.Println("Computer turns on...")
	time.Sleep(time.Duration(1) * time.Second)
	p.Monitor.TODO("Monitor turns on...")
	p.Keyboard.TODO("Keyboard turns on...")
	p.Mouse.TODO("Mouse turns on...")
}
func (p PC) TurnOFF() {
	fmt.Println("Computer turns off...")
	time.Sleep(time.Duration(1) * time.Second)
	p.Monitor.TODO("Monitor turns off...")
	p.Keyboard.TODO("Keyboard turns off...")
	p.Mouse.TODO("Mouse turns off...")
}

type Monitor struct {
}

func (m Monitor) TODO(do string) {
	fmt.Println(do)
	time.Sleep(time.Duration(1) * time.Second)
}

type Keyboard struct {
}

func (k Keyboard) TODO(do string) {
	fmt.Println(do)
	time.Sleep(time.Duration(1) * time.Second)
}

type Mouse struct {
}

func (m Mouse) TODO(do string) {
	fmt.Println(do)
	time.Sleep(time.Duration(1) * time.Second)
}

func NewPC() *PC {
	return &PC{
		Monitor:  &Monitor{},
		Keyboard: &Keyboard{},
		Mouse:    &Mouse{},
	}
}

//Паттерн Facade относится к структурным паттернам уровня объекта.
//Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс
//в виде набора имен методов к набору взаимосвязанных классов или объектов некоторой подсистемы,
//что облегчает ее использование.
//Разбиение сложной системы на подсистемы позволяет упростить процесс разработки,
//а также помогает максимально снизить зависимости одной подсистемы от другой.
//Однако использовать такие подсистемы становиться довольно сложно.
//Один из способов решения этой проблемы является паттерн Facade.
//Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.
func main() {
	comp := NewPC()
	comp.TurnON()
	for i := 0; i < 10; i++ {
		fmt.Println("WORK!")
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	comp.TurnOFF()
}
