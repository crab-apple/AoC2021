package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	copyStructure()
	downloadInputFiles()
}

func copyStructure() {
	fileInfo, err := os.Stat("cmd/day02")
	if err != nil {
		log.Panic(err)
	}
	for i := 3; i <= 25; i++ {
		os.RemoveAll(dirname(i))
		os.Mkdir(dirname(i), fileInfo.Mode())

		copyFile("cmd/day02/solution.go", dirname(i)+"/solution.go", replaceDay(2, i))
		copyFile("cmd/day02/solution_test.go", dirname(i)+"/solution_test.go", replaceDay(2, i))
	}
}

func dirname(i int) string {
	return fmt.Sprintf("cmd/day%02d", i)
}

func downloadInputFiles() {
	for i := 1; i <= 25; i++ {
		downloadInputFile(i)
	}

}

func downloadInputFile(day int) {

	client := http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)

	req, err := http.NewRequest("GET", url, nil)

	cookieVal := "replace me"
	req.Header.Set("Cookie", cookieVal)

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(resp)

	var b []byte

	b, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(b))

	err = os.WriteFile(dirname(day)+"/input.txt", b, os.FileMode(0664))

	if err != nil {
		log.Panic(err)
		//} else {
		//	fmt.Println("All ok?")
	}
}

func replaceDay(old, new int) func(string) string {
	oldStr := fmt.Sprintf("day%02d", old)
	newStr := fmt.Sprintf("day%02d", new)
	return func(s string) string {
		return strings.Replace(s, oldStr, newStr, -1)
	}
}

func copyFile(from string, to string, transformer func(s string) string) {

	fileInfo, _ := os.Stat(from)

	b, _ := os.ReadFile(from)
	contents := string(b)
	contents = transformer(contents)
	os.WriteFile(to, []byte(contents), fileInfo.Mode())
}
