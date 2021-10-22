package dict

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//type Words struct {
//	Dict map[string][]string
//	Done bool
//}

var wordsFromFile map[string]string = make(map[string]string)

func Read() *map[string]string {
	file, err := os.Open("english")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strArray := strings.Split(scanner.Text(), "/")
		if len(strArray) == 2 {
			key := strings.TrimSpace(strArray[0])
			val := strings.TrimSpace(strArray[1])

			var arr []string
			for _, element := range strings.Split(val, ",") {
				arr = append(arr, strings.TrimSpace(element))
			}
			wordsFromFile[key] = strings.Join(arr, ",")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &wordsFromFile
}
