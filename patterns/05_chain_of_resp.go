package main

import "fmt"

type Service interface {
	Execute(d *Data)
	SetNext(service Service)
}
type Data struct {
	GetSource    bool
	UpdateSource bool
}

//источник данных
type Dev struct {
	Name string
	Next Service
}

func (device *Dev) Execute(d *Data) {
	if d.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", device.Name)
		device.Next.Execute(d)
	} else {
		fmt.Printf("Get data from device [%s]\n", device.Name)
		d.GetSource = true
		device.Next.Execute(d)
	}
}
func (device *Dev) SetNext(service Service) {
	device.Next = service
}

//обновление данных
type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(d *Data) {
	if d.UpdateSource {
		fmt.Printf("Data from device [%s] already update.\n", upd.Name)
		upd.Next.Execute(d)
	} else {
		fmt.Printf("Update data from device [%s]\n", upd.Name)
		d.GetSource = true
		upd.Next.Execute(d)
	}
}
func (upd *UpdateDataService) SetNext(service Service) {
	upd.Next = service
}

//сохранение данных
type SaveDataService struct {
	Next Service
}

func (save *SaveDataService) Execute(d *Data) {
	if !d.UpdateSource {
		fmt.Println("Data not update")
	} else {
		fmt.Println("Data save")
	}

}
func (save *SaveDataService) SetNext(service Service) {
	save.Next = service
}

func NewDevice(name string) *Dev {
	return &Dev{
		Name: name,
	}
}
func NewUpdateSvc(name string) *UpdateDataService {
	return &UpdateDataService{
		Name: name,
	}
}
func NewSaveDataService() *SaveDataService {
	return &SaveDataService{}
}

func NewData() *Data {
	return &Data{}
}

//Цепочка обязанностей — поведенческий шаблон проектирования предназначенный
//для организации в системе уровней ответственности.

//Цепочка обязанностей помогает строить цепочки объектов.
//Запрос входит с одного конца и проходит через каждый объект, пока не найдет подходящий обработчик.

//По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет.
//Если запрос не обработан, то он передается дальше по цепочке.
//Если же он обработан, то паттерн сам решает передавать его дальше или нет.
//Если запрос не обработан ни одним обработчиком, то он просто теряется.
func main() {
	Device := NewDevice("device-1")
	UpdateSvc := NewUpdateSvc("update-1")
	SaveData := NewSaveDataService()
	Device.SetNext(UpdateSvc)
	UpdateSvc.SetNext(SaveData)
	Data := NewData()
	Device.Execute(Data)
}
