package dict

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Words struct {
	Dict map[string][]string
	Done bool
}

func Read() *[]Words {

	var dict []Words

	file, err := os.Open("english")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strArray := strings.Split(scanner.Text(), "/")
		if len(strArray) == 2 {
			dictLocal := make(map[string][]string)
			key := strings.ReplaceAll(strArray[0], " ", "")
			val := strings.Split(strArray[1], ",")
			stripStringInArray(val)
			dictLocal[key] = val
			element := Words{
				Dict: dictLocal,
				Done: false,
			}
			dict = append(dict, element)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &dict
}

func stripStringInArray(strArr []string) {
	for _, s := range strArr {
		s = strings.ReplaceAll(s, " ", "")
	}
}
