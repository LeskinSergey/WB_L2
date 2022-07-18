package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadData(file string) [][]string {
	dataFile := make([]string, 0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		dataFile = append(dataFile, scan.Text())
		//fmt.Println(scan.Text(), "  ")
	}
	//fmt.Println(dataFile)
	parts := make([][]string, 0)
	for _, elem := range dataFile {
		parts = append(parts, strings.Fields(elem))
	}
	return parts
}
func RemoveDup(before [][]string) [][]string {
	afterMap := make(map[string]bool, 0)
	for _, val := range before {
		str := strings.Join(val, " ")
		_, ok := afterMap[str]
		if !ok {
			afterMap[str] = true
		}
		//fmt.Println(str)
	}
	after := make([][]string, 0)
	for k := range afterMap {
		after = append(after, strings.Split(k, " "))
	}
	//fmt.Println(after)
	return after
}
func GoToSort(data [][]string, k int, n, r bool) [][]string {
	if r && !n {
		sort.Slice(data, func(i, j int) bool {
			return data[i][k-1] > data[j][k-1]
		})
	} else if n && !r {
		sort.Slice(data, func(i, j int) bool {
			vi, _ := strconv.Atoi(data[i][k-1])
			vj, _ := strconv.Atoi(data[j][k-1])
			return vi < vj
		})
	} else if r && n {
		sort.Slice(data, func(i, j int) bool {
			vi, _ := strconv.Atoi(data[i][k-1])
			vj, _ := strconv.Atoi(data[j][k-1])
			return vi > vj
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i][k-1] < data[j][k-1]
		})
	}
	return data
}

//Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.
//
//Реализовать поддержку утилитой следующих ключей:
//
//-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
//-n — сортировать по числовому значению
//-r — сортировать в обратном порядке
//-u — не выводить повторяющиеся строки

func main() {
	k := flag.Int("k", 1, "колонка для сортировки")
	n := flag.Bool("n", false, "сортировка по числовому значению")
	r := flag.Bool("r", false, "сортировка в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	filename := flag.Arg(0)
	data := ReadData(filename)
	if *u {
		data = RemoveDup(data)
	}
	data = GoToSort(data, *k, *n, *r)
	for _, i := range data {
		for _, j := range i {
			fmt.Printf("%s ", j)
		}
		fmt.Println()
	}
}
