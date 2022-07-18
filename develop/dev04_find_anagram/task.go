package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagram(s []string) map[string][]string {
	res := make(map[string][]string)
	tmpS := GetLower(s)
	tmpS = RemoveDup(tmpS)
	tmpM := make(map[string][]string)
	for _, val := range tmpS {
		runeS := []rune(val)
		sort.Slice(runeS, func(i, j int) bool {
			return runeS[i] < runeS[j]
		})
		tmpM[string(runeS)] = append(tmpM[string(runeS)], val)
	}
	for _, v := range tmpM {
		if len(v) > 1 {
			res[v[0]] = v
		}
	}

	return res

}

func GetLower(s []string) []string {
	for _, v := range s {
		v = strings.ToLower(v)
	}
	return s
}

func RemoveDup(s []string) []string {
	m := make(map[string]bool)
	res := make([]string, 0)
	for _, val := range s {
		_, ok := m[val]
		if !ok {
			m[val] = true
			res = append(res, val)
		}
	}
	return res
}

//Написать функцию поиска всех множеств анаграмм по словарю.
//
//
//Например:
//'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
//'листок', 'слиток' и 'столик' - другому.
//
//Требования:
//Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
//Выходные данные: ссылка на мапу множеств анаграмм
//Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
//слово из множества.
//Массив должен быть отсортирован по возрастанию.
//Множества из одного элемента не должны попасть в результат.
//Все слова должны быть приведены к нижнему регистру.
//В результате каждое слово должно встречаться только один раз.
func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "пятка", "листок", "кольцо"}
	output := FindAnagram(input)
	fmt.Println(output)
}
