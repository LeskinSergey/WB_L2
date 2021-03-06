package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

//Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
//
//Реализовать поддержку утилитой следующих ключей:
//-f - "fields" - выбрать поля (колонки)
//-d - "delimiter" - использовать другой разделитель
//-s - "separated" - только строки с разделителем
func main() {
	f := flag.Int("f", 1, "выбор колонки")
	d := flag.String("d", "\t", "выбор делиметра")
	s := flag.Bool("s", false, "только строки с разделителями")
	flag.Parse()
	if *f <= 0 {
		log.Fatal("f <= 0")
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		txt := sc.Text()
		splitTxt := strings.Split(txt, *d)
		if *s && !strings.Contains(txt, *d) {
			fmt.Println("")
		} else if len(splitTxt) < *f {
			fmt.Println(txt)
		} else {
			fmt.Println(splitTxt[*f-1])
		}
	}
}
