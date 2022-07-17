package main

import "fmt"

const (
	AsusCollectorType = "Asus"
	HpCollectorType   = "HP"
)

type Collector interface {
	SetCore()
	SetMemory()
	SetGraphicCard()
	SetBrand()
	GetComputer() ComputerB
}

type ComputerB struct {
	Core        int
	Memory      int
	GraphicCard int
	Brand       string
}
type AsusCollector struct {
	Core        int
	Memory      int
	GraphicCard int
	Brand       string
}

func (a *AsusCollector) SetCore() {
	a.Core = 8
}
func (a *AsusCollector) SetMemory() {
	a.Memory = 8
}
func (a *AsusCollector) SetGraphicCard() {
	a.GraphicCard = 1
}
func (a *AsusCollector) SetBrand() {
	a.Brand = "Asus"
}
func (a *AsusCollector) GetComputer() ComputerB {
	return ComputerB{
		Core:        a.Core,
		Memory:      a.Memory,
		GraphicCard: a.GraphicCard,
		Brand:       a.Brand,
	}
}

type HpCollector struct {
	Core        int
	Memory      int
	GraphicCard int
	Brand       string
}

func (h *HpCollector) SetCore() {
	h.Core = 6
}
func (h *HpCollector) SetMemory() {
	h.Memory = 16
}
func (h *HpCollector) SetGraphicCard() {
	h.GraphicCard = 1
}
func (h *HpCollector) SetBrand() {
	h.Brand = "HP"
}
func (h *HpCollector) GetComputer() ComputerB {
	return ComputerB{
		Core:        h.Core,
		Memory:      h.Memory,
		GraphicCard: h.GraphicCard,
		Brand:       h.Brand,
	}
}

func (c *ComputerB) Print() {
	fmt.Printf("%s Core:[%d] Mem:[%d] GR:[%d]\n", c.Brand, c.Core, c.Memory, c.GraphicCard)
}
func GetCollector(CollectorType string) Collector {
	switch CollectorType {
	case HpCollectorType:
		return &HpCollector{}
	case AsusCollectorType:
		return &AsusCollector{}
	default:
		return nil
	}
}

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}
func (f *Factory) SetCollector(collector Collector) {
	f.Collector = collector
}
func (f *Factory) CreateComputer() ComputerB {
	f.Collector.SetCore()
	f.Collector.SetMemory()
	f.Collector.SetGraphicCard()
	f.Collector.SetBrand()
	return f.Collector.GetComputer()
}

//Паттерн Builder относится к порождающим паттернам уровня объекта.
//
//Паттерн Builder определяет процесс поэтапного построения сложного продукта.
//После того как будет построена последняя его часть, продукт можно использовать.
func main() {
	AsusCollector := GetCollector(AsusCollectorType)
	HpCollector := GetCollector(HpCollectorType)

	factory := NewFactory(AsusCollector)
	AsusComputer := factory.CreateComputer()
	AsusComputer.Print()

	factory.SetCollector(HpCollector)
	HpComputer := factory.CreateComputer()
	HpComputer.Print()
}
