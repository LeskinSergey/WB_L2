package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func NewFlags() *flags {
	return &flags{
		A: *flag.Int("A", 0, "печатать +N строк после совпадения"),
		B: *flag.Int("B", 0, "печатать +N строк до совпадения"),
		C: *flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения"),
		c: *flag.Bool("c", false, "(количество строк)"),
		i: *flag.Bool("i", false, "(игнорировать регистр)"),
		v: *flag.Bool("v", false, "(вместо совпадения, исключать)"),
		F: *flag.Bool("F", false, "точное совпадение со строкой, не паттерн"),
		n: *flag.Bool("n", false, "напечатать номер строки"),
	}
}

func grep(F flags) {
	patt, data := getPatternAndSource(F)

	a, b, intRes, lineRes, err := checkFlagsAndDo(F, patt, data)
	if err != nil {
		return
	} else {
		sum := 0
		for _, val := range intRes {
			sum = min(b, val)
			for sum != 0 {
				lineRes = append(lineRes, data[val-sum])
				sum--
			}
			lineRes = append(lineRes, data[val])
			sum = 0
			for sum != min(a, len(data)-val-1) {
				lineRes = append(lineRes, data[val+sum+1])
				sum++
			}
		}
		fmt.Println(strings.Join(lineRes, "\n"))
	}
}
func getPatternAndSource(F flags) (string, []string) {
	patt := flag.Arg(0)
	fileName := flag.Arg(1)
	data := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if F.i {
			data = append(data, strings.ToLower(sc.Text()))
		} else {
			data = append(data, sc.Text())
		}
	}
	return patt, data
}
func checkFlagsAndDo(F flags, patt string, data []string) (int, int, []int, []string, error) {
	if F.F {
		patt = `\Q` + patt + `\E`
	}
	if F.i {
		patt = `(?i)` + patt
	}
	reg := regexp.MustCompile(patt)
	a := max(F.A, F.C)
	b := max(F.B, F.C)
	intRes := make([]int, 0)
	for i, val := range data {
		if reg.Match([]byte(val)) {
			intRes = append(intRes, i)
		}
	}

	if F.n {
		for i, val := range data {
			val = strconv.Itoa(i) + " " + val
		}
	}
	lineRes := make([]string, 0)
	if F.v {
		for i, val := range data {
			if findElem(intRes, i) {
				continue
			}
			lineRes = append(lineRes, val)
		}
		fmt.Println(strings.Join(lineRes, "\n"))
		return 0, 0, intRes, lineRes, errors.New("end")
	}
	if F.c {
		fmt.Println(len(intRes))
	}
	return a, b, intRes, lineRes, nil
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findElem(str []int, elem int) bool {
	for _, index := range str {
		if index == elem {
			return true
		}
	}
	return false
}

//Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
//
//Реализовать поддержку утилитой следующих ключей:
//-A - "after" печатать +N строк после совпадения
//-B - "before" печатать +N строк до совпадения
//-C - "context" (A+B) печатать ±N строк вокруг совпадения
//-c - "count" (количество строк)
//-i - "ignore-case" (игнорировать регистр)
//-v - "invert" (вместо совпадения, исключать)
//-F - "fixed", точное совпадение со строкой, не паттерн
//-n - "line num", напечатать номер строки
func main() {
	F := NewFlags()
	flag.Parse()
	grep(*F)
}
