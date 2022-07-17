package main

import "fmt"

type MobileAlertStater interface {
	Alert() string
}

type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{
		state: &MobileAlertVibration{},
	}
}

type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "What can i say except you're welcome..."
}

//Паттерн State относится к поведенческим паттернам уровня объекта.
//
//Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и
//является объектно-ориентированной реализацией конечного автомата.
//Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.
//
func main() {
	phone := NewMobileAlert()
	phone.SetState(&MobileAlertSong{})
	fmt.Println(phone.Alert())
}
