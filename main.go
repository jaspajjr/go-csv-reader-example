package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json"address,omitempty"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func readCSV() []byte {
	csvFile, _ := os.Open("example_data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}
	peopleJSON, _ := json.Marshal(people)
	return peopleJSON
}

func main() {

	foo := readCSV()
	fmt.Println(string(foo))
}
