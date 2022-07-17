package main

import "fmt"

type Strategy interface {
	Route(StartPoint, EndPoint int)
}

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(str Strategy) {
	n.Strategy = str
}

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(StartPoint, EndPoint int) {
	AvgSpeed := 30
	TrafficJam := 2
	Total := EndPoint - StartPoint
	TotalTime := Total * 40 * TrafficJam
	fmt.Printf("RoadStrategy A:[%d] B:[%d] Avg Speed:[%d] Traffic jam:[%d] Total:[%d] Total time:[%d] \n",
		StartPoint, EndPoint, AvgSpeed, TrafficJam, Total, TotalTime)

}

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(StartPoint, EndPoint int) {
	AvgSpeed := 40
	Total := EndPoint - StartPoint
	TotalTime := Total * 40
	fmt.Printf("PublicTransportStrategy A:[%d] B:[%d] Avg Speed:[%d]  Total:[%d] Total time:[%d] \n",
		StartPoint, EndPoint, AvgSpeed, Total, TotalTime)

}

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(StartPoint, EndPoint int) {
	AvgSpeed := 4
	Total := EndPoint - StartPoint
	TotalTime := Total * 60
	fmt.Printf("WalkStrategy A:[%d] B:[%d] Avg Speed:[%d] Total:[%d] Total time:[%d] \n",
		StartPoint, EndPoint, AvgSpeed, Total, TotalTime)

}

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&RoadStrategy{},
		&PublicTransportStrategy{},
		&WalkStrategy{},
	}
)

//Стратегия — поведенческий шаблон проектирования, предназначенный
//для определения семейства алгоритмов, инкапсуляции каждого из них и обеспечения их взаимозаменяемости.
//Это позволяет выбирать алгоритм путём определения соответствующего класса.
//Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов, которые его используют.

//Шаблон стратегия позволяет переключаться между алгоритмами или стратегиями в зависимости от ситуации.
func main() {
	nav := Navigator{}
	for _, str := range strategies {
		nav.SetStrategy(str)
		nav.Route(start, end)
	}
}
