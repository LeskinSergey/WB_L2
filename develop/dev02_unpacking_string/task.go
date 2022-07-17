package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	s = strings.Trim(s, "\n")
	Srune := []rune(s)
	if s == "" || Srune[len(Srune)-1] == 92 || unicode.IsDigit(Srune[0]) {
		return "", errors.New("некорректная строка")
	}
	var isSlash bool
	res := make([]rune, 0)
	for i, val := range Srune {
		if isSlash && unicode.IsLetter(val) {
			return "", errors.New("некорректная строка")
		}
		if !isSlash && val == '\\' {
			isSlash = true
			continue
		}
		if isSlash {
			res = append(res, val)
			isSlash = false
			continue
		}
		if unicode.IsDigit(val) {
			n, err := strconv.Atoi(string(val))
			if err != nil {
				return "", err
			}
			if n == 0 {
				res = res[:len(res)-1]
				continue
			}
			n--
			for n > 0 {
				res = append(res, res[i-1])
				n--
			}
			continue
		}
		res = append(res, val)
	}
	return string(res), nil
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s, err := Unpack(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
