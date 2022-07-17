package main

import "fmt"

type Visitor interface {
	VisitSushi(p *SushiBar) string
	VisitPizza(p *Pizzeria) string
	VisitBurger(p *BurgerBar) string
}
type Place interface {
	Accept(v Visitor) string
}
type People struct {
}

func (p *People) VisitSushi(s *SushiBar) string {
	return s.Buy()
}
func (p *People) VisitPizza(s *Pizzeria) string {
	return s.Buy()
}
func (p *People) VisitBurger(s *BurgerBar) string {
	return s.Buy()
}

type City struct {
	Places []Place
}

func (c *City) Add(p Place) {
	c.Places = append(c.Places, p)
}
func (c *City) Accept(v Visitor) string {
	var res string
	for _, val := range c.Places {
		res += val.Accept(v)
	}
	return res
}

type SushiBar struct {
}

func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushi(s)
}
func (s *SushiBar) Buy() string {
	return "...buy sushi..."
}

type Pizzeria struct {
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizza(p)
}
func (p *Pizzeria) Buy() string {
	return "...buy pizza..."
}

type BurgerBar struct {
}

func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurger(b)
}
func (b *BurgerBar) Buy() string {
	return "...buy burger..."
}

func NewCity() *City {
	return &City{
		Places: make([]Place, 0),
	}
}

//Паттерн Visitor относится к поведенческим паттернам уровня объекта.
//
//Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
//а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.
func main() {
	city := NewCity()
	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})
	res := city.Accept(&People{})
	fmt.Println(res)
}
