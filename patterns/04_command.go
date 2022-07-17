package main

import (
	"fmt"
	"time"
)

type Command interface {
	Execute()
}
type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type Device interface {
	On()
	Off()
}
type OnCommand struct {
	device Device
}

func (o *OnCommand) Execute() {
	o.device.On()
}

type OffCommand struct {
	device Device
}

func (o *OffCommand) Execute() {
	o.device.Off()
}

type Tv struct {
	IsRunning bool
}

func (t *Tv) On() {
	t.IsRunning = true
	fmt.Println("Turning on...")
}
func (t *Tv) Off() {
	t.IsRunning = false
	fmt.Println("Turning off...")
}

func NewTv() *Tv {
	return &Tv{}
}
func NewOn(tv *Tv) *OnCommand {
	return &OnCommand{
		device: tv,
	}
}
func NewOff(tv *Tv) *OffCommand {
	return &OffCommand{
		device: tv,
	}
}
func NewButtonOn(command *OnCommand) *Button {
	return &Button{
		Command: command,
	}
}
func NewButtonOff(command *OffCommand) *Button {
	return &Button{
		Command: command,
	}
}

//Паттерн Command относится к поведенческим паттернам уровня объекта.
//
//Паттерн Command позволяет представить запрос в виде объекта.
//Из этого следует, что команда - это объект.
//Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
//
//Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
//Единственное, что должен знать инициатор, это как отправить команду.
//
//Позволяет вам инкапсулировать действия в объекты.
//Основная идея, стоящая за шаблоном — это предоставление средств, для разделения клиента и получателя.
func main() {
	tv := NewTv()
	on := NewOn(tv)
	off := NewOff(tv)
	buttonOn := NewButtonOn(on)
	buttonOn.Press()
	time.Sleep(time.Duration(3) * time.Second)
	buttonOff := NewButtonOff(off)
	buttonOff.Press()
}
