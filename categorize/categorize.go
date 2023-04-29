package categorize

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Cat struct {
	Category string
	Words    Entry
}

type Entry struct {
	K string `json:"k"`
	R string `json:"r"`
	E string `json:"e"`
}

func include(str string, array []string) bool {
	for _, a := range array {
		if a == str {
			return true
		}
	}
	return false
}

func rf(filepath string) []string {
	out := []string{}
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return out
}

func read_categories_from_json(str string) map[string]Entry {
	jsonFile, err := os.Open(str)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	var temp map[string]Entry
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &temp)
	return temp
}

var N1 = read_categories_from_json("./data/N1_vocab.json")
var N2 = read_categories_from_json("./data/N2_vocab.json")
var N3 = read_categories_from_json("./data/N3_vocab.json")
var N4 = read_categories_from_json("./data/N4_vocab.json")
var N5 = read_categories_from_json("./data/N5_vocab.json")

func Categorize(str string) Cat {

	// Open our jsonFile

	// fmt.Println(N1["唱える"])
	// var categorization_list Cat
	if val, ok := N1[str]; ok {
		// print(val.K, " is N1 \n")
		return Cat{"N1", val}
	}
	if val, ok := N2[str]; ok {
		// print(val.K, " is N2 \n")
		return Cat{"N2", val}
	}
	if val, ok := N3[str]; ok {
		// print(val.K, " is N3 \n")
		return Cat{"N3", val}
	}
	if val, ok := N4[str]; ok {
		// print(val.K, " is N4 \n")
		return Cat{"N4", val}

	}
	if val, ok := N5[str]; ok {
		print(val.K, " is N5 \n")
		// return Cat{"N5", val}

	}
	return Cat{"UN", Entry{str, "", ""}}
	// categorization_list["undefined"][c] = Entry{c, "", ""}
}
