package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f := flag.Int("f", 1, "выбор колонки")
	d := flag.String("d", "\t", "выбор делиметра")
	s := flag.Bool("s", false, "только строки с разделителями")
	flag.Parse()
	if *f <= 0 {
		log.Fatal("f <= 0")
	}
	//fmt.Println("arg", flag.Arg(0), "----d", *d, "-----s", *s, "----f", *f)
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
