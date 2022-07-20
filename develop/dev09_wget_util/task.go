package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getURLAndFileName() (string, string) {
	fullURLFile := flag.String("url", "", "enter website address")
	flag.Parse()
	fileURL, err := url.Parse(*fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	return *fullURLFile, fileName
}
func createFile(fileName string) (*os.File, *http.Client) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return file, &client
}
func putContentToFile(client *http.Client, file *os.File, fullURLFile string) int64 {
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	defer file.Close()

	return size
}
func main() {
	// Получаем URL, как аргумент и возвращаем URL и имя файла
	fullURLFile, fileName := getURLAndFileName()
	// Создаем файл
	file, client := createFile(fileName)
	// Получаем содержимое(HTML) и кладем в файл
	size := putContentToFile(client, file, fullURLFile)

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}
